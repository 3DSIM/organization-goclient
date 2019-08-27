package organization

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/3dsim/auth0/auth0fakes"
	"github.com/3dsim/organization-goclient/models"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	apiBasePath = "base-path"
	audience    = "test audience"
)

func TestOrganizationsWhenSuccessfulExpectsOrganizationListReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	listToReturn := []*models.Organization{
		&models.Organization{
			ID:   1,
			Name: swag.String("Organization 1"),
		},
		&models.Organization{
			ID:   2,
			Name: swag.String("Organization 2"),
		},
	}

	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		bytes, err := json.Marshal(listToReturn)
		if err != nil {
			t.Error("Failed to marshal organization list")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationEndpoint := "/" + apiBasePath + "/organizations"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.Organizations()

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, list, "Expected returned organization list to not be nil")
	assert.Equal(t, len(listToReturn), len(list), "Expected organization list length to match reference list length")
	assert.Equal(t, int32(1), listToReturn[0].ID, "Expected IDs to match")
	assert.Equal(t, "Organization 1", *listToReturn[0].Name, "Expected names to match")
}

func TestOrganizationsWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", apiBasePath, audience)

	// act
	list, err := client.Organizations()

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}

func TestOrganizationsWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationEndpoint := "/" + apiBasePath + "/organizations"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.Organizations()

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api send a 500 error")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}

func TestOrganizationWhenSuccessfulExpectsOrganizationReturned(t *testing.T) {
	// arrange
	organizationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	organizationToReturn := &models.Organization{
		ID:   organizationID,
		Name: swag.String("Organization name"),
	}
	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedOrganizationID, err := strconv.Atoi(mux.Vars(r)["organizationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(organizationID), receivedOrganizationID, "Expected organization id received to match what was passed in")
		bytes, err := json.Marshal(organizationToReturn)
		if err != nil {
			t.Error("Failed to marshal organization")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationEndpoint := "/" + apiBasePath + "/organizations/{organizationID}"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	organization, err := client.Organization(organizationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, organization, "Expected returned organization to not be nil")
	assert.Equal(t, *organizationToReturn.Name, *organization.Name, "Expected names to match")
	assert.Equal(t, organizationToReturn.ID, organization.ID, "Expected IDs to match")
}

func TestOrganizationWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	organizationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", apiBasePath, audience)

	// act
	response, err := client.Organization(organizationID)

	// assert
	assert.Nil(t, response, "Expected response to be nil because token fetcher returned error")
	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestOrganizationWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	organizationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationEndpoint := "/" + apiBasePath + "/organizations/{organizationID}"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	response, err := client.Organization(organizationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api sent a 500 error")
	assert.Nil(t, response, "Expected response to be nil because organization api sent a 500 error")
}

func TestNewClientWithRetryWhen500ExpectsRetry(t *testing.T) {
	// arrange
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	callCounter := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCounter++
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/"+apiBasePath+"/organizations/1", handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClientWithRetry(fakeTokenFetcher, testServer.URL, apiBasePath, audience, 3*time.Second)

	// act
	_, err := client.Organization(1)

	// assert
	assert.True(t, callCounter > 1, "Expected to retry the failed call at least once")
	assert.NotNil(t, err, "Expected an error returned because organization api sent a 500 error")
}

func TestSubscriptionsWhenSuccessfulExpectsSubscriptionsListReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Subscription
	listToReturn := []*models.Subscription{
		&models.Subscription{
			ID: 1,
		},
		&models.Subscription{
			ID: 2,
		},
	}

	subscriptionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		bytes, err := json.Marshal(listToReturn)
		if err != nil {
			t.Error("Failed to marshal subscription list")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	subscriptionsEndpoint := "/" + apiBasePath + "/subscriptions"
	r.HandleFunc(subscriptionsEndpoint, subscriptionHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.Subscriptions(nil)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, list, "Expected returned organization list to not be nil")
	assert.Equal(t, len(listToReturn), len(list), "Expected organization list length to match reference list length")
	assert.Equal(t, int32(1), listToReturn[0].ID, "Expected IDs to match")
}

func TestSubscriptionsWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", apiBasePath, audience)

	// act
	list, err := client.Subscriptions(nil)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}

func TestSubscriptionsWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	subscriptionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	subscriptionEndpoint := "/" + apiBasePath + "/subscriptions"
	r.HandleFunc(subscriptionEndpoint, subscriptionHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.Subscriptions(nil)

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api send a 500 error")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}

func TestUpdateSubscriptionWhenSuccessfulExpectsSubscriptionReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Subscription
	subscription := &models.Subscription{
		ID:             1,
		OrganizationID: 1,
	}

	subscriptionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		bytes, err := json.Marshal(subscription)
		if err != nil {
			t.Error("Failed to marshal organization list")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	subscriptionsEndpoint := "/" + apiBasePath + "/organizations/1/subscriptions/1"
	r.HandleFunc(subscriptionsEndpoint, subscriptionHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	subResponse, err := client.UpdateSubscription(subscription)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, subResponse, "Expected subResponse to not be nil")
	assert.Equal(t, int32(1), subResponse.ID, "Expected IDs to match")
}

func TestUpdateSubscriptionWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	expectedError := errors.New("Some auth0 error")
	subscription := &models.Subscription{
		ID:             1,
		OrganizationID: 1,
	}
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", apiBasePath, audience)

	// act
	subResponse, err := client.UpdateSubscription(subscription)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
	assert.Nil(t, subResponse, "Expected subResponse to be nil")
}

func TestUpdateSubscriptionWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	subscription := &models.Subscription{
		ID:             1,
		OrganizationID: 1,
	}

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	subscriptionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	subscriptionEndpoint := "/" + apiBasePath + "/organizations/1/subscriptions/1"
	r.HandleFunc(subscriptionEndpoint, subscriptionHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	subResponse, err := client.UpdateSubscription(subscription)

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api send a 500 error")
	assert.Nil(t, subResponse, "Expected subResponse to be nil")
}

func TestPlanWhenSuccessfulExpectsPlanReturned(t *testing.T) {
	// arrange
	planID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	planToReturn := &models.Plan{
		ID:   planID,
		Name: swag.String("Plan name"),
	}
	planHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		receivedPlanID, err := strconv.Atoi(mux.Vars(r)["planID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(planID), receivedPlanID, "Expected plan id received to match what was passed in")
		bytes, err := json.Marshal(planToReturn)
		if err != nil {
			t.Error("Failed to marshal organization")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	planEndpoint := "/" + apiBasePath + "/plans/{planID}"
	r.HandleFunc(planEndpoint, planHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	plan, err := client.Plan(planID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, plan, "Expected returned plan to not be nil")
	assert.Equal(t, *planToReturn.Name, *plan.Name, "Expected names to match")
	assert.Equal(t, planToReturn.ID, plan.ID, "Expected IDs to match")
}

func TestPlanWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	planID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	planHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	planEndpoint := "/" + apiBasePath + "/plans/{planID}"
	r.HandleFunc(planEndpoint, planHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	response, err := client.Plan(planID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api sent a 500 error")
	assert.Nil(t, response, "Expected response to be nil because organization api sent a 500 error")
}

func TestOrganizationUsersWhenSuccessfulExpectsUserListReturned(t *testing.T) {
	// arrange
	orgID := int32(1)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	listToReturn := []*models.User{
		&models.User{
			FirstName: "User 1",
		},
		&models.User{
			FirstName: "User 2",
		},
	}

	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		bytes, err := json.Marshal(listToReturn)
		if err != nil {
			t.Error("Failed to marshal organization list")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationUserEndpoint := "/" + apiBasePath + "/organizations/{orgId}/users"
	r.HandleFunc(organizationUserEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.OrganizationUsers(orgID)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, list, "Expected returned organization list to not be nil")
	assert.Equal(t, len(listToReturn), len(list), "Expected organization list length to match reference list length")
	assert.Equal(t, "User 1", listToReturn[0].FirstName, "Expected Names to match")
}

func TestOrganizationUsersWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	orgID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", apiBasePath, audience)

	// act
	list, err := client.OrganizationUsers(orgID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}

func TestOrganizationUsersWhenOrganizationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	orgID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Organization
	organizationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	organizationEndpoint := "/" + apiBasePath + "/organizations/{orgID}/users"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, apiBasePath, audience)

	// act
	list, err := client.OrganizationUsers(orgID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because organization api send a 500 error")
	assert.Nil(t, list, "Expected list of organizations to be nil")
}
