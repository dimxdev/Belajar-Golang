package main

import (
	"belajar-gin/materi"
	"fmt"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, "Hello From Gin Gonic👻")
}

func main() {
	router := gin.Default()
	router.Use(materi.LoggerMiddleware) //pakai middleware di semua route

	router.GET("/", handler)
	router.GET("/profile",materi.SimpleMiddleware, materi.ProfilHandler)
	router.GET("/search", materi.GetQuery)
	router.GET("/user/:id", materi.GetId)

	router.POST("/product", materi.CreateProductHandler)
	router.POST("/register", materi.RegisterHandler)

	fmt.Println("Server jalan di http://localhost:8000")
	router.Run(":8000")
}