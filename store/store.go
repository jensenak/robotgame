package store

import "github.com/jensenak/robotgame/qwirk"

type Store struct {
	games map[string]*qwirk.Game
}

func NewStore() *Store {
	m := map[string]*qwirk.Game{}
	return &Store{games: m}
}

func (s *Store) Add(g *qwirk.Game) error {
	s.games[g.ID] = g
	return nil
}

func (s *Store) Find(id string) (*qwirk.Game, error) {
	if g, ok := s.games[id]; ok {
		return g, nil
	}
	return nil, nil
}
