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
	p := &Player{
		Name: name,
	}
	query := gs.db.Create(p)
	if errs := query.GetErrors(); len(errs) > 0 {
		return nil, multierror.Append(nil, errs...)
	}
	return p, nil
}
