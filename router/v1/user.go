package v1

import (
	"github.com/anyingiit/GoReactResourceManagement/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUser(routerGroup *gin.RouterGroup, db *gorm.DB) {
	controller := controller.NewUserController(db)
	user := routerGroup.Group("/user")
	{
		user.GET("", controller.Get)
	}
}
