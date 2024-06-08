package handler

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) SignUpPage(ctx *gin.Context) {
}

func (h Handler) SignInPage(ctx *gin.Context) {
	Render(ctx, 200, layouts.SignInPage())
}

func (h Handler) SignIn(ctx *gin.Context) {
	creds := models.SignInDTO{}
	err := ctx.ShouldBind(&creds)
	if err != nil {
		ctx.String(400, "Bad request")
		ctx.Abort()
		return
	}
	ctx.Header("HX-Redirect", "/")
	ctx.String(200, "/")

}

func (h Handler) SignUp(ctx *gin.Context) {
}

func (h Handler) SignOut(ctx *gin.Context) {}
