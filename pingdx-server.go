package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type compareParams struct {
	Lista string `json:"lista"`
	Listb string `json:"listb"`
}

func main() {
	gin.SetMode(gin.DebugMode) //todo, prod mode
	router := gin.Default()
	// ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/compare", func(c *gin.Context) {
		params := &compareParams{}
		_ = c.BindJSON(params)
		fmt.Println(params)
		fmt.Println(params.Lista)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = router.Run(":8080")
}
