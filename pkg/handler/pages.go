package handler

import (
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) IndexPage(ctx *gin.Context) {

	//	url := "/public/logo.png"
	//	posts := []models.Article{
	//		{
	//			Title:    "title1",
	//			Content:  "Content1",
	//			ImageUrl: &url,
	//		},
	//		{
	//			Title:    "title2",
	//			Content:  "Conten2",
	//			ImageUrl: &url,
	//		},
	//		{
	//			Title:   "title3",
	//			Content: "Content3",
	//		},
	//	}
	//	data := struct {
	//		Articles []models.Article
	//	}{
	//		Articles: posts,
	//	}
	utils.RenderTempl(ctx, 200, layouts.IndexPage())
}
