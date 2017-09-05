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
	audience = "test audience"
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
	organizationEndpoint := "/" + OrganizationAPIBasePath + "/organizations"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", audience)

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
	organizationEndpoint := "/" + OrganizationAPIBasePath + "/organizations"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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
	organizationEndpoint := "/" + OrganizationAPIBasePath + "/organizations/{organizationID}"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", audience)

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
	organizationEndpoint := "/" + OrganizationAPIBasePath + "/organizations/{organizationID}"
	r.HandleFunc(organizationEndpoint, organizationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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
	r.HandleFunc("/"+OrganizationAPIBasePath+"/organizations/1", handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClientWithRetry(fakeTokenFetcher, testServer.URL, audience, 3*time.Second)

	// act
	_, err := client.Organization(1)

	// assert
	assert.True(t, callCounter > 1, "Expected to retry the failed call at least once")
	assert.NotNil(t, err, "Expected an error returned because organization api sent a 500 error")
}
