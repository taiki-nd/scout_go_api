package models

import "time"

type User struct {
	Id             uint         `json:"id" gorm:"primaryKey"`
	Uuid           string       `json:"uuid" gorm:"not null; size:256"`
	LastName       string       `json:"last_name" gorm:"not null; size:256"`
	LastNameKana   string       `json:"last_name_kana" gorm:"not null; size:256"`
	FirstName      string       `json:"first_name" gorm:"not null; size:256"`
	FirstNameKana  string       `json:"first_name_kana" gorm:"not null; size:256"`
	Nickname       string       `json:"nickname" gorm:"not null; size:256"`
	Sex            string       `json:"sex" gorm:"not null; size:256"`
	BirthYear      int          `json:"birth_year" gorm:"not null"`
	BirthMonth     int          `json:"birth_month" gorm:"not null"`
	BirthDay       int          `json:"birth_day" gorm:"not null"`
	AutoPermission bool         `json:"auto_permission" gorm:"not null; default:false"`
	IsExample      bool         `json:"is_example" gorm:"not null; default:false"`
	IsAdmin        bool         `json:"is_admin" gorm:"not null; default:false"`
	Statuses       []Status     `json:"statuses" gorm:"many2many:user_statuses"`
	Prefectures    []Prefecture `json:"prefectures" gorm:"many2many:user_prefectures"`
	Licenses       []License    `json:"licenses" gorm:"foreignKey:UserId"`
	Schools        []School     `json:"schools" gorm:"foreignKey:UserId"`
	Activities     []Activity   `json:"activities" gorm:"foreignKey:UserId"`
	Works          []Work       `json:"works" gorm:"foreignKey:UserId"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at`
}
