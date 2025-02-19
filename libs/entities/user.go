package entities

import "github.com/google/uuid"

type Users struct {
	UserId    uuid.UUID `gorm:"column:id; primary_key; type:varchar(36); default:gen_random_uuid()"`
	Firstname string    `gorm:"column:firstname; type:varchar(255); not null"`
	Lastname  string    `gorm:"column:lastname; type:varchar(255); not null"`
}

func (Users) TableName() string {
	return "users"
}
