package controller

import (
	"github.com/SkiSale/cnswap-legacy-api/service"
	"github.com/gin-gonic/gin"
)

type VendorController interface {
	GetAllVendors(c *gin.Context)
	GetVendorById(c *gin.Context)
}

type VendorControllerImpl struct {
	svc service.VendorService
}

func (u VendorControllerImpl) GetAllVendors(c *gin.Context) {
	u.svc.GetAllVendors(c)
}

func (u VendorControllerImpl) GetVendorById(c *gin.Context) {
	u.svc.GetVendorById(c)
}

func VendorControllerInit(VendorService service.VendorService) *VendorControllerImpl {
	return &VendorControllerImpl{
		svc: VendorService,
	}
}
