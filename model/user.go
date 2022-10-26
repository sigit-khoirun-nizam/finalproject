package model

import (
	"errors"
	"finalProject/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age <= 8 {
		errCreate = errors.New("age should be more than 8")
	}

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !govalidator.IsEmail(u.Email) {
		err = errors.New("email not valid")
		return
	}

	if !govalidator.IsNotNull(u.Username) {
		err = errors.New("username required")
		return
	}

	err = nil
	return
}
