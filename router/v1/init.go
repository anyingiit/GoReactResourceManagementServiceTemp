package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitV1(engin *gin.Engine, db *gorm.DB) {
	v1 := engin.Group("/v1")
	InitUser(v1, db)
	InitClient(v1, db)
	InitToken(v1, db)
	InitProject(v1, db)
}
