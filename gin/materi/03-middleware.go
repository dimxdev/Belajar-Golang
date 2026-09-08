package materi

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	fmt.Println("Ada request masuk ke: ", c.Request.URL.Path)
	c.Next()
}

func SimpleMiddleware(c *gin.Context) {
	fmt.Println("ini middleware paling simple")
	c.Next()
}

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(401, gin.H{
			"status": "error",
			"message": "token kosong",
		})

		c.Abort() //berhenti disini middleware ga di lanjut ke setelahnya
		return
	}

	c.Next()
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		fmt.Println("Request Masuk:", c.Request.Method, c.Request.URL.Path)
		c.Next()

		duration := time.Since(start)
		fmt.Println("Request selesai dalam:", duration)
	}
}

// middleware cors manual/native
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}