package models

import "time"

type Work struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null; size:256"`
	InYear          int       `json:"in_year" gorm:"not null"`
	InMonth         int       `json:"in_month" gorm:"not null"`
	OutYear         int       `json:"out_year" gorm:"not null"`
	OutMonth        int       `json:"out_month" gorm:"not null"`
	BusinessContent string    `json:"business_content" gorm:"not null; size:256"`
	EmployeeCount   int       `json:"employee_count" gorm:"not null; size:256"`
	EmployStyle     string    `json:"employ_style" gorm:"not null; size:256"`
	WorkName        string    `json:"work_name" gorm:"not null; size:256"`
	WorkContent     string    `json:"work_content" gorm:"not null; size:256"`
	Public          bool      `json:"public" gorm:"not null; default:false"`
	Publish         bool      `json:"publish" gorm:"not null; default:false"`
	UserId          uint      `json:"user_id"`
	Projects        []Project `json:"projects" gorm:"foreignKey:WorkId"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at`
}
