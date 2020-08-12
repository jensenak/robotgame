package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jensenak/robotgame/qwirk"
)

func main() {
	fmt.Println("Started")

	router := gin.Default()

	router.POST("/games", func(c *gin.Context) {
		gameID := qwirk.NewGameID()
		c.JSON(200, gin.H{
			"gameID": gameID,
		})
	})

	router.Run()
}
