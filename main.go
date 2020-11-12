package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jensenak/robotgame/qwirk"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type gameResp struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type createPlayerRequest struct {
	Name string `json:"name"`
}

type playerResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Started")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=qwirk password=qwirk dbname=qwirk sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&qwirk.Game{}, &qwirk.Player{})

	gs := qwirk.NewGameServer(db)

	router := gin.Default()

	router.GET("/games/:code", func(c *gin.Context) {
		code := c.Param("code")

		game, err := gs.Find(code)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error accessing store")
			return
		}
		if game == nil {
			c.String(http.StatusNotFound, "Game not found")
			return
		}
		c.JSON(http.StatusOK, gameResp{Code: game.Code, State: game.State.String()})
	})

	router.POST("/games", func(c *gin.Context) {
		game := gs.NewGame()

		c.JSON(http.StatusCreated, gameResp{Code: game.Code, State: game.State.String()})
	})

	router.POST("/players", func(c *gin.Context) {
		var request createPlayerRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid JSON")
			return
		}

		p, err := gs.CreatePlayer(request.Name)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating player")
			return
		}
		c.JSON(http.StatusOK, playerResp{ID: p.ID, Name: p.Name})
	})

	router.Run()
}
