package config

import (
	"github.com/SkiSale/cnswap-legacy-api/controller"
	"github.com/SkiSale/cnswap-legacy-api/repository"
	"github.com/SkiSale/cnswap-legacy-api/service"
)

type Container struct {
	VendorRepository repository.VendorRepository
	VendorService    service.VendorService
	VendorController controller.VendorController

	EquipmentRepository repository.EquipmentRepository
	EquipmentService    service.EquipmentService
	EquipmentController controller.EquipmentController
}

func NewContainer(
	vendorService service.VendorService,
	vendorController controller.VendorController,
	vendorRepository repository.VendorRepository,
	equipmentService service.EquipmentService,
	equipmentController controller.EquipmentController,
	equipmentRepository repository.EquipmentRepository) *Container {
	return &Container{
		VendorRepository:    vendorRepository,
		VendorService:       vendorService,
		VendorController:    vendorController,
		EquipmentRepository: equipmentRepository,
		EquipmentService:    equipmentService,
		EquipmentController: equipmentController,
	}
}
