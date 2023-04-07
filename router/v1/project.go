package v1

import (
	"github.com/anyingiit/GoReactResourceManagement/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProject(routerGroup *gin.RouterGroup, db *gorm.DB) {
	controller := controller.NewProjectController(db)
	project := routerGroup.Group("/project")
	{
		project.POST("", controller.Post)
	}
}
