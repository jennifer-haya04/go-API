package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id         int    `gorm:"primaryKey; autoIncrement" json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	DNI        string `gorm:"unique; not null" json:"dni"`
	Department string `json:"department"`
	Cel        string `json:"cel"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
