package models

import (
	"encoding/json"
	"time"
)

// User is a model representing a user.
type User struct {
	tableName struct{} `sql:"users,alias:users"`

	ID           int       `json:"id"`
	Username     string    `json:"username"`
	FriendCode   *string   `json:"friend_code"`
	Dexes        []Dex     `json:"dexes"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`

	Password  string     `json:"-"`
	LastIP    *string    `json:"-"`
	LastLogin *time.Time `json:"-"`
	Referrer  *string    `json:"-"`
	StripeID  *string    `json:"-"`
}

// MarshalJSON is used to satisfy the Marshaler interface to customize the way
// it's serialized into JSON.
func (u User) MarshalJSON() ([]byte, error) {
	type alias User
	return json.Marshal(&struct {
		alias
		Donated bool `json:"donated"`
	}{
		alias:   (alias)(u),
		Donated: u.StripeID != nil,
	})
}
