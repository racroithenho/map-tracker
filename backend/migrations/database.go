package migrations

import (
	"github.com/racroithenho/map-tracker/backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func Connect() (*gorm.DB, error){
	dsn := "host=localhost user=postgres password=racroithenho dbname=map port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto migrate
    if err := database.AutoMigrate(&model.Device{}, &model.DeviceLocation{}); err != nil {
        return nil, err
    }

    return database, nil
}