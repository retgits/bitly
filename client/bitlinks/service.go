// Package bitlinks contains the methods to interact with the Bitlinks in Bitly
package bitlinks

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/retgits/bitly/client"
)

const (
	expandEndpoint                  = "expand"
	bitlinksEndpoint                = "bitlinks"
	updateBitlinkEndpoint           = "bitlinks/%s"
	shortenEndpoint                 = "shorten"
	bitlinksClickSummaryEndpoint    = "bitlinks/%s/clicks/summary"
	bitlinksClickEndpoint           = "bitlinks/%s/clicks"
	bitlinksCountryEndpoint         = "bitlinks/%s/country"
	bitlinksReferrersEndpoint       = "bitlinks/%s/referrers"
	bitlinksReferrersDomainEndpoint = "bitlinks/%s/referrers_by_domains"
	bitlinksReferringDomainEndpoint = "bitlinks/%s/referring_domains"
)

// Bitlinks is how we refer to shortened links. You can see these with the bit.ly domain or your
// custom branded short domain. (Example: bit.ly/ABCDE)
type Bitlinks struct {
	*client.Client
}

// New creates a new instance of the Bitlinks client.
func New(c *client.Client) *Bitlinks {
	return &Bitlinks{
		c,
	}
}

// ExpandBitlink returns public information for a Bitlink.
func (b *Bitlinks) ExpandBitlink(link Link) (LinkInfo, error) {
	payload, err := link.marshal()
	if err != nil {
		return LinkInfo{}, err
	}

	data, err := b.Call(expandEndpoint, http.MethodPost, payload)
	if err != nil {
		return LinkInfo{}, err
	}

	return unmarshalLinkInfo(data)
}

// GetMetricsByCountries will return metrics about the countries referring click traffic to a single Bitlink.
func (b *Bitlinks) GetMetricsByCountries(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksCountryEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// GetMetricsByReferrers will return metrics about the referrers referring click traffic to a single Bitlink.
func (b *Bitlinks) GetMetricsByReferrers(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksReferrersEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// GetMetricsByReferrersAndDomain will group referrers metrics about a single Bitlink.
func (b *Bitlinks) GetMetricsByReferrersAndDomain(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksReferrersDomainEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// GetMetricsByReferringDomains will rollup the click counts to a referrer about a single Bitlink.
func (b *Bitlinks) GetMetricsByReferringDomains(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksReferringDomainEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// CreateBitlink will convert a long url to a Bitlink and set additional parameters.
func (b *Bitlinks) CreateBitlink(bitlink *Bitlink) (BitlinkDetails, error) {
	payload, err := bitlink.marshal()
	if err != nil {
		return BitlinkDetails{}, err
	}

	data, err := b.Call(bitlinksEndpoint, http.MethodPost, payload)
	if err != nil {
		return BitlinkDetails{}, err
	}

	return unmarshalBitlinkDetails(data)
}

// GetClicksSummary will return the click counts for a specified Bitlink. This rolls up all the data into a single field of clicks.
func (b *Bitlinks) GetClicksSummary(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksClickSummaryEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// GetClicks will return the click counts for a specified Bitlink. This returns an array with clicks based on a date.
func (b *Bitlinks) GetClicks(bitlink string, input *MetricsRequest) (Metrics, error) {
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

	url := fmt.Sprintf(bitlinksClickEndpoint, bitlink)
	if len(queryParams) > 1 {
		url = fmt.Sprintf("%s?%s", url, queryParams)
	}

	data, err := b.Call(url, http.MethodGet, nil)
	if err != nil {
		return Metrics{}, err
	}

	return unmarshalMetrics(data)
}

// UpdateBitlink will update fields in the Bitlink.
func (b *Bitlinks) UpdateBitlink(bitlink string, bitlinkDetails *BitlinkDetails) (BitlinkDetails, error) {
	payload, err := bitlinkDetails.marshal()
	if err != nil {
		return BitlinkDetails{}, err
	}

	data, err := b.Call(fmt.Sprintf(updateBitlinkEndpoint, bitlink), http.MethodPatch, payload)
	if err != nil {
		return BitlinkDetails{}, err
	}

	return unmarshalBitlinkDetails(data)
}

// RetrieveBitlink returns information for a Bitlink.
func (b *Bitlinks) RetrieveBitlink() {}

// ShortenLink will convert a long url to a Bitlink.
func (b *Bitlinks) ShortenLink(bitlink *ShortenRequest) (BitlinkDetails, error) {
	payload, err := bitlink.marshal()
	if err != nil {
		return BitlinkDetails{}, err
	}

	data, err := b.Call(shortenEndpoint, http.MethodPost, payload)
	if err != nil {
		return BitlinkDetails{}, err
	}

	return unmarshalBitlinkDetails(data)
}
