package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racroithenho/map-tracker/backend/internal/model"
	"github.com/racroithenho/map-tracker/backend/internal/service"
)

type DeviceHandler struct {
	services *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler{
	return &DeviceHandler{
		services: service,
	}
}

func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device, err := h.services.CreateDevice(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, device)
}

func (h *DeviceHandler) ReceiveLocation(c *gin.Context) {
	var device model.DeviceLocation

	if err := c.ShouldBindJSON(&device); err != nil {
		log.Println("Failed to bind device json in create function: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.SaveLocation(device); err != nil {
		log.Println("Failed to create device: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *DeviceHandler) GetLatestLocation(c *gin.Context) {
	deviceID := c.Param("device_id")

	device, err := h.services.GetLatestLocation(deviceID)
	if err != nil {
		log.Println("Failed to get latest device location: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if device == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no location found for this device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"device": device})
}

func (h *DeviceHandler) GetDeviceHistory(c *gin.Context) {
	deviceID := c.Param("device_id")
	fromTimeStr := c.Query("from_time")
	toTimeStr := c.Query("to_time")

	deviceLocation, err := h.services.GetDeviceHistory(deviceID, fromTimeStr, toTimeStr)
	if err != nil {
		log.Println("Failed to get device history: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"device_id": deviceID,
		"history":   deviceLocation,
	})
}

func (h *DeviceHandler) GetRoute(c *gin.Context) {
	startLatStr := c.Query("start_latitude")
	startLotStr := c.Query("start_longitude")
	endLatStr := c.Query("end_latitude")
	endLotStr := c.Query("end_longitude")

	response, err := h.services.GetRoute(startLatStr, startLotStr, endLatStr, endLotStr)
	if err != nil {
		log.Println("Failed to get route: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (h *DeviceHandler) GetAllDevices(c *gin.Context) {
	devices, err := h.services.GetAllDevices()
	if err != nil {
		log.Println("Failed to get all devices: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

func (h *DeviceHandler) GetAllDevicesLatestLocation(c *gin.Context) {
	devices, err := h.services.GetAllDevicesLatestLocation()
	if err != nil {
		log.Println("Failed to get all devices latest location: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"devices": devices})
}
