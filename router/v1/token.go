package v1

import (
	"github.com/anyingiit/GoReactResourceManagement/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitToken(routerGroup *gin.RouterGroup, db *gorm.DB) {
	controller := controller.NewTokenController(db)
	token := routerGroup.Group("/token")
	{
		token.POST("", controller.Post)
	}
}
