package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/iBaiYang/go-gin-blog-eddycjy/docs"
	"github.com/iBaiYang/go-gin-blog-eddycjy/global"
	"github.com/iBaiYang/go-gin-blog-eddycjy/internal/middleware"
	"github.com/iBaiYang/go-gin-blog-eddycjy/internal/routers/api"
	v1 "github.com/iBaiYang/go-gin-blog-eddycjy/internal/routers/api/v1"
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.AccessLog())

	//if global.ServerSetting.RunMode == "debug" {
	//	r.Use(gin.Logger())
	//	r.Use(gin.Recovery())
	//} else {
	//	r.Use(middleware.AccessLog())
	//	r.Use(middleware.Recovery())
	//}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))

	r.Use(middleware.Translations())

	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 初步定义
	//apiv1 := r.Group("/api/v1")
	//{
	//	apiv1.POST("/tags")
	//	apiv1.DELETE("/tags/:id")
	//	apiv1.PUT("/tags/:id")
	//	apiv1.PATCH("/tags/:id/state")
	//  apiv1.GET("/tags/:id")
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
		apiv1.GET("/tags/:id", tag.Get)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	// 增加JWT认证过滤
	//apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
	//{
	//	apiv1.POST("/tags", tag.Create)
	//	apiv1.DELETE("/tags/:id", tag.Delete)
	//	apiv1.PUT("/tags/:id", tag.Update)
	//	apiv1.PATCH("/tags/:id/state", tag.Update)
	//	apiv1.GET("/tags/:id", tag.Get)
	//	apiv1.GET("/tags", tag.List)
	//
	//	apiv1.POST("/articles", article.Create)
	//	apiv1.DELETE("/articles/:id", article.Delete)
	//	apiv1.PUT("/articles/:id", article.Update)
	//	apiv1.PATCH("/articles/:id/state", article.Update)
	//	apiv1.GET("/articles/:id", article.Get)
	//	apiv1.GET("/articles", article.List)
	//}

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	return r
}
