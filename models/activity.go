package models

import "time"

type Activity struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null; size:256"`
	Content   string    `json:"content" gorm:"not null"`
	ImageUrl  string    `json:"image_url" gorm:"not null"`
	Url       string    `json:"url" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
