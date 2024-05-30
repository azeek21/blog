package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllArticles(ctx *gin.Context) {
	ctx.String(http.StatusOK, "asdfasdf asdfasdf asdfasdf asdfa df asdf asdf asdf")
}

func (h Handler) GetArticleById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}

func (h Handler) CreateArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}

func (h Handler) UpdateArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}

func (h Handler) DeleteArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}
