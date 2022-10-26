package controller

import (
	"finalProject/database"
	"finalProject/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	Comment := model.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))
	Comment.UserID = userIdJWT

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []model.Comment{}

	err := db.Preload("User").Preload("Photo").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	responseComments := []model.ResponseGetComment{}

	for _, comment := range Comments {
		c := model.ResponseGetComment{}
		c.ID = comment.ID
		c.Message = comment.Message
		c.PhotoID = comment.PhotoID
		c.UserID = comment.UserID
		c.CreatedAt = comment.CreatedAt
		c.UpdatedAt = comment.UpdatedAt
		c.User.ID = comment.User.ID
		c.User.Email = comment.User.Email
		c.User.Username = comment.User.Username
		c.Photo.ID = comment.Photo.ID
		c.Photo.Title = comment.Photo.Title
		c.Photo.Caption = comment.Photo.Caption
		c.Photo.PhotoUrl = comment.Photo.PhotoUrl
		c.Photo.UserID = comment.Photo.UserID
		responseComments = append(responseComments, c)
	}

	c.JSON(http.StatusOK, responseComments)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	commentIdParam, _ := strconv.Atoi(c.Param("commentId"))
	contentType := c.Request.Header.Get("Content-Type")
	Comment := model.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentIdParam)

	err := db.Model(&Comment).Where("id = ?", commentIdParam).Updates(model.Comment{Message: Comment.Message}).Take(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentIdParam, _ := strconv.Atoi(c.Param("commentId"))
	Comment := model.Comment{}

	err := db.Where("id = ?", commentIdParam).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been succesfully deleted",
	})
}
