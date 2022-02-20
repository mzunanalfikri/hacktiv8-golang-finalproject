package middleware

import (
	"fmt"
	"net/http"
	"project-4/service"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim := AuthContext(c)

		if claim == nil {
			fmt.Println("lalalalallilili")

			c.AbortWithStatusJSON(http.StatusForbidden, "Authorization header required!")
			return
		}

		user, err := service.GetUserDetail(claim.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, "User is not admin")
			return
		}

		c.Next()
	}
}
