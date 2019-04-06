// Package bsds contains the methods to interact with the Branded Short Domains in Bitly
package bsds

import (
	"net/http"

	"github.com/retgits/bitly/client"
)

const (
	bsdEndpoint = "bsds"
)

// BSDs is an acronym for branded short domains. This is a custom 15 character or less domain for bitlinks.
// This allows you to customize the domain to your brand.
type BSDs struct {
	*client.Client
}

// New creates a new instance of the BSDs client.
func New(c *client.Client) *BSDs {
	return &BSDs{
		c,
	}
}

// GetBSDs is to Fetch all Branded Short Domains
func (b *BSDs) GetBSDs() (BSD, error) {
	data, err := b.Call(bsdEndpoint, http.MethodGet, nil)
	if err != nil {
		return BSD{}, err
	}

	return unmarshalAllBSD(data)
}
