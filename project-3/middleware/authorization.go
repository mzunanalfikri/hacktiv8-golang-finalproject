package middleware

import (
	"context"
	"net/http"
	"project-3/tool"

	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")

		// Allow unauthenticated users in
		if authToken == "" {
			c.Next()
			return
		}

		//validate jwt token
		jwtToken, err := tool.TokenValidate(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid token")
			return
		}

		//validate claim
		claims, ok := jwtToken.Claims.(*tool.MyClaim)
		if !ok && !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid token")
			return
		}

		//return user data to req
		ctx := context.WithValue(c.Request.Context(), userCtxKey, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func AuthContext(c *gin.Context) *tool.MyClaim {
	raw, _ := c.Request.Context().Value(userCtxKey).(*tool.MyClaim)
	return raw
}
