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
	router.Use(materi.MyLogger()) //pakai middleware di semua route

	router.GET("/", handler)
	router.GET("/profile",materi.SimpleMiddleware, materi.ProfilHandler)

	user := router.Group("/user") //group routing (semua dibungkus "user/")
	user.Use(materi.SimpleMiddleware) //semua router user memakai simplemiddleware
	{
		user.GET("/search", materi.GetQuery)
		user.GET("/me/:id", materi.GetId)
	}

	admin := router.Group("admin/")	// "admin/..."
	{
		v1 := admin.Group("v1") // "admin/v1/..." 
		{
			v1.POST("/product", materi.CreateProductHandler)
		}
		
		v2 := admin.Group("v2")// "admin/v2/..."
		{
			v2.POST("/register",materi.AuthMiddleware, materi.RegisterHandler)
		}
	}

	fmt.Println("Server jalan di http://localhost:8000")
	router.Run(":8000")
}