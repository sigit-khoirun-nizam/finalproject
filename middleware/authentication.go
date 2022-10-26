package middleware

import (
	"finalProject/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	verifyToken, err := helper.VerifyToken(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthenticated",
			"message": err.Error(),
		})
		return
	}
	c.Set("userData", verifyToken)
	c.Next()
}
