package router

import (
	"github.com/SkiSale/cnswap-legacy-api/config"
	"github.com/gin-gonic/gin"
)

func Init(ioc *config.Container) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	vendor := router.Group("/vendor")
	vendor.GET("", ioc.VendorController.GetAllVendors)
	vendor.GET("/:vendorId", ioc.VendorController.GetVendorById)
	vendor.GET("/:vendorId/equipment", ioc.EquipmentController.GetAllEquipmentByVendor)

	equipment := router.Group("/equipment")
	equipment.GET("", ioc.EquipmentController.GetAllEquipment)
	equipment.GET("/:barcode", ioc.EquipmentController.GetEquipmentByBarcode)

	return router
}
