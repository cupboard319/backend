package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	pevents "github.com/m3o/services/pkg/events"
	custpb "github.com/m3o/services/pkg/events/proto/customers"
	stripepb "github.com/m3o/services/pkg/events/proto/stripe"
	mevents "github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

func (b *Billing) consumeEvents() {
	go pevents.ProcessTopic(custpb.Topic, "billing", b.processCustomerEvents)
	go pevents.ProcessTopic(stripepb.Topic, "billing", b.processStripeEvents)
}

func (b *Billing) processCustomerEvents(ev mevents.Event) error {
	ctx := context.Background()
	ce := &custpb.Event{}
	if err := json.Unmarshal(ev.Payload, ce); err != nil {
		logger.Errorf("Error unmarshalling customer event: $s", err)
		return nil
	}
	switch ce.Type {
	case custpb.EventType_EventTypeDeleted:
		if err := b.processCustomerDelete(ctx, ce); err != nil {
			logger.Errorf("Error processing request event %s", err)
			return err
		}
	case custpb.EventType_EventTypeCreated:
		if err := b.processCustomerCreate(ctx, ce); err != nil {
			logger.Errorf("Error processing request event %s", err)
			return err
		}
	case custpb.EventType_EventTypeAddPaymentMethod:
		if err := b.processAddPaymentMethod(ctx, ce); err != nil {
			logger.Errorf("Error processing request event %s", err)
			return err
		}
	default:
		logger.Infof("Skipping event %+v", ce)
	}
	return nil

}

const (
	prefixAccByID    = "accByID"
	prefixAccByAdmin = "accByAdmin"
)

func billingAccKey(id string) string {
	return fmt.Sprintf("%s/%s", prefixAccByID, id)
}

func adminKey(userID string) string {
	return fmt.Sprintf("%s/%s", prefixAccByAdmin, userID)
}

func (b *Billing) processAddPaymentMethod(ctx context.Context, event *custpb.Event) error {
	// add a billing account if required
	// does the user who added the card have a billing account already?
	recs, err := store.Read(fmt.Sprintf("%s/%s", prefixAccByAdmin, event.Customer.Id))
	if err != nil && err != store.ErrNotFound {
		// try again
		logger.Errorf("Error looking up billing account %s", err)
		return err
	}
	if len(recs) > 0 {
		// nothing to do, billing acc already exists for this user
		return nil
	}
	// create a new billing account
	billingAcc := BillingAccount{
		ID:     uuid.New().String(),
		Admins: []string{event.Customer.Id},
	}
	return b.storeBillingAccount(&billingAcc)
}

func (b *Billing) storeBillingAccount(billingAcc *BillingAccount) error {
	for _, v := range billingAcc.Admins {
		if err := store.Write(store.NewRecord(adminKey(v), billingAcc)); err != nil {
			logger.Errorf("Error storing billing account %s", err)
			return err
		}
	}
	if err := store.Write(store.NewRecord(billingAccKey(billingAcc.ID), billingAcc)); err != nil {
		logger.Errorf("Error storing billing account %s", err)
		return err
	}
	return nil
}

func (b *Billing) processCustomerDelete(ctx context.Context, event *custpb.Event) error {
	// TODO SUBSCRIPTIONS remove sa owner of billing account, delete if no other owners of the billing account
	return nil
}

func (b *Billing) processCustomerCreate(ctx context.Context, event *custpb.Event) error {
	return nil
}

func (b *Billing) processStripeEvents(ev mevents.Event) error {
	//ctx := context.Background()
	ce := &stripepb.Event{}
	if err := json.Unmarshal(ev.Payload, ce); err != nil {
		logger.Errorf("Error unmarshalling customer event: $s", err)
		return nil
	}
	switch ce.Type {
	// TODO
	default:
		logger.Infof("Skipping event %+v", ce)
	}
	return nil

}
