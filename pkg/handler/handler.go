package handler

import (
	"github.com/azeek21/blog/pkg/middleware"
	"github.com/azeek21/blog/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

		auth_grup := api_group.Group("/auth")
		{
			auth_grup.POST("/sign-in", h.SignIn)
			auth_grup.POST("/sign-up", h.SignUp)
			auth_grup.POST("/sign-out", h.SignOut)
		}

		markdown_group := api_group.Group("/markdown")
		{
			markdown_group.POST("/show-preview", h.ShowMarkdownPreview)
			markdown_group.POST("/show-edit", h.ShowMarkdownEditor)
		}

	}

	page_group := enginge.Group("", middleware.AuthMiddleware(h.service.UserService, h.service.JwtService, viper.GetString("PRIVATE_KEY"), false))
	{
		page_group.GET("/", middleware.NewPagingMiddleware(), h.IndexPage)
		page_group.GET("/sign-in", h.SignInPage)
		page_group.GET("/sign-up", h.SignUpPage)
		page_group.GET("/articles/new", middleware.AuthMiddleware(h.service.UserService, h.service.JwtService, viper.GetString("PRIVATE_KEY"), true), h.NewArticlePage)
		page_group.GET("/articles/:id", h.ArticleByIdPage)
	}

	return nil
}
