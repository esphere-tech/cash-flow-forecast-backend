package models

import (
	"github.com/google/uuid"
)

type CashEntry struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Type        string    `gorm:"type:text;not null;check:type IN ('inflow', 'outflow')"`
	Amount      float64   `gorm:"not null"`
	Category    string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
	Date        string    `gorm:"type:date;not null"`
	CreatedAt   int64     `gorm:"autoCreateTime"`
}
