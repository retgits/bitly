// Package groups contains the methods to interact with the Groups in Bitly
package groups

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/retgits/bitly/client"
)

const (
	groupsEndpoint             = "groups"
	groupDetailsEndpoint       = "groups/%s"
	groupPreferencesEndpoint   = "groups/%s/preferences"
	bitlinksByGroupEndpoint    = "groups/%s/bitlinks"
	tagsByGroupEndpoint        = "groups/%s/tags"
	metricsByCountryEndpoint   = "groups/%s/countries"
	metricsByReferrersEndpoint = "groups/%s/referring_networks"
	groupShortenCountsEndpoint = "groups/%s/shorten_counts"
	sortedBitlinksEndpoint     = "groups/%s/bitlinks/%s"
)

// Groups are a subdivision within an organization. A user will belong to a group within an organization.
// Most actions on our API will be on behalf of a group. For example, when you shorten a link, it will be
// on behalf of a user and a group.
type Groups struct {
	*client.Client
}

// New creates a new instance of the Groups client.
func New(c *client.Client) *Groups {
	return &Groups{
		c,
	}
}

// RetrieveGroups is to retrieve all groups from Bitly by organization
func (g *Groups) RetrieveGroups(organizationGUID string) (BitlyGroups, error) {
	v := url.Values{}

	if len(organizationGUID) > 0 {
		v.Add("organization_guid", organizationGUID)
	}

	queryParams := v.Encode()

	url := groupsEndpoint
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := g.Call(url, http.MethodGet, nil)
	if err != nil {
		return BitlyGroups{}, err
	}

	return unmarshalGroups(data)
}

