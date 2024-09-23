package controller

import (
	"github.com/SkiSale/cnswap-legacy-api/service"
	"github.com/gin-gonic/gin"
)

type EquipmentController interface {
	GetAllEquipment(c *gin.Context)
	GetEquipmentByBarcode(c *gin.Context)
	GetAllEquipmentByVendor(c *gin.Context)
}

type EquipmentControllerImpl struct {
	svc service.EquipmentService
}

func (u EquipmentControllerImpl) GetAllEquipment(c *gin.Context) {
	u.svc.GetAllEquipment(c)
}

func (u EquipmentControllerImpl) GetEquipmentByBarcode(c *gin.Context) {
	u.svc.GetEquipmentByBarcode(c)
}

func (u EquipmentControllerImpl) GetAllEquipmentByVendor(c *gin.Context) {
	u.svc.GetAllEquipmentByVendor(c)
}

func EquipmentControllerInit(EquipmentService service.EquipmentService) *EquipmentControllerImpl {
	return &EquipmentControllerImpl{
		svc: EquipmentService,
	}
}
