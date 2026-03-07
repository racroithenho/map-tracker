package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/racroithenho/map-tracker/backend/internal/model"
	"github.com/racroithenho/map-tracker/backend/internal/repository"
)

type DeviceService struct{
	repository *repository.DeviceRepository
}

func NewDeviceService(repo *repository.DeviceRepository) *DeviceService {
	return &DeviceService{
		repository: repo,
	}
}

func (s *DeviceService) CreateDevice(name string) (*model.Device, error) {
	if name == "" {
		return nil, errors.New("Invalid name device")
	}

	device, err := s.repository.CreateDevice(name)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (s *DeviceService) SaveLocation(device model.DeviceLocation) error {
	if device.DeviceID == "" {
		return errors.New("Select a device to save")
	}

	if device.Latitude == 0 || device.Longitude == 0 {
		return errors.New("Location invalid")
	}

	exists, err := s.repository.DeviceExists(device.DeviceID)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("device not found")
    }

	save := &model.DeviceLocation{
		DeviceID: device.DeviceID,
		Latitude: device.Latitude,
		Longitude: device.Longitude,
	}

	if err := s.repository.SaveLocation(save); err != nil {
		return err
	}
return nil
}

func (s *DeviceService) GetLatestLocation(deviceID string) (*model.DeviceLocation, error) {
	if deviceID == "" {
		return nil, errors.New("device_id invalid")
	}

	response, err := s.repository.GetLatestLocation(deviceID)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DeviceService) GetDeviceHistory(deviceID, fromTimeStr, toTimeStr string) ([]model.DeviceLocation, error) {
	if deviceID == "" {
		return nil, errors.New("device_id invalid")
	}

	if fromTimeStr == "" || toTimeStr == "" {
		return nil, errors.New("time selected invalid")
	}

	fromTime, err := time.Parse(time.RFC3339, fromTimeStr)
	if err != nil {
		return nil, err
	}
	
	toTime, err := time.Parse(time.RFC3339, toTimeStr)
	if err != nil {
		return nil, err
	}

	deviceHistory, err := s.repository.GetDeviceHistory(deviceID, fromTime, toTime)
	if err != nil {
		return nil, err
	}

	return deviceHistory, nil
}

func (s *DeviceService) GetRoute(startLatStr, startLotStr, endLatStr, endLotStr string) (any, error) {
	if startLatStr == "" || startLotStr == "" || endLatStr == "" || endLotStr == "" {
		return nil, errors.New("invalid location")
	}

	startLat, err := strconv.ParseFloat(startLatStr, 64)
	if err != nil {
		return nil, err
	}

	startLot, err := strconv.ParseFloat(startLotStr, 64)
	if err != nil {
		return nil, err
	}

	endLat, err := strconv.ParseFloat(endLatStr, 64)
	if err != nil {
		return nil, err
	}

	endLot, err := strconv.ParseFloat(endLotStr, 64)
	if err != nil {
		return nil, err
	}

	response, err := s.repository.GetRoute(startLat, startLot, endLat, endLot)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DeviceService) GetAllDevices() ([]model.Device, error) {
	devices, err := s.repository.GetAllDevices()
	if err != nil {
		return nil, err
	}

	return devices,  nil
}

func (s *DeviceService) GetAllDevicesLatestLocation() ([]model.DeviceLocation, error) {
	devices, err := s.repository.GetAllDevicesLatestLocation()
	if err != nil {
		return nil, err
	}

	return devices,  nil
}