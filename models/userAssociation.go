package models

type UserAssociation struct {
	Uuid           string
	LastName       string
	LastNameKana   string
	FirstName      string
	FirstNameKana  string
	Nickname       string
	Sex            string
	BirthYear      int
	BirthMonth     int
	BirthDay       int
	AutoPermission bool
	IsExample      bool
	IsAdmin        bool
	Statuses       []int
}
