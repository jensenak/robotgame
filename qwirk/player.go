package qwirk

import (
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
)

// Player is a person
type Player struct {
	gorm.Model
	Name  string
	Games []Game `gorm:"many2many:game_players;"`
}

// CreatePlayer returns a Player
func (gs *GameServer) CreatePlayer(name string) (*Player, error) {
	p, err := gs.FindPlayer(name)
	if err != nil {
		return nil, err
	}
	if p != nil {
		return p, nil
	}

	p = &Player{
		Name: name,
	}

	query := gs.db.Create(p)
	if errs := query.GetErrors(); len(errs) > 0 {
		return nil, multierror.Append(nil, errs...)
	}
	return p, nil
}

// FindPlayer takes a code and uses it to find an existing game
func (gs *GameServer) FindPlayer(name string) (*Player, error) {
	var p Player
	q := gs.db.First(&p, "name = ?", name)
	if q.RecordNotFound() {
		return nil, nil
	}
	if errs := q.GetErrors(); len(errs) > 0 {
		return nil, multierror.Append(nil, errs...)
	}
	return &p, nil
}
