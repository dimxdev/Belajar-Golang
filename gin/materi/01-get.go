package materi

import "github.com/gin-gonic/gin"

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ProfilHandler(c *gin.Context) {
	data := Profile{
		Name: "Dimas",
		Age:  20,
	}

	c.JSON(200, data)
}

func GetQuery(c *gin.Context) {
	name := c.Query("name")
	page := c.DefaultQuery("page", "1")

	c.JSON(200, gin.H{
		"name": name,
		"page": page,
		"status": "succes",
	})
}

func GetId(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"status": "succes",
		"id": id,
	})
}