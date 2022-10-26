package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `json:"title" form:"title" valid:"required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint      `json:"user_id"`
	User     *User     `json:"user"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
}

type ResponseGetPhoto struct {
	GormModel
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	User     struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if !govalidator.IsNotNull(p.Title) {
		err = errors.New("title required")
		return
	}

	if !govalidator.IsNotNull(p.PhotoUrl) {
		err = errors.New("photo url required")
		return
	}

	err = nil
	return
}
