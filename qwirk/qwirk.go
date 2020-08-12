package qwirk

import "github.com/hashicorp/go-uuid"

// NewGameID generates a game ID
func NewGameID() string {
	u, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	return u
}
