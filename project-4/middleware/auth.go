package middleware

import (
	"context"
	"net/http"
	"project-4/tool"

	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")

		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "Authorization header required!")
			return
		}

		jwtToken, err := tool.TokenValidate(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid token")
			return
		}

		claims, ok := jwtToken.Claims.(*tool.MyClaim)
		if !ok && !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid token")
			return
		}

		ctx := context.WithValue(c.Request.Context(), userCtxKey, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AuthContext(c *gin.Context) *tool.MyClaim {
	raw, _ := c.Request.Context().Value(userCtxKey).(*tool.MyClaim)
	return raw
}
