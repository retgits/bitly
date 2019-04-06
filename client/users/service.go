// Package users contains the methods to interact with the Users in Bitly
package users

import (
	"net/http"

	"github.com/retgits/bitly/client"
)

const (
	userEndpoint = "user"
)

// Users is the client object which allows you to perform operations such as changing your name
// or fetching basic user information apply only to the authenticated user.
type Users struct {
	*client.Client
}

// New creates a new instance of the Users client.
func New(c *client.Client) *Users {
	return &Users{
		c,
	}
}

// UpdateUser is to update fields in the user
func (u *Users) UpdateUser(user User) (User, error) {
	payload, err := user.marshal()
	if err != nil {
		return User{}, err
	}

	data, err := u.Call(userEndpoint, http.MethodPatch, payload)
	if err != nil {
		return User{}, err
	}

	return unmarshalUser(data)
}

// RetrieveUser is to retrieve information for the current authenticated user
func (u *Users) RetrieveUser() (User, error) {
	data, err := u.Call(userEndpoint, http.MethodGet, nil)
	if err != nil {
		return User{}, err
	}

	return unmarshalUser(data)
}
