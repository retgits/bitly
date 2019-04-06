// Package organizations contains the methods to interact with the Organizations in Bitly
package organizations

import "encoding/json"

// BitlyOrganizations is a toplevel struct containing all organizations
type BitlyOrganizations struct {
	Organizations []OrganizationDetails `json:"organizations"`
}

// Metrics contains the data for a metrics request
type Metrics struct {
	UnitReference string   `json:"unit_reference"`
	Metrics       []Metric `json:"metrics"`
	Units         int64    `json:"units"`
	Unit          string   `json:"unit"`
	Facet         string   `json:"facet"`
}

// Metric contains data on the chosen metric
type Metric struct {
	Value int64  `json:"value"`
	Key   string `json:"key"`
}

// OrganizationDetails contains detailed information on the organization
type OrganizationDetails struct {
	Created         string        `json:"created"`
	Modified        string        `json:"modified"`
	Bsds            []interface{} `json:"bsds"`
	GUID            string        `json:"guid"`
	Name            string        `json:"name"`
	IsActive        bool          `json:"is_active"`
	Tier            string        `json:"tier"`
	TierFamily      string        `json:"tier_family"`
	TierDisplayName string        `json:"tier_display_name"`
	Role            string        `json:"role"`
	References      References    `json:"references"`
}

// References contains the groups which an organization is part of
type References struct {
	Groups string `json:"groups"`
}

func unmarshalBitlyOrganizations(data []byte) (BitlyOrganizations, error) {
	var r BitlyOrganizations
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalMetrics(data []byte) (Metrics, error) {
	var r Metrics
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalOrganizationDetails(data []byte) (OrganizationDetails, error) {
	var r OrganizationDetails
	err := json.Unmarshal(data, &r)
	return r, err
}
