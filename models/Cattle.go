package models

import (
	"time"

	"gorm.io/gorm"
)

type Cattle struct {
	UUID                      string         `gorm:"primaryKey;unique"`
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	Birthday                  time.Time
	LastInseminationDate      time.Time `gorm:"default:null"`
	LastGiveBirthDate         time.Time `gorm:"default:null"`
	Gender                    string    `gorm:"index"`
	Period                    string    `gorm:"-"`
	Classification            string    `gorm:"-"`
	TagNumber                 string    `gorm:"primaryKey;index"`
	DeadTagNumber             string
	MotherUUID                string `gorm:"index"`
	FatherUUID                string `gorm:"index"`
	OwnerUUID                 string `gorm:"index"`
	InseminationRecordUUID    string
	ChildrenCount             uint
	PregnancyStatus           string              `gorm:"index"` // not-pregnant | inseminated | pregnant
	IsAlive                   bool                `gorm:"default:true;index"`
	IsCastrated               bool                `gorm:"default:false;index"`
	Mother                    *Cattle             `gorm:"foreignKey:MotherUUID;references:UUID"`
	Father                    *Cattle             `gorm:"foreignKey:FatherUUID;references:UUID"`
	Owner                     *User               `gorm:"foreignKey:OwnerUUID;references:UUID"`
	CurrentInseminationRecord *InseminationRecord `gorm:"foreignKey:InseminationRecordUUID;references:UUID"`
	WeightRecords             []WeightRecord
	MilkingRecords            []MilkingRecord
	IllnessRecords            []IllnessRecord
	InseminationRecords       []InseminationRecord `gorm:"foreignKey:MotherUUID;references:UUID"`
}
