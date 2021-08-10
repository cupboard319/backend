package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	mevents "github.com/micro/micro/v3/service/events"
	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	cproto "github.com/m3o/services/customers/proto"
	eproto "github.com/m3o/services/emails/proto"
	onboarding "github.com/m3o/services/onboarding/proto"
	authproto "github.com/micro/micro/v3/proto/auth"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/client"
	mconfig "github.com/micro/micro/v3/service/config"
	cont "github.com/micro/micro/v3/service/context"
	merrors "github.com/micro/micro/v3/service/errors"
	logger "github.com/micro/micro/v3/service/logger"
	model "github.com/micro/micro/v3/service/model"
	mstore "github.com/micro/micro/v3/service/store"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://127.0.0.1:4200/google-login",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

const (
	microNamespace   = "micro"
	internalErrorMsg = "An error occurred during onboarding. Contact #m3o-support at slack.m3o.com if the issue persists"
	topic            = "onboarding"
)

const (
	expiryDuration = 5 * time.Minute

	onboardingTopic = "onboarding"
)

type tokenToEmail struct {
	Email      string `json:"email"`
	Token      string `json:"token"`
	Created    int64  `json:"created"`
	CustomerID string `json:"customerID"`
}

type Signup struct {
	customerService cproto.CustomersService
	emailService    eproto.EmailsService
	auth            auth.Auth
	accounts        authproto.AccountsService
	config          conf
	cache           *cache.Cache
	resetCode       model.Model
	track           model.Model
}

type ResetToken struct {
	Created int64
	ID      string
	Token   string
}

type sendgridConf struct {
	TemplateID         string `json:"template_id"`
	RecoveryTemplateID string `json:"recovery_template_id"`
}

type googleConf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

type oauthConf struct {
	Google googleConf `json:"google"`
}

type conf struct {
	Sendgrid     sendgridConf `json:"sendgrid"`
	PromoCredit  int64        `json:"promoCredit"`
	PromoMessage string       `json:"promoMessage"`
	Oauth        oauthConf    `json:"oauth"`
}

func NewSignup(srv *service.Service, auth auth.Auth) *Signup {
	c := conf{}
	val, err := mconfig.Get("micro.onboarding")
	if err != nil {
		logger.Fatalf("Error getting config: %v", err)
	}
	err = val.Scan(&c)
	if err != nil {
		logger.Fatalf("Error scanning config: %v", err)
	}
	if len(c.Sendgrid.TemplateID) == 0 {
		logger.Fatalf("No sendgrid template ID provided")
	}

	oauthConfGl.ClientID = c.Oauth.Google.ClientID
	oauthConfGl.ClientSecret = c.Oauth.Google.ClientSecret
	if c.Oauth.Google.RedirectURL != "" {
		oauthConfGl.RedirectURL = c.Oauth.Google.RedirectURL
	}

	s := &Signup{
		customerService: cproto.NewCustomersService("customers", srv.Client()),
		emailService:    eproto.NewEmailsService("emails", srv.Client()),
		auth:            auth,
		accounts:        authproto.NewAccountsService("auth", srv.Client()),
		config:          c,
		cache:           cache.New(1*time.Minute, 5*time.Minute),
		resetCode:       model.New(ResetToken{}, nil),
		track: model.New(onboarding.TrackRequest{}, &model.Options{
			Key: "id",
		}),
	}
	return s
}

// taken from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// SendVerificationEmail is the first step in the onboarding flow.SendVerificationEmail
// A stripe customer and a verification token will be created and an email sent.
func (e *Signup) SendVerificationEmail(ctx context.Context,
	req *onboarding.SendVerificationEmailRequest,
	rsp *onboarding.SendVerificationEmailResponse) error {
	err := e.sendVerificationEmail(ctx, req, rsp)
	if err != nil {
		logger.Warnf("Error during sending verification email: %v", err)
	}
	return err
}

