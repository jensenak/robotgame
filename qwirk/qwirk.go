package qwirk

import (
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GameServer struct {
	db *gorm.DB
}

type Game struct {
	gorm.Model
	Code    string `json:"code"`
	State   State
	Players []Player
}

type Player struct {
	gorm.Model
	GameID uint
	Name   string
}

func NewGameServer(db *gorm.DB) *GameServer {
	return &GameServer{
		db: db,
	}
}

func (gs *GameServer) Stop() {
	gs.db.Close()
}

//NewGame returns a new game with a join code
func (gs *GameServer) NewGame() *Game {
	g := &Game{
		Code:  newGameCode(),
		State: Ready,
	}

	gs.db.Create(g)

	return g
}

func (gs *GameServer) Find(code string) (*Game, error) {
	var g Game
	q := gs.db.First(&g, "code = ?", code)
	if q.RecordNotFound() {
		return nil, nil
	}
	if errs := q.GetErrors(); len(errs) > 0 {
		return nil, multierror.Append(nil, errs...)
	}
	return &g, nil
}

func newGameCode() string {
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
