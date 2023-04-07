package middleware

import "github.com/gin-gonic/gin"

func AuthBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do something
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do something
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do something
	}
}

func AuthSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do something
	}
}
