package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/iBaiYang/go-gin-blog-eddycjy/docs"
	"github.com/iBaiYang/go-gin-blog-eddycjy/internal/middleware"
	v1 "github.com/iBaiYang/go-gin-blog-eddycjy/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//apiv1 := r.Group("/api/v1")
	//{
	//	apiv1.POST("/tags")
	//	apiv1.DELETE("/tags/:id")
	//	apiv1.PUT("/tags/:id")
	//	apiv1.PATCH("/tags/:id/state")
	//	apiv1.GET("/tags")
	//
	//	apiv1.POST("/articles")
	//	apiv1.DELETE("/articles/:id")
	//	apiv1.PUT("/articles/:id")
	//	apiv1.PATCH("/articles/:id/state")
	//	apiv1.GET("/articles/:id")
	//	apiv1.GET("/articles")
	//}

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
