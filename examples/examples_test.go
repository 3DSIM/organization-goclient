package examples

import (
	"fmt"
	organizationclient "github.com/3dsim/organization-goclient/client"
	"github.com/3dsim/organization-goclient/client/operations"
	"github.com/3dsim/organization-goclient/models"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func ExampleUseOfAPIWithAuthentication() {
	token := "my token fetched from auth0"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := organizationclient.New(openapiclient.New("3dsim.cloud.tyk.io", "organization-api", []string{"https"}), strfmt.Default)
	organizationToCreate := &models.Organization{
		Name:                   stringToPointer("Sample Org"),
		Street:                 stringToPointer("1 way"),
		City:                   stringToPointer("Park City"),
		State:                  stringToPointer("UT"),
		PostalCode:             stringToPointer("84098"),
		ProductID:              int64ToPointer(0),
		RunningSimulationLimit: int64ToPointer(1),
		Active:                 boolToPointer(true),
	}
	createdOrganization, err := client.Operations.PostOrganizations(operations.NewPostOrganizationsParams().WithOrganization(organizationToCreate), bearerTokenAuth)
	if err != nil {
		// You have a problem
		return
	}

	fmt.Printf("Result: %v\n", createdOrganization)
}

func stringToPointer(text string) *string {
	return &text
}

func boolToPointer(b bool) *bool {
	return &b
}

func int64ToPointer(i int64) *int64 {
	return &i
}
