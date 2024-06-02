package handler

import (
	"fmt"
	"net/http"

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

	}

	page_group := enginge.Group("")
	{
		page_group.GET("/", h.IndexPage)
		page_group.GET("/login", h.LoginPage)
		page_group.GET("/articles/:id", h.ArticlePage)
	}

	// static
	return nil
}
