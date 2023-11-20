package models

type Users struct {
	Base
	Email    string `gorm:"size:128;not null;unique"`
	Password string `gorm:"size:130;not null;"`
}