func (e *Signup) sendVerificationEmail(ctx context.Context,
	req *onboarding.SendVerificationEmailRequest,
	rsp *onboarding.SendVerificationEmailResponse) error {
	logger.Info("Received Signup.SendVerificationEmail request")

	// create entry in customers service
	crsp, err := e.customerService.Create(ctx, &cproto.CreateRequest{Email: req.Email}, client.WithAuthToken())
	if err != nil {
		logger.Error(err)
		return merrors.InternalServerError("onboarding.SendVerificationEmail", internalErrorMsg)
	}

	k := randStringBytesMaskImprSrc(8)
	tok := &tokenToEmail{
		Token:      k,
		Email:      req.Email,
		Created:    time.Now().Unix(),
		CustomerID: crsp.Customer.Id,
	}

	bytes, err := json.Marshal(tok)
	if err != nil {
		logger.Error(err)
		return merrors.InternalServerError("onboarding.SendVerificationEmail", internalErrorMsg)
	}

	if err := mstore.Write(&mstore.Record{
		Key:   req.Email,
		Value: bytes,
	}); err != nil {
		logger.Error(err)
		return merrors.InternalServerError("onboarding.SendVerificationEmail", internalErrorMsg)
	}

	// Send email
	// @todo send different emails based on if the account already exists
	// ie. registration vs login email.
	err = e.sendEmail(ctx, req.Email, e.config.Sendgrid.TemplateID, map[string]interface{}{
		"token": k,
	})
	if err != nil {
		logger.Errorf("Error when sending email to %v: %v", req.Email, err)
		return merrors.InternalServerError("onboarding.SendVerificationEmail", internalErrorMsg)
	}

	return nil
}

func (e *Signup) sendEmail(ctx context.Context, email, templateID string, templateData map[string]interface{}) error {
	b, _ := json.Marshal(templateData)
	_, err := e.emailService.Send(ctx, &eproto.SendRequest{To: email, TemplateId: templateID, TemplateData: b}, client.WithAuthToken())
	return err
}

