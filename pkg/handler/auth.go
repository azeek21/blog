package handler

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) SignUpPage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.SignUpPage())
}

func (h Handler) SignInPage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.SignInPage())
}

func (h Handler) SignIn(ctx *gin.Context) {
	creds := models.SignInDTO{}
	err := ctx.ShouldBind(&creds)
	if err != nil {
		ctx.String(400, "Bad request")
		ctx.Abort()
		return
	}
}

func (h Handler) SignUp(ctx *gin.Context) {
	var creds models.SignUpDTO
	err := ctx.ShouldBind(&creds)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}
	if creds.Password != creds.PasswordRepeat {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "Password and it's repeat should match"))
		ctx.Abort()
		return
	}
	// TODO implement signin jwt
	utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.SUCCESS, "Sign Up success. You'll be redirected to main page soon..."))
}

func (h Handler) SignOut(ctx *gin.Context) {}
