// Package users contains the methods to interact with the Users in Bitly
package users

import "encoding/json"

// Email contains email data from the logged in user
type Email struct {
	Email      string `json:"email"`
	IsPrimary  bool   `json:"is_primary"`
	IsVerified bool   `json:"is_verified"`
}

// User is the currently logged in user
type User struct {
	Created          string  `json:"created,omitempty"`
	Modified         string  `json:"modified,omitempty"`
	Login            string  `json:"login,omitempty"`
	IsActive         bool    `json:"is_active,omitempty"`
	Is2FaEnabled     bool    `json:"is_2fa_enabled,omitempty"`
	Name             string  `json:"name"`
	Emails           []Email `json:"emails,omitempty"`
	IsSsoUser        bool    `json:"is_sso_user,omitempty"`
	DefaultGroupGUID string  `json:"default_group_guid"`
}

func (r *User) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}
