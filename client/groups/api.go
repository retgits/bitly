// Package groups contains the methods to interact with the Groups in Bitly
package groups

import "encoding/json"

// Bitlinks contains the Bitlink information
type Bitlinks struct {
	Links       []Link       `json:"links,omitempty"`
	Pagination  Pagination   `json:"pagination,omitempty"`
	SortedLinks []SortedLink `json:"sorted_links,omitempty"`
}

// BitlinksGroupRequest contains information to search for details of Bitlinks in a group
type BitlinksGroupRequest struct {
	// The quantity of items to be be returned
	Size int
	// Integer specifying the numbered result at which to start
	Page int
	// Custom keyword to filter on history entries
	Keyword string
	// The value that you would like to search
	Query string
	// Timestamp as an integer unix epoch
	CreatedBefore int
	// Timestamp as an integer unix epoch
	CreatedAfter int
	// Timestamp as an integer unix epoch
	ModifiedAfter int
	// Whether or not to include archived bitlinks
	Archived string
	// Filter to only Bitlinks that contain deeplinks
	Deeplinks string
	// Filter to only Bitlinks that contain deeplinks configured with a custom domain
	DomainDeeplinks string
	// Filter to return only links for the given campaign GUID, can be provided
	CampaignGUID string
	// Filter to return only links for the given channel GUID, can be provided, overrides all other parameters
	ChannelGUID    string
	CustomBitlinks string
	// Filter by given tags
	Tags []string
	// Filter by the login of the authenticated user that created the Bitlink
	EncodingLogin []string
}

// BitlyGroupPreferences is used to retrieve or update preferences
type BitlyGroupPreferences struct {
	GroupGUID        string `json:"group_guid"`
	DomainPreference string `json:"domain_preference"`
}

// BitlyGroups contains all group information
type BitlyGroups struct {
	Groups []Group `json:"groups"`
}

// Group contains group details
type Group struct {
	Created          string        `json:"created,omitempty"`
	Modified         string        `json:"modified,omitempty"`
	Bsds             []interface{} `json:"bsds,omitempty"`
	GUID             string        `json:"guid,omitempty"`
	OrganizationGUID string        `json:"organization_guid"`
	Name             string        `json:"name"`
	IsActive         bool          `json:"is_active,omitempty"`
	Role             string        `json:"role,omitempty"`
	References       References    `json:"references,omitempty"`
}

// Link contains details information on Bitlinks
type Link struct {
	CreatedAt      string     `json:"created_at"`
	ID             string     `json:"id"`
	Link           string     `json:"link"`
	CustomBitlinks []string   `json:"custom_bitlinks"`
	LongURL        string     `json:"long_url"`
	Title          string     `json:"title,omitempty"`
	Archived       bool       `json:"archived"`
	CreatedBy      string     `json:"created_by"`
	ClientID       string     `json:"client_id"`
	Tags           []string   `json:"tags"`
	Deeplinks      []string   `json:"deeplinks"`
	References     References `json:"references"`
}

// Metric contains information on the selected metric
type Metric struct {
	Value  string `json:"value"`
	Clicks int64  `json:"clicks"`
}

// Metrics is the response for a metrics request
type Metrics struct {
	UnitReference string   `json:"unit_reference"`
	Metrics       []Metric `json:"metrics"`
	Units         int64    `json:"units"`
	Unit          string   `json:"unit"`
	Facet         string   `json:"facet"`
}

// Pagination contains data if more pages are available
type Pagination struct {
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Size  int64  `json:"size"`
	Page  int64  `json:"page"`
	Total int64  `json:"total"`
}

// References contains information ob the references for a Bitlink
type References struct {
	Organization string `json:"organization,omitempty"`
	Group        string `json:"group,omitempty"`
}

// SortedBitlinksGroupRequest is the request to Bitly to create a sorted list
type SortedBitlinksGroupRequest struct {
	// The type of sorting that you would like to do
	SortType string
	// A unit of time
	Unit string
	// An integer representing the time units to query data for. pass -1 to return all units of time.
	Units int
	// An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time.
	UnitReference string
	// The quantity of items to be be returned
	Size int
}

// SortedLink contains the link details of sorted links
type SortedLink struct {
	ID     string `json:"id"`
	Clicks int64  `json:"clicks"`
}

// Tags are the tags you can associate with Bitlinks
type Tags struct {
	Tags []string `json:"tags"`
}

func (r *BitlyGroupPreferences) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Group) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalBitlinks(data []byte) (Bitlinks, error) {
	var r Bitlinks
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalGroups(data []byte) (BitlyGroups, error) {
	var r BitlyGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalGroupDetails(data []byte) (Group, error) {
	var r Group
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalGroupPreferences(data []byte) (BitlyGroupPreferences, error) {
	var r BitlyGroupPreferences
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalMetrics(data []byte) (Metrics, error) {
	var r Metrics
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalTags(data []byte) (Tags, error) {
	var r Tags
	err := json.Unmarshal(data, &r)
	return r, err
}
