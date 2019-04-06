// Package bsds contains the methods to interact with the Branded Short Domains in Bitly
package bsds

import "encoding/json"

// BSD contains all Branded Short Domains
type BSD struct {
	AllBSDs []string `json:"bsds"`
}

func unmarshalAllBSD(data []byte) (BSD, error) {
	var r BSD
	err := json.Unmarshal(data, &r)
	return r, err
}
