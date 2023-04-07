package middleware

import (
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
)

// 判断项目是否已经初始化
func CheckProjectInitialized() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := globalVars.Db.GetDb()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		var count int64
		if err := db.Model(&models.Sys{}).Count(&count).Error; err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		
	}
}
