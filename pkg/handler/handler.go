package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/azeek21/blog/pkg/middleware"
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

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c, c.Writer)
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

		api_group.GET("/count", middleware.AuthMiddleware(h.service, "asdfasdf"), func(ctx *gin.Context) {
			count++
			ctx.HTML(http.StatusOK, "count", fmt.Sprintf("%v", count))
		})

		auth_grup := api_group.Group("/auth")
		{
			auth_grup.POST("/sign-in", h.SignIn)
			auth_grup.POST("/sign-up", h.SignIn)
			auth_grup.POST("/sign-out", h.SignOut)
		}

	}

	page_group := enginge.Group("")
	{
		page_group.GET("/", h.IndexPage)
		page_group.GET("/sign-in", h.SignInPage)
		page_group.GET("/articles/:id", h.ArticleByIdPage)
	}

	return nil
}
