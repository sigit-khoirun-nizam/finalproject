package controller

import (
	"finalProject/database"
	"finalProject/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	SocialMedia := model.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))
	SocialMedia.UserID = userIdJWT

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetAllSocialMediasByUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userIdJWT := uint(userData["id"].(float64))

	SocialMedias := []model.SocialMedia{}
	err := db.Where("user_id = ?", userIdJWT).Preload("User").Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	responseSocialMedias := []model.ResponseGetSocialMedia{}

	for _, socialMedia := range SocialMedias {
		s := model.ResponseGetSocialMedia{}
		s.ID = socialMedia.ID
		s.Name = socialMedia.Name
		s.SocialMediaUrl = socialMedia.SocialMediaUrl
		s.UserID = socialMedia.UserID
		s.CreatedAt = socialMedia.CreatedAt
		s.UpdatedAt = socialMedia.UpdatedAt
		s.User.ID = socialMedia.User.ID
		s.User.Username = socialMedia.User.Username
		responseSocialMedias = append(responseSocialMedias, s)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": responseSocialMedias,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaIdParam, _ := strconv.Atoi(c.Param("socialMediaId"))
	contentType := c.Request.Header.Get("Content-Type")
	SocialMedia := model.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.ID = uint(socialMediaIdParam)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaIdParam).Updates(model.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_urk": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaIdParam, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := model.SocialMedia{}

	err := db.Where("id = ?", socialMediaIdParam).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been succesfully deleted",
	})
}
