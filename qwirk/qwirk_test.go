package qwirk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGameServer_Find_Existing(t *testing.T) {
	db := TestDB(t)

	gs := NewGameServer(db)
	game := gs.NewGame()

	found, err := gs.Find(game.Code)
	require.NoError(t, err)

	assert.Equal(t, game.Code, found.Code)
}

func TestGameServer_Find_NotFound(t *testing.T) {
	db := TestDB(t)

	gs := NewGameServer(db)

	found, err := gs.Find("arbitrary")
	require.NoError(t, err)
	assert.Nil(t, found)
}

func TestGameServer_CreatePlayer(t *testing.T) {
	db := TestDB(t)

	gs := NewGameServer(db)

	player, err := gs.CreatePlayer("alice")
	require.NoError(t, err)
	assert.Equal(t, "alice", player.Name)

	var found Player
	q := db.Find(&found, player.ID)
	require.Empty(t, q.GetErrors())
	assert.Equal(t, player.Name, found.Name)
}

// func TestGameServer_Join_Existing(t *testing.T) {
// 	db := TestDB(t)

// 	gs := NewGameServer(db)
// 	game := gs.NewGame()

// 	player, err := gs.Join(game.Code, "Alice")
// 	require.NoError(t, err)

// }