// RetrieveGroupDetails is to retrieve details for a group
func (g *Groups) RetrieveGroupDetails(groupGUID string) (Group, error) {
	data, err := g.Call(fmt.Sprintf(groupDetailsEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return Group{}, err
	}

	return unmarshalGroupDetails(data)
}

// RetrieveGroupPreferences is to retrieve preferences for a specific group
func (g *Groups) RetrieveGroupPreferences(groupGUID string) (BitlyGroupPreferences, error) {
	data, err := g.Call(fmt.Sprintf(groupPreferencesEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return BitlyGroupPreferences{}, err
	}

	return unmarshalGroupPreferences(data)
}

// UpdateGroupDetails is to update details for a specific group
func (g *Groups) UpdateGroupDetails(groupGUID string, prefs Group) (Group, error) {
	payload, err := prefs.marshal()
	if err != nil {
		return Group{}, err
	}

	data, err := g.Call(fmt.Sprintf(groupDetailsEndpoint, groupGUID), http.MethodPatch, payload)
	if err != nil {
		return Group{}, err
	}

	return unmarshalGroupDetails(data)
}

// UpdateGroupPreferences is to update preferences for a specific group
func (g *Groups) UpdateGroupPreferences(groupGUID string, prefs BitlyGroupPreferences) (BitlyGroupPreferences, error) {
	payload, err := prefs.marshal()
	if err != nil {
		return BitlyGroupPreferences{}, err
	}

	data, err := g.Call(fmt.Sprintf(groupPreferencesEndpoint, groupGUID), http.MethodPatch, payload)
	if err != nil {
		return BitlyGroupPreferences{}, err
	}

	return unmarshalGroupPreferences(data)
}

// RetrieveBitlinksByGroup is to retrieve a paginated collection of Bitlinks for a Group
func (g *Groups) RetrieveBitlinksByGroup(groupGUID string, input *BitlinksGroupRequest) (Bitlinks, error) {
	v := url.Values{}

	if input.Size != 0 {
		v.Add("size", fmt.Sprintf("%d", input.Size))
	}

	if input.Page != 0 {
		v.Add("page", fmt.Sprintf("%d", input.Page))
	}

	if input.CreatedBefore != 0 {
		v.Add("created_before", fmt.Sprintf("%d", input.CreatedBefore))
	}

	if input.CreatedAfter != 0 {
		v.Add("created_after", fmt.Sprintf("%d", input.CreatedAfter))
	}

	if input.ModifiedAfter != 0 {
		v.Add("modified_after", fmt.Sprintf("%d", input.ModifiedAfter))
	}

	if len(input.Keyword) > 0 {
		v.Add("keyword", input.Keyword)
	}

	if len(input.Query) > 0 {
		v.Add("query", input.Query)
	}

	if len(input.Archived) > 0 {
		v.Add("archived", input.Archived)
	}

	if len(input.Deeplinks) > 0 {
		v.Add("deeplinks", input.Deeplinks)
	}

	if len(input.DomainDeeplinks) > 0 {
		v.Add("domain_deeplinks", input.DomainDeeplinks)
	}

	if len(input.CampaignGUID) > 0 {
		v.Add("campaign_guid", input.CampaignGUID)
	}

	if len(input.ChannelGUID) > 0 {
		v.Add("channel_guid", input.ChannelGUID)
	}

	if len(input.CustomBitlinks) > 0 {
		v.Add("custom_bitlink", input.CustomBitlinks)
	}

	if input.Tags != nil {
		for idx := range input.Tags {
			v.Add("tags", input.Tags[idx])
		}
	}

	if input.EncodingLogin != nil {
		for idx := range input.EncodingLogin {
			v.Add("encoding_login", input.EncodingLogin[idx])
		}
	}

	queryParams := v.Encode()

	url := fmt.Sprintf(bitlinksByGroupEndpoint, groupGUID)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := g.Call(url, http.MethodGet, nil)
	if err != nil {
		return Bitlinks{}, err
	}

	return unmarshalBitlinks(data)
}

// RetrieveTagsByGroup is to retrieve the currently used tags for a group
func (g *Groups) RetrieveTagsByGroup(groupGUID string) (Tags, error) {
	data, err := g.Call(fmt.Sprintf(tagsByGroupEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return Tags{}, err
	}

	return unmarshalTags(data)
}

// GetGroupClickMetricsByCountries will return metrics about the countries referring click traffic rolled up to a Group
func (g *Groups) GetGroupClickMetricsByCountries(groupGUID string) (Metrics, error) {
	data, err := g.Call(fmt.Sprintf(metricsByCountryEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// GetGroupClickMetricsByReferringNetworks will return metrics about the referring network click traffic rolled up to a Group
func (g *Groups) GetGroupClickMetricsByReferringNetworks(groupGUID string) (Metrics, error) {
	data, err := g.Call(fmt.Sprintf(metricsByReferrersEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// RetrieveGroupShortenCounts will get all the shorten counts for a specific group
func (g *Groups) RetrieveGroupShortenCounts(groupGUID string) (Metrics, error) {
	data, err := g.Call(fmt.Sprintf(groupShortenCountsEndpoint, groupGUID), http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// RetrieveSortedBitlinksForGroup will retrieve a paginated response for Bitlinks that are sorted for the Group
func (g *Groups) RetrieveSortedBitlinksForGroup(groupGUID string, input *SortedBitlinksGroupRequest) (Bitlinks, error) {
	v := url.Values{}

	if len(input.Unit) > 0 {
		v.Add("unit", input.Unit)
	}

	if input.Units != 0 {
		v.Add("units", fmt.Sprintf("%d", input.Units))
	}

	if len(input.UnitReference) > 0 {
		v.Add("unit_reference", input.UnitReference)
	}

	if input.Size != 0 {
		v.Add("size", fmt.Sprintf("%d", input.Size))
	}

	queryParams := v.Encode()

	url := fmt.Sprintf(sortedBitlinksEndpoint, groupGUID, input.SortType)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := g.Call(url, http.MethodGet, nil)
	if err != nil {
		return Bitlinks{}, err
	}

	return unmarshalBitlinks(data)
}
