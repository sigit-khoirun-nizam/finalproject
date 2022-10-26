package controller

import (
	"finalProject/database"
	"finalProject/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	Photo := model.Photo{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))
	Photo.UserID = userIdJWT

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))
	Photos := []model.Photo{}

	err := db.Where("user_id = ?", userIdJWT).Preload("User").Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	responsePhotos := []model.ResponseGetPhoto{}

	for _, photo := range Photos {
		p := model.ResponseGetPhoto{}
		p.ID = photo.ID
		p.Title = photo.Title
		p.Caption = photo.Caption
		p.PhotoUrl = photo.PhotoUrl
		p.UserID = photo.UserID
		p.CreatedAt = photo.CreatedAt
		p.UpdatedAt = photo.UpdatedAt
		p.User.Email = photo.User.Email
		p.User.Username = photo.User.Username
		responsePhotos = append(responsePhotos, p)
	}

	c.JSON(http.StatusOK, responsePhotos)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	photoIdParam, _ := strconv.Atoi(c.Param("photoId"))
	contentType := c.Request.Header.Get("Content-Type")
	Photo := model.Photo{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.ID = uint(photoIdParam)

	err := db.Model(&Photo).Where("id = ?", photoIdParam).Updates(model.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoIdParam, _ := strconv.Atoi(c.Param("photoId"))
	Photo := model.Photo{}

	err := db.Where("id = ?", photoIdParam).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been succesfully deleted",
	})
}
