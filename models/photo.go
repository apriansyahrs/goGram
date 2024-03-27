package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	Title     string     `json:"title" valid:"required~Judul harus diisi"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url" valid:"required~URL Photo harus diisi,url~URL Photo harus valid"`
	UserID    uint       `json:"user_id" valid:"required~ID User harus diisi"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      *User      `json:"user,omitempty"`
	Comments  []*Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}

type PhotoResponse struct {
	ID       uint               `json:"id"`
	Title    string             `json:"title"`
	Caption  string             `json:"caption"`
	PhotoURL string             `json:"photo_url"`
	UserID   uint               `json:"user_id"`
	User     UserResponseRelasi `json:"user"`
}

func (p *Photo) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}
