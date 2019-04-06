package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// BitlyBaseURL is the base URL for the Bitly v4 API
	BitlyBaseURL = "https://api-ssl.bitly.com/v4/"
)

// Client contains all the functions to communicate between Bitly and your app.
type Client struct {
	// Many of Bitly's API methods require an OAuth access token for authentication.
	// You can generate a generic access token by confirming your password on https://bitly.is/accesstoken.
	AccessToken string
}

// NewClient returns a new Client pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewClient() *Client {
	return &Client{}
}

// WithAccessToken sets a config AccessToken value returning a Client pointer for
// chaining.
func (c *Client) WithAccessToken(accessToken string) *Client {
	c.AccessToken = accessToken
	return c
}

// Call sends a request to Bitly and receives the response.
func (c *Client) Call(urlSuffix string, httpMethod string, payload []byte) ([]byte, error) {
	var req *http.Request
	var err error

	if len(payload) > 0 {
		req, err = http.NewRequest(httpMethod, fmt.Sprintf("%s%s", BitlyBaseURL, urlSuffix), strings.NewReader(string(payload)))
	} else {
		req, err = http.NewRequest(httpMethod, fmt.Sprintf("%s%s", BitlyBaseURL, urlSuffix), nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
