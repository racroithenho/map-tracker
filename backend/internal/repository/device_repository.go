package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/racroithenho/map-tracker/backend/internal/model"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

func (r *DeviceRepository) CreateDevice(name string) (*model.Device, error) {
	device := model.Device{
		DeviceID: uuid.NewString(),
		Name:     name,
	}

	err := r.db.Create(&device).Error
	return &device, err
}

func (r *DeviceRepository) SaveLocation(device *model.DeviceLocation) error {
	device.Timestamp = time.Now()
	return r.db.Create(device).Error
}

func (r *DeviceRepository) GetLatestLocation(deviceID string) (*model.DeviceLocation, error) {
	var location *model.DeviceLocation

	err := r.db.Where("device_id = ?", deviceID).Order("timestamp DESC").First(&location).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return location, nil
}

func (r *DeviceRepository) GetDeviceHistory(deviceID string, fromTime, toTime time.Time) ([]model.DeviceLocation, error) {
	var history []model.DeviceLocation

	err := r.db.
		Where("device_id = ? AND timestamp BETWEEN ? AND ?", deviceID, fromTime, toTime).
		Order("timestamp ASC").
		Find(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *DeviceRepository) GetRoute(startLat, startLon, endLat, endLon float64) (*model.OSRMResponse, error) {

	url := fmt.Sprintf(
		"http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		startLon, startLat, endLon, endLat,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("osrm api error: %s", resp.Status)
	}

	var result model.OSRMResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Routes) == 0 {
		return nil, errors.New("no route found")
	}

	return &result, nil
}

func (r *DeviceRepository) GetAllDevices() ([]model.Device, error) {
    var devices []model.Device
    err := r.db.Find(&devices).Error
    return devices, err
}


func (r *DeviceRepository) GetAllDevicesLatestLocation() ([]model.DeviceLocation, error) {
	var devices []model.DeviceLocation

	err := r.db.Raw(`
					SELECT DISTINCT ON (device_id) *
					FROM device_locations
					ORDER BY device_id, timestamp DESC
					`).Scan(&devices).Error
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *DeviceRepository) DeviceExists(deviceID string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Device{}).Where("device_id = ?", deviceID).Count(&count).Error
	return count > 0, err
}