package store

import (
	"testing"

	"github.com/jensenak/robotgame/qwirk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	s := NewStore()
	g := qwirk.NewGame()

	err := s.Add(g)
	require.NoError(t, err)

	found, err := s.Find(g.ID)
	require.NoError(t, err)
	assert.Equal(t, g.ID, found.ID)
}
