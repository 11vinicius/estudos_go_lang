package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time `sql:"index"`
}
