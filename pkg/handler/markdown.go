package handler

import (
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/gin-gonic/gin"
)

func (h Handler) ShowMarkdownPreview(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	utils.RenderTempl(ctx, 200, components.MarkdownEditor(content, content, true))
}

func (h Handler) ShowMarkdownEditor(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	utils.RenderTempl(ctx, 200, components.MarkdownEditor(content, content, false))
}
