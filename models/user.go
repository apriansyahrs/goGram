package models

import (
	"errors"
	"goGram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	Username        string        `gorm:"not null; uniqueIndex" valid:"required~Nama harus diisi" json:"username"`
	Email           string        `gorm:"not null; uniqueIndex" valid:"required~Email harus diisi, email~Format email tidak valid" json:"email"`
	Password        string        `json:"password" gorm:"not null" valid:"required~Password harus diisi, minstringlength(6)~Password harus memiliki panjang minimal 6 karakter"`
	Age             int           `json:"age" gorm:"not null" valid:"required~Umur harus diisi, range(8|120)~Umur harus 8 atau lebih"`
	ProfileImageURL string        `json:"profile_image_url,omitempty" form:"profile_image_url" valid:"url~URL profil gambar tidak valid"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

type UserResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Age             int    `json:"age"`
	ProfileImageURL string `json:"profile_image_url,omitempty"`
}

type UserResponseRelasi struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err := u.Validate(); err != nil {
		return err
	}

	u.Password = helpers.HashPassword(u.Password)

	if u.ProfileImageURL != "" && !isValidURL(u.ProfileImageURL) {
		return errors.New("URL profil gambar tidak valid")
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := u.Validate(); err != nil {
		return err
	}

	return nil
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

func isValidURL(url string) bool {
	regex := govalidator.IsURL(url)
	return regex
}
