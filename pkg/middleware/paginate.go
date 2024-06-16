package middleware

import (
	"fmt"

	"github.com/azeek21/blog/models"
	"github.com/gin-gonic/gin"
)

func NewPagingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paging := models.PagingIncoming{}
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			ctx.String(400, err.Error())
			ctx.Abort()
			return
		}
		fmt.Printf("GOT PAGING INCOMING: %+v\n", paging)
		ctx.Set(models.PAGING_MODEL_NAME, paging)
		ctx.Next()
	}
}
