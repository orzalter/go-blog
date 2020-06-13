package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/orzalter/go-blog/middleware/jwt"
	"github.com/orzalter/go-blog/pkg/setting"
	"github.com/orzalter/go-blog/routers/api"
	v1 "github.com/orzalter/go-blog/routers/api/v1"
)

// InitRouter init gin & routers.
func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags/:id", v1.EditTags)
		apiv1.DELETE("/tags/:id", v1.DeleteTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