func (e *Signup) CompleteSignup(ctx context.Context, req *onboarding.CompleteSignupRequest, rsp *onboarding.CompleteSignupResponse) error {
	err := e.completeSignup(ctx, req, rsp)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (e *Signup) completeSignup(ctx context.Context, req *onboarding.CompleteSignupRequest, rsp *onboarding.CompleteSignupResponse) error {
	logger.Info("Received Signup.CompleteSignup request")

	recs, err := mstore.Read(req.Email)
	if err == mstore.ErrNotFound {
		logger.Errorf("Can't verify record for %v: record not found", req.Email)
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	} else if err != nil {
		logger.Errorf("Error reading store: err")
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	}

	tok := &tokenToEmail{}
	if err := json.Unmarshal(recs[0].Value, tok); err != nil {
		logger.Errorf("Error when unmarshaling stored token object for %v: %v", req.Email, err)
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	}
	if tok.Token != req.Token {
		return merrors.Forbidden("onboarding.CompleteSignup", "The token you provided is invalid")
	}

	if time.Since(time.Unix(tok.Created, 0)) > expiryDuration {
		return merrors.Forbidden("onboarding.CompleteSignup", "The token you provided has expired")
	}

	rsp.CustomerID = tok.CustomerID
	if _, err := e.customerService.MarkVerified(ctx, &cproto.MarkVerifiedRequest{Email: tok.Email}, client.WithAuthToken()); err != nil {
		logger.Errorf("Error marking customer as verified: %v", err)
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	}

	// take secret from the request
	secret := req.Secret

	// generate a random secret
	if len(req.Secret) == 0 {
		secret = uuid.New().String()
	}
	_, err = e.auth.Generate(tok.CustomerID,
		auth.WithScopes("customer"),
		auth.WithSecret(secret),
		auth.WithIssuer(microNamespace),
		auth.WithName(req.Email),
		auth.WithType("customer"))
	if err != nil {
		logger.Errorf("Error generating token for %v: %v", tok.CustomerID, err)
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	}

	t, err := e.auth.Token(auth.WithCredentials(tok.CustomerID, secret), auth.WithTokenIssuer(microNamespace))
	if err != nil {
		logger.Errorf("Can't get token for %v: %v", tok.CustomerID, err)
		return merrors.InternalServerError("onboarding.CompleteSignup", internalErrorMsg)
	}
	rsp.AuthToken = &onboarding.AuthToken{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry.Unix(),
		Created:      t.Created.Unix(),
	}
	rsp.CustomerID = tok.CustomerID
	rsp.Namespace = microNamespace
	if err := mevents.Publish(topic, &onboarding.Event{Type: "newSignup", NewSignup: &onboarding.NewSignupEvent{Email: tok.Email, Id: tok.CustomerID}}); err != nil {
		logger.Warnf("Error publishing %s", err)
	}

	return nil
}

func (e *Signup) Recover(ctx context.Context, req *onboarding.RecoverRequest, rsp *onboarding.RecoverResponse) error {
	logger.Info("Received Signup.Recover request")
	_, found := e.cache.Get(req.Email)
	if found {
		return merrors.BadRequest("onboarding.recover", "We have issued a recovery email recently. Please check that.")
	}

	token := uuid.New().String()
	created := time.Now().Unix()
	err := e.resetCode.Create(ResetToken{
		ID:      req.Email,
		Token:   token,
		Created: created,
	})
	logger.Infof("Sent recovery code %v to email %v", token, req.Email)
	if err != nil {
		return merrors.InternalServerError("onboarding.recover", err.Error())
	}

	err = e.sendEmail(ctx, req.Email, e.config.Sendgrid.RecoveryTemplateID, map[string]interface{}{
		"token": token,
	})
	if err == nil {
		e.cache.Set(req.Email, true, cache.DefaultExpiration)
	}

	if err := mevents.Publish(topic, &onboarding.Event{Type: "passwordReset", PasswordReset: &onboarding.PasswordResetEvent{Email: req.Email}}); err != nil {
		logger.Warnf("Error publishing %s", err)
	}

	return err
}

func (e *Signup) ResetPassword(ctx context.Context, req *onboarding.ResetPasswordRequest, rsp *onboarding.ResetPasswordResponse) error {
	m := ResetToken{}
	if req.Email == "" {
		return errors.New("Email is required")
	}
	err := e.resetCode.Read(model.QueryEquals("ID", req.Email), &m)
	if err != nil {
		return err
	}

	if m.ID == "" {
		return errors.New("can't find token")
	}
	if m.Token == "" {
		return errors.New("can't find token")
	}
	if m.Created == 0 {
		return errors.New("expiry can't be calculated")
	}
	if m.Token != req.Token {
		return errors.New("tokens don't match")
	}
	if time.Unix(m.Created, 0).Before(time.Now().Add(-1 * 10 * time.Minute)) {
		return errors.New("token expired")
	}

	_, err = e.accounts.ChangeSecret(cont.DefaultContext, &authproto.ChangeSecretRequest{
		Id:        req.Email,
		NewSecret: req.Password,
		Options: &authproto.Options{
			Namespace: microNamespace,
		},
	}, client.WithAuthToken())
	if err != nil {
		return err
	}
	e.resetCode.Delete(model.QueryByID(m.ID))
	return err
}

func (e *Signup) Track(ctx context.Context,
	req *onboarding.TrackRequest,
	rsp *onboarding.TrackResponse) error {
	if req.Id == "" {
		return errors.New("no tracking id")
	}
	oldTrack := []*onboarding.TrackRequest{}
	err := e.track.Read(model.QueryEquals("id", req.Id), &oldTrack)
	if err != nil {
		return err
	}
	if len(oldTrack) == 0 {
		return e.track.Create(req)
	}
	// support partial update
	if req.GetFirstVisit() == 0 {
		req.FirstVisit = oldTrack[0].FirstVisit
	}
	if req.GetFirstVerification() == 0 {
		req.FirstVerification = oldTrack[0].FirstVerification
	}
	if req.Referrer == "" {
		req.Referrer = oldTrack[0].Referrer
	}
	if req.Registration == 0 {
		req.Registration = oldTrack[0].Registration
	}
	if req.Email == "" {
		req.Email = oldTrack[0].Email
	}
	return e.track.Update(req)
}

func (e *Signup) GoogleOauthURL(ctx context.Context, req *onboarding.GoogleOauthURLRequest, rsp *onboarding.GoogleOauthURLResponse) error {
	URL, err := url.Parse(oauthConfGl.Endpoint.AuthURL)
	if err != nil {
		return err
	}

	parameters := url.Values{}
	parameters.Add("client_id", oauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(oauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	//parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	logger.Info(URL.String())
	url := URL.String()
	rsp.Url = url
	return nil
}

func (e *Signup) GoogleOauthCallback(ctx context.Context, req *onboarding.GoogleOauthCallbackRequest, rsp *onboarding.GoogleOauthCallbackResponse) error {
	state := req.State

	if state != oauthStateStringGl {
		return fmt.Errorf("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
	}

	code := req.Code

	if code == "" {
		reason := req.ErrorReason
		if reason == "user_denied" {
			return fmt.Errorf("user has denied permission")
		}
		return fmt.Errorf("code not found")
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			return err
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			return fmt.Errorf("Get: " + err.Error() + "\n")
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		rsp.Response = string(response)
		return nil
	}

	return nil
}
