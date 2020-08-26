package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jensenak/robotgame/qwirk"
	"github.com/jensenak/robotgame/store"
)

func main() {
	fmt.Println("Started")

	router := gin.Default()
	s := store.NewStore()

	router.GET("/games/:id", func(c *gin.Context) {
		id := c.Param("id")

		game, err := s.Find(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error accessing store")
			return
		}
		if game == nil {
			c.String(http.StatusNotFound, "Game not found")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": game.ID,
		})
	})

	router.POST("/games", func(c *gin.Context) {
		game := qwirk.NewGame()
		err := s.Add(game)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error storing game")
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id": game.ID,
		})
	})

	router.Run()
}
