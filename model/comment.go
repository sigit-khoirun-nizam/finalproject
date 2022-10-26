package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required"`
	UserID  uint   `json:"user_id"`
	User    *User  `json:"user"`
	PhotoID uint   `json:"photo_id"`
	Photo   *Photo `json:"photo"`
}

type ResponseGetComment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required"`
	UserID  uint   `json:"user_id"`
	User    struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	PhotoID uint `json:"photo_id"`
	Photo   struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	}
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if !govalidator.IsNotNull(c.Message) {
		err = errors.New("message required")
		return
	}

	err = nil
	return
}
