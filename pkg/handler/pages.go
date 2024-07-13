package handler

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) IndexPage(ctx *gin.Context) {
	articles, err := h.service.ArticleService.GetArticles(models.PagingIncoming{})
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}
	utils.RenderTempl(ctx, 200, layouts.IndexPage(articles))
}
