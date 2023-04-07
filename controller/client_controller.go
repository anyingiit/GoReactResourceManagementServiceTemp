package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientController struct {
	*BaseController
}

func NewClientController(db *gorm.DB) *ClientController {
	return &ClientController{NewBaseController(db)}
}

// require RESTful API

// Get get all clients
func (i *ClientController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user1, user2, user3",
	})
}

// Post create new client
func (i *ClientController) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

// Put update client
func (i *ClientController) Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "put",
	})
}

// Delete delete client
func (i *ClientController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete",
	})
}
