package qwirk

import (
	"strings"

	"github.com/hashicorp/go-uuid"
)

// Game is a game
type Game struct {
	ID string
}

//NewGame returns a game with an ID
func NewGame() *Game {
	return &Game{ID: newGameID()}
}

func newGameID() string {
	u, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	parts := strings.Split(u, "-")
	if len(parts) == 0 {
		panic("Empty ID")
	}
	return parts[0]
}
