package utils

import (
	"api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSqlite() {

	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("DATABASE_CONNECTION_ERROR")
	}
	migrate()
}

func migrate() {
	err := DB.AutoMigrate(
		&models.Cattle{},
		&models.BirthRecord{},
		&models.User{},
		&models.DeathRecord{},
		&models.WeightRecord{},
		&models.MilkingRecord{},
		&models.IllnessRecord{},
		&models.InseminationRecord{},
	)
	if err != nil {
		panic("DATABASE_MIGRATION_ERROR")
	}
}
