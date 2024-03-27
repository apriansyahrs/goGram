package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"user_id" valid:"required~ID User harus diisi"`
	PhotoID   uint      `json:"photo_id" valid:"required~ID Photo harus diisi"`
	Message   string    `json:"message" valid:"required~Pesan harus diisi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user,omitempty"`
	Photo     *Photo    `json:"photo,omitempty"`
}

func (c *Comment) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}
