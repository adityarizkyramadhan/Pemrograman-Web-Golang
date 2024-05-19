package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jika path url hanya / saja maka kembalikan 303
		bearerToken, err := c.Cookie("session_token")
		if errors.Is(err, http.ErrNoCookie) && c.Request.URL.Path == "/" {
			c.Status(http.StatusSeeOther)
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			c.JSON(400, model.NewErrorResponse(err.Error()))
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["user_id"].(float64)
			userIdInt := int(userId)
			c.Set("id", userIdInt)
			c.Next()
			return
		} else {
			c.Abort()
			c.JSON(401, model.NewErrorResponse("unauthorized"))
			return
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return model.JwtKey, nil
}
