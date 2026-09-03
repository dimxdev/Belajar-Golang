package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, "Hello From Gin Gonic👻")
}

// GET
type Profile struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func profilHandler(c *gin.Context) {
	data := Profile{
		Name: "Dimas",
		Age: 20,
	}

	c.JSON(200, data)
}

func main() {
	router := gin.Default()

	router.GET("/", handler)
	router.GET("/profile", profilHandler)

	fmt.Println("Server jalan di http://localhost:8000")
	router.Run(":8000")
}