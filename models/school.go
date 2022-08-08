package models

import "time"

type School struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null; size:256"`
	InYear    int       `json:"in_year" gorm:"not null"`
	InMonth   int       `json:"in_month" gorm:"not null"`
	OutYear   int       `json:"out_year" gorm:"not null"`
	OutMonth  int       `json:"out_month" gorm:"not null"`
	Public    bool      `json:"public" gorm:"not null; default:false"`
	Publish   bool      `json:"publish" gorm:"not null; default:false"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
