// Package organizations contains the methods to interact with the Organizations in Bitly
package organizations

import (
	"fmt"
	"net/http"

	"github.com/retgits/bitly/client"
)

const (
	organizationDetailsEndpoint       = "organizations/%s"
	organizationsEndpoint             = "organizations"
	organizationShortenCountsEndpoint = "organizations/%s/shorten_counts"
)

// Organizations are part of our hierarchy. This is the top level where a group and user will belong.
type Organizations struct {
	*client.Client
}

// New creates a new instance of the Organizations client.
func New(c *client.Client) *Organizations {
	return &Organizations{
		c,
	}
}

// RetrieveOrganizationDetails is to retrieve details for an organization
func (o *Organizations) RetrieveOrganizationDetails(organizationGUID string) (OrganizationDetails, error) {
	data, err := o.Call(fmt.Sprintf(organizationDetailsEndpoint, organizationGUID), http.MethodGet, nil)
	if err != nil {
		return OrganizationDetails{}, err
	}

	return unmarshalOrganizationDetails(data)
}

// RetrieveOrganizations is to retrieve all organizations
func (o *Organizations) RetrieveOrganizations() (BitlyOrganizations, error) {
	data, err := o.Call(organizationsEndpoint, http.MethodGet, nil)
	if err != nil {
		return BitlyOrganizations{}, err
	}

	return unmarshalBitlyOrganizations(data)
}

// RetrieveOrganizationShortenCounts is to retrieve all the shorten counts for a specific organization
func (o *Organizations) RetrieveOrganizationShortenCounts(organizationGUID string) (Metrics, error) {
	data, err := o.Call(fmt.Sprintf(organizationShortenCountsEndpoint, organizationGUID), http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}
