package routers

import (
	_ "github.com/KarasWinds/blog-service/docs"
	"github.com/KarasWinds/blog-service/internal/middleware"
	v1 "github.com/KarasWinds/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	articles := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
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
