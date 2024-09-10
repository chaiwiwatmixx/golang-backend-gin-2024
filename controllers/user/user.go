package usercontroller

import (
	"net/http"

	"example.com/gin-api/configs"
	"example.com/gin-api/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "users",
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	result := configs.DB.Debug().Create(&user)

	// db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(201, gin.H{
		"message": "สมัครสมาชิกสำหรับ",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"data": id,
	})
}

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}
