package materi

import "github.com/gin-gonic/gin"

type Product struct {
	Name string `json:"name"`
	Stock int `json:"stock"`
}

func CreateProductHandler(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"message": "data yg dimasukkan tidak valid",
		})

		return
	}

	if product.Stock < 0 {
		c.JSON(400, gin.H{
			"status": "error",
			"message": "stock tidak boleh kurang dari 0",
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "succes",
		"data": product,
	})
}


type Registerdata struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"min=5,max=10"`
}

func RegisterHandler(c *gin.Context) {
	var registerData Registerdata

	if err:= c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "succes",
		"data": registerData,
	})
}