package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/racroithenho/map-tracker/backend/internal/handler"
)

func RegisterDeviceRoutes(r *gin.Engine, deviceHandler *handler.DeviceHandler) {
	v1 := r.Group("/api/v1/devices")
	{
		v1.POST("/", deviceHandler.CreateDevice)
		v1.POST("/locations", deviceHandler.ReceiveLocation)
		v1.GET("/:device_id/latest", deviceHandler.GetLatestLocation)
		v1.GET("/:device_id/history", deviceHandler.GetDeviceHistory)
		v1.GET("/all_latest", deviceHandler.GetAllDevicesLatestLocation)
		v1.GET("/", deviceHandler.GetAllDevices)
		v1.GET("/route", deviceHandler.GetRoute)
	}
}