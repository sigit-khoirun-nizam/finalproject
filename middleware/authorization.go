package middleware

import (
	"finalProject/database"
	"finalProject/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization(c *gin.Context) {
	db := database.GetDB()

	userIdParam, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	User := model.User{}

	err = db.Select("id").First(&User, uint(userIdParam)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))

	if User.ID != userIdJWT {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}

func PhotoAuthorization(c *gin.Context) {
	db := database.GetDB()

	photoIdParam, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	Photo := model.Photo{}
	err = db.Select("user_id").First(&Photo, uint(photoIdParam)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))

	if Photo.UserID != userIdJWT {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}

func CommentAuthorization(c *gin.Context) {
	db := database.GetDB()

	commentIdParam, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	Comment := model.Comment{}
	err = db.Select("user_id").First(&Comment, uint(commentIdParam)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))

	if Comment.UserID != userIdJWT {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}

func SocialMediaAuthorization(c *gin.Context) {
	db := database.GetDB()

	socialMediaIdParam, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	SocialMedia := model.SocialMedia{}
	err = db.Select("user_id").First(&SocialMedia, uint(socialMediaIdParam)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))

	if SocialMedia.UserID != userIdJWT {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}
