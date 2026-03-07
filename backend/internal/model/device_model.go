package model

import (
	"time"
)

type Device struct {
	DeviceID  string           `gorm:"primaryKey" json:"device_id"`
	Name      string           `json:"name"`
	Locations []DeviceLocation `gorm:"foreignKey:DeviceID"`
}

type DeviceLocation struct {
	ID        uint      `gorm:"primaryKey"`
	DeviceID  string    `gorm:"not null;index" json:"device_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}
