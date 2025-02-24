package models

import (
	"time"

	"gorm.io/gorm"
)

type InseminationRecord struct {
	UUID             string `gorm:"primaryKey;unique"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Date             time.Time
	MotherUUID       string  `gorm:"index"`
	FatherUUID       string  `gorm:"index"`
	Mother           *Cattle `gorm:"foreignKey:MotherUUID;references:UUID"`
	Father           *Cattle `gorm:"foreignKey:FatherUUID;references:UUID"`
	Type             string  // natural | artificial
	Status           string  // uncertain | pregnant | done | failed
	StatusUpdateDate time.Time
}
