package v1

import (
	"github.com/anyingiit/GoReactResourceManagement/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitClient(routerGroup *gin.RouterGroup, db *gorm.DB) {
	clientController := controller.NewClientController(db)

	client := routerGroup.Group("/client")
	{
		client.GET("", clientController.Get)
	}
}
