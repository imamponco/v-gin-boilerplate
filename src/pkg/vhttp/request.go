package vhttp

import (
	"encoding/json"
	"net/http"
)

func ParseJSONBody(r *http.Request, dest interface{}) error {
	// Decode body request body
	d := json.NewDecoder(r.Body)
	if err := d.Decode(dest); err != nil {
		return err
	}
	return nil
}
