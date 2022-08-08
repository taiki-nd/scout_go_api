package models

import "time"

type Resume struct {
	Id                  uint      `json:"id" gorm:"primaryKey"`
	ImadeUrl            string    `json:"image_url" gorm:"not null; size:256"`
	Prefecture          string    `json:"prefecture" gorm:"not null; size:256"`
	Address             string    `json:"address" gorm:"not null; size:256"`
	PostalCode          string    `json:"postal_code" gorm:"not null; size:256"`
	Tell                string    `json:"tell" gorm:"not null; size:256"`
	Email               string    `json:"email" gorm:"not null; size:256"`
	Line                string    `json:"line" gorm:"not null; size:256"`
	Station             string    `json:"station" gorm:"not null; size:256"`
	Family              int       `json:"family" gorm:"not null; size:256"`
	IsMarried           bool      `json:"is_married" gorm:"not null; default:false"`
	IsMarriedObligation bool      `json:"is_married_obligation" gorm:"not null; default:false"`
	IsMove              bool      `json:"is_ move" gorm:"not null; default:false"`
	DesiredSalary       int       `json:"desired_salary" gorm:"not null"`
	PreviousSalary      int       `json:"previous_salary" gorm:"not null"`
	OtherInfoTitle1     string    `json:"other_info_Title1"`
	OtherInfoUrl1       string    `json:"other_info_Url1"`
	OtherInfoTitle2     string    `json:"other_info_Title2"`
	OtherInfoUrl2       string    `json:"other_info_Url2"`
	OtherInfoTitle3     string    `json:"other_info_Title3"`
	OtherInfoUrl3       string    `json:"other_info_Url3"`
	OtherInfoTitle4     string    `json:"other_info_Title4"`
	OtherInfoUrl4       string    `json:"other_info_Url4"`
	OtherInfoTitle5     string    `json:"other_info_Title5"`
	OtherInfoUrl5       string    `json:"other_info_Url5"`
	WorkArea            string    `json:"work_area" gorm:"not null	; size:256"`
	Pr                  string    `json:"pr" gorm:"not null"`
	UserId              uint      `json:"user_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
