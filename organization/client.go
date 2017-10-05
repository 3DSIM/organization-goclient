//go:generate counterfeiter ./ Client

package organization

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/3dsim/auth0"
	"github.com/3dsim/organization-goclient/genclient"
	"github.com/3dsim/organization-goclient/genclient/operations"
	"github.com/3dsim/organization-goclient/models"
	"github.com/PuerkitoBio/rehttp"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/inconshreveable/log15"
)

// Log is a github.com/inconshreveable/log15.Logger.  Log is exposed so that users of this library can set
// their own log handler.  By default this Log uses the DiscardHandler, which discards log statements.
// See: https://godoc.org/github.com/inconshreveable/log15#hdr-Library_Use
//
// To set a different log handler do something like this:
//
// 		Log.SetHandler(log.LvlFilterHandler(log.LvlInfo, log.CallerFileHandler(log.StdoutHandler)))
var Log = log.New()

func init() {
	Log.SetHandler(log.DiscardHandler())
}

const (
	// OrganizationAPIBasePath is the base path or "slug" for the organization api
	OrganizationAPIBasePath = "organization-api"
)

// Client is a wrapper around the generated client found in the "genclient" package.  It provides convenience methods
// for common operations.  If the operation needed is not found in Client, use the "genclient" package using this client
// as an example of how to utilize the genclient.  PRs are welcome if more functionality is wanted in this client package.
type Client interface {
	Organizations() ([]*models.Organization, error)
	Organization(organizationID int32) (*models.Organization, error)
	Subscriptions(limit *int32) ([]*models.Subscription, error)
	UpdateSubscription(subscription *models.Subscription) (a *models.Subscription, err error)
	Plan(planID int32) (org *models.Plan, err error)
	OrganizationUsers(organizationID int32) (users []*models.User, err error)
}

type client struct {
	tokenFetcher auth0.TokenFetcher
	client       *genclient.Organization
	audience     string
}

// NewClient creates a new client for interacting with the 3DSIM organization api.  See the auth0 package for how to construct
// the token fetcher.  The apiGatewayURL's are as follows:
//
// 		QA 				= https://3dsim-qa.cloud.tyk.io
//		Prod and Gov 	= https://3dsim.cloud.tyk.io
//
// The audience's are:
//
// 		QA 		= https://organization-qa.3dsim.com/v2
//		Prod 	= https://organization.3dsim.com/v2
// 		Gov 	= https://organization-gov.3dsim.com
func NewClient(tokenFetcher auth0.TokenFetcher, apiGatewayURL, audience string) Client {
	return newClient(tokenFetcher, apiGatewayURL, audience, nil, openapiclient.DefaultTimeout)
}

// NewClientWithRetry creates the same type of client as NewClient, but allows for retrying any temporary errors or
// any responses with status >= 400 and < 600 for a specified amount of time.
func NewClientWithRetry(tokenFetcher auth0.TokenFetcher, apiGatewayURL, audience string, retryTimeout time.Duration) Client {
	tr := rehttp.NewTransport(
		nil, // will use http.DefaultTransport
		rehttp.RetryAny(rehttp.RetryStatusInterval(400, 600), rehttp.RetryTemporaryErr()),
		rehttp.ExpJitterDelay(1*time.Second, retryTimeout),
	)
	return newClient(tokenFetcher, apiGatewayURL, audience, tr, retryTimeout)
}

func newClient(tokenFetcher auth0.TokenFetcher, apiGatewayURL, audience string,
	roundTripper http.RoundTripper, defaultRequestTimeout time.Duration) Client {
	parsedURL, err := url.Parse(apiGatewayURL)
	if err != nil {
		message := "API Gateway URL was invalid!"
		Log.Error(message, "apiGatewayURL", apiGatewayURL)
		panic(message + " " + err.Error())
	}
	organizationTransport := openapiclient.New(parsedURL.Host, OrganizationAPIBasePath, []string{parsedURL.Scheme})
	organizationTransport.Debug = true
	if roundTripper != nil {
		organizationTransport.Transport = roundTripper
	}
	openapiclient.DefaultTimeout = defaultRequestTimeout
	organizationClient := genclient.New(organizationTransport, strfmt.Default)
	return &client{
		tokenFetcher: tokenFetcher,
		client:       organizationClient,
		audience:     audience,
	}
}

func (c *client) Organizations() (orgList []*models.Organization, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetOrganizationsParams()
	response, err := c.client.Operations.GetOrganizations(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) Organization(organizationID int32) (org *models.Organization, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewFindOrganizationByIDParams().WithID(organizationID)
	response, err := c.client.Operations.FindOrganizationByID(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) Subscriptions(limit *int32) (subscriptionList []*models.Subscription, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetSubscriptionsParams()
	if limit != nil {
		params.SetLimit(limit)
	}
	response, err := c.client.Operations.GetSubscriptions(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) UpdateSubscription(subscription *models.Subscription) (a *models.Subscription, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewPutSubscriptionParams().WithOrgID(subscription.OrganizationID).WithSubID(subscription.ID).WithSubscription(subscription)
	response, err := c.client.Operations.PutSubscription(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) Plan(planID int32) (plan *models.Plan, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	params := operations.NewGetPlanParams().WithID(planID)
	response, err := c.client.Operations.GetPlan(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) OrganizationUsers(organizationID int32) (users []*models.User, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetUsersByOrganizationParams().WithID(organizationID)
	response, err := c.client.Operations.GetUsersByOrganization(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}
