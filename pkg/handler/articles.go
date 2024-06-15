package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) ArticleByIdPage(ctx *gin.Context) {
	_articleId := ctx.Param("id")
	log.Println("ARTICLE ID: ", _articleId)

	articleId, err := strconv.ParseUint(_articleId, 0, 64)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	article, err := h.service.GetArticleById(uint(articleId))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	utils.RenderTempl(ctx, 200, layouts.ArticleByIdPage(*article))
}

func (h Handler) NewArticlePage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.NewArticlePage())
}
func (h Handler) GetAllArticles(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.ArticlesPage())
}

func (h Handler) GetArticleById(ctx *gin.Context) {
	imageUrl := "/public/logo.png"
	utils.RenderTempl(ctx, 200, layouts.ArticleByIdPage(models.Article{Title: "Article title", Content: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis. Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.", ImageUrl: &imageUrl}))
}

func (h Handler) CreateArticle(ctx *gin.Context) {
	title := ctx.Request.FormValue("title")
	image := ctx.Request.FormValue("image")
	content := ctx.Request.FormValue("content")
	newArticle := &models.Article{
		Title:    title,
		Content:  content,
		ImageUrl: &image,
	}
	fmt.Printf("New Article: %+v\n", newArticle)
	_, err := h.service.CreateArticle(newArticle, 1)
	if err != nil {
		ctx.String(200, "Erroration")
		ctx.Abort()
		return
	}
	ctx.Header("HX-Redirect", fmt.Sprintf("/articles/%v", newArticle.ID))
	//utils.RenderTempl(ctx, 200, components.MarkdownEditor(ttt.Content, ttt.Content, true))
}

func (h Handler) UpdateArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}

func (h Handler) DeleteArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}
