package materi

import (
	"fmt"
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