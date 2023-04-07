package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SuperAdminController struct {
	*BaseController
}

func NewSuperAdminController(db *gorm.DB) *ClientController {
	return &ClientController{NewBaseController(db)}
}

// require RESTful api

// POST /super_admin
func (i *SuperAdminController) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}
