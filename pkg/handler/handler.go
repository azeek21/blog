package handler

import (
	"fmt"
	"net/http"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) RegisterHandlers(enginge *gin.Engine, staticPath string) error {
	enginge.Static("public", staticPath)
	count := 0

	api_group := enginge.Group("/api")
	{
		articles := api_group.Group("/articles")
		{
			articles.GET("/", h.GetAllArticles)
			articles.GET("/:id", h.GetArticleById)
			articles.POST("/", h.CreateArticle)
			articles.PUT("/:id", h.UpdateArticle)
			articles.DELETE("/:id", h.DeleteArticle)
		}

		api_group.GET("/count", func(ctx *gin.Context) {
			count++
			ctx.HTML(http.StatusOK, "count", fmt.Sprintf("%v", count))
		})

		auth_grup := api_group.Group("/auth")
		{
			auth_grup.POST("/signin", func(ctx *gin.Context) {
				creds := models.SignInDTO{}
				err := ctx.ShouldBind(&creds)
				if err != nil {
					ctx.String(400, "Bad request")
					ctx.Abort()
					return
				}
				ctx.Header("HX-Redirect", "/")
				ctx.String(200, "/")
			})
			auth_grup.POST("/sign-out")
			auth_grup.POST("/update-user")
		}

	}

	page_group := enginge.Group("")
	{
		page_group.GET("/", h.IndexPage)
		page_group.GET("/signin", h.LoginPage)
		page_group.GET("/articles/:id", h.ArticlePage)
	}

	// static
	return nil
}
