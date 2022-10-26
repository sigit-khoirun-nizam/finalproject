package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required"`
	UserID         uint   `json:"user_id"`
	User           *User  `json:"user"`
}

type ResponseGetSocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required"`
	UserID         uint   `json:"user_id"`
	User           struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if !govalidator.IsNotNull(s.Name) {
		err = errors.New("name required")
		return
	}

	if !govalidator.IsNotNull(s.SocialMediaUrl) {
		err = errors.New("social media url required")
		return
	}

	err = nil
	return
}
