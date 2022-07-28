package models

type UserAssociation struct {
	Uuid           string `json:"uuid" gorm:"not null; size:256"`
	LastName       string `json:"last_name" gorm:"not null; size:256"`
	LastNameKana   string `json:"last_name_kana" gorm:"not null; size:256"`
	FirstName      string `json:"first_name" gorm:"not null; size:256"`
	FirstNameKana  string `json:"first_name_kana" gorm:"not null; size:256"`
	Nickname       string `json:"nickname" gorm:"not null; size:256"`
	Sex            string `json:"sex" gorm:"not null; size:256"`
	BirthYear      int    `json:"birth_year" gorm:"not null"`
	BirthMonth     int    `json:"birth_month" gorm:"not null"`
	BirthDay       int    `json:"birth_day" gorm:"not null"`
	AutoPermission bool   `json:"auto_permission" gorm:"not null; default:false"`
	IsExample      bool   `json:"is_example" gorm:"not null; default:false"`
	IsAdmin        bool   `json:"is_admin" gorm:"not null; default:false"`
	Statuses       []int  `json:"statuses" gorm:"many2many:user_statuses"`
}
