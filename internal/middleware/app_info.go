package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "go-gin-blog-eddycjy")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
