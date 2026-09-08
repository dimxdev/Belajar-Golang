package main

import (
	_ "belajar-gin/api-docs"
	"belajar-gin/materi"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func handler(c *gin.Context) {
	c.JSON(200, "Hello From Gin Gonic👻")
}

// @title Belajar Gin API
// @version 1.0
// @description Ini API buat latihan Gin
// @host localhost:8000
// @BasePath /
func main() {
	router := gin.Default()
	// router.Use(materi.CORSMiddleware())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.Use(materi.MyLogger()) //pakai middleware di semua route

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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