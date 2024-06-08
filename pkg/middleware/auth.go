package middleware

import (
	"errors"
	"net/http"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const AUTH_SLUG = "blog-session"

func AuthMiddleware(userService service.UserService, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenCookie, err := c.Request.Cookie(AUTH_SLUG)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		token := tokenCookie.Value
		jwtParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unauthorized")
			}
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil || !jwtParsed.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := jwtParsed.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		user, err := userService.GetUserById(claims["sub"].(uint))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set(models.USER_MODEL_NAME, user)
		c.Next()
	}
}

func RoleMiddleware(allowRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _user, ok := c.Get(models.USER_MODEL_NAME); ok {
			user := _user.(models.User)
			if user.RoleCode != allowRole {
				c.String(http.StatusUnauthorized, "You don't have enough privilages to complete this action")
				c.Abort()
				return
			}
			c.Next()
			return
		}
		c.String(http.StatusUnauthorized, "You don't have enough privilages to complete this action")
		c.Abort()
	}
}
