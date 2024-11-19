package models

import (
	"time"

	"github.com/google/uuid"
)

type SearchLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Query     string    `gorm:"type:varchar(255);not null"`
	Results   int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
