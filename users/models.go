package users

import (
	"time"

	"github.com/pokedextracker/api.pokedextracker.com/dexes"
)

// User is a model representing a user.
type User struct {
	tableName struct{} `sql:"users,alias:users"`

	ID           int         `json:"id"`
	Username     string      `json:"username"`
	FriendCode   string      `json:"friend_code"`
	Dexes        []dexes.Dex `json:"dexes"`
	Donated      bool        `sql:"-" json:"donated"`
	DateCreated  time.Time   `json:"date_created"`
	DateModified time.Time   `json:"date_modified"`

	Password  string    `json:"-"`
	LastIP    string    `json:"-"`
	LastLogin time.Time `json:"-"`
	Referrer  string    `json:"-"`
	StripeID  string    `json:"-"`
}
