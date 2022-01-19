package handler

import (
	"context"
	"testing"
	"time"

	billing "github.com/m3o/services/billing/proto"
	custevents "github.com/m3o/services/pkg/events/proto/customers"
	stripe "github.com/m3o/services/stripe/proto"
	"github.com/m3o/services/stripe/proto/fakes"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/events"
	evmem "github.com/micro/micro/v3/service/events/stream/memory"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/micro/v3/service/store/memory"

	. "github.com/onsi/gomega"
)

func TestSubscribeTier(t *testing.T) {
	tcs := []struct {
		name       string
		tier       string
		subErr     error
		unsubErr   error
		testErr    error
		subCount   int
		unsubCount int
		billingAcc *BillingAccount
	}{
		{
			name:     "upgrade to pro",
			tier:     "pro",
			subCount: 1,
			billingAcc: &BillingAccount{
				ID:     "1234",
				Admins: []string{"1", "2"},
			},
		},
		{
			name:       "downgrade to free",
			tier:       "free",
			subCount:   0,
			unsubCount: 1,
			billingAcc: &BillingAccount{
				ID:      "1234",
				Admins:  []string{"1"},
				PriceID: "pro",
				SubID:   "sub_12134",
			},
		},
		{
			name:       "switch between paid tiers",
			tier:       "team",
			subCount:   1,
			unsubCount: 1,
			billingAcc: &BillingAccount{
				ID:      "1234",
				Admins:  []string{"1"},
				PriceID: "pro",
				SubID:   "sub_12134",
			},
		},
		{
			// error state, we shouldn't get in to this position
			name:       "no billing account",
			tier:       "pro",
			testErr:    errors.InternalServerError("billing.SubscribeTier", "Error processing subscription, please try again"),
			subCount:   0,
			unsubCount: 0,
			billingAcc: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			store.DefaultStore = memory.NewStore()
			events.DefaultStream, _ = evmem.NewStream()
			stripeSvc := &fakes.FakeStripeService{}
			stripeSvc.SubscribeReturns(&stripe.SubscribeResponse{SubscriptionId: "100"}, tc.subErr)
			stripeSvc.UnsubscribeReturns(&stripe.UnsubscribeResponse{}, tc.unsubErr)
			b := Billing{
				stripeSvc: stripeSvc,
				tiers: map[string]string{
					"free": "free",
					"pro":  "2",
					"team": "3",
				},
			}

			if tc.billingAcc != nil {
				b.storeBillingAccount(tc.billingAcc)
			}

			evs, _ := events.DefaultStream.Consume(custevents.Topic)

			g := NewWithT(t)

			req := billing.SubscribeTierRequest{Id: tc.tier, CardId: "pm_1234"}
			rsp := billing.SubscribeTierResponse{}
			err := b.SubscribeTier(auth.ContextWithAccount(context.TODO(), &auth.Account{
				ID:     "1",
				Type:   "customer",
				Issuer: "micro",
				Scopes: []string{"customer"},
				Name:   "dom",
			}), &req, &rsp)
			if tc.testErr == nil {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(Equal(tc.testErr))
			}
			g.Expect(stripeSvc.SubscribeCallCount()).To(Equal(tc.subCount))
			g.Expect(stripeSvc.UnsubscribeCallCount()).To(Equal(tc.unsubCount))

			if tc.testErr == nil {
				select {
				case <-evs:
					// event published successfully
				case <-time.After(time.Second):
					t.Errorf("Failed to publish event")
				}

			}

		})
	}

}
