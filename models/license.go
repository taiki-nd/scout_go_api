package models

import "time"

type License struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null; size:256"`
	GetYear   int       `json:"get_year" gorm:"not null"`
	GetMonth  int       `json:"get_month" gorm:"not null"`
	Public    bool      `json:"public" gorm:"not null; default:false"`
	Publish   bool      `json:"publish" gorm:"not null; default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
}
