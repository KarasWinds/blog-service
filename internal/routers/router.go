package routers

import (
	"net/http"
	"time"

	_ "github.com/KarasWinds/blog-service/docs"
	"github.com/KarasWinds/blog-service/global"
	"github.com/KarasWinds/blog-service/internal/middleware"
	"github.com/KarasWinds/blog-service/internal/routers/api"
	v1 "github.com/KarasWinds/blog-service/internal/routers/api/v1"
	"github.com/KarasWinds/blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	articles := v1.NewArticle()
	tag := v1.NewTag()

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.POST("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", articles.Create)
		apiv1.DELETE("/articles/:id", articles.Delete)
		apiv1.PUT("/articles/:id", articles.Update)
		apiv1.PATCH("/articles/:id/state", articles.Update)
		apiv1.GET("/articles/:id", articles.Get)
		apiv1.GET("/articles", articles.List)
	}

	return r
}
