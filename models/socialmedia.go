package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `json:"name" valid:"required~Nama sosial media harus diisi"`
	SocialMediaURL string    `json:"social_media_url" valid:"required~URL media sosial harus diisi,url~URL sosial media harus valid"`
	UserID         uint      `json:"user_id" valid:"required~ID User harus diisi"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           *User
}

type SocialMediaResponse struct {
	ID             uint               `json:"id"`
	Name           string             `json:"name"`
	SocialMediaURL string             `json:"social_media_url"`
	UserID         uint               `json:"user_id"`
	User           UserResponseRelasi `json:"user"`
}

func (s *SocialMedia) Validate() error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}
	if !isValidURL(s.SocialMediaURL) {
		return errors.New("URL media sosial tidak valid")
	}
	return nil
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	return s.Validate()
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	return s.Validate()
}
