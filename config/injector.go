//go:build wireinject
// +build wireinject

package config

import (
	"github.com/SkiSale/cnswap-legacy-api/controller"
	"github.com/SkiSale/cnswap-legacy-api/repository"
	"github.com/SkiSale/cnswap-legacy-api/repository/foxpro"
	"github.com/SkiSale/cnswap-legacy-api/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDBF)

var vendorServiceSet = wire.NewSet(service.VendorServiceInit,
	wire.Bind(new(service.VendorService), new(*service.VendorServiceImpl)),
)

var vendorRepositorySet = wire.NewSet(foxpro.VendorRepositoryInit,
	wire.Bind(new(repository.VendorRepository), new(*foxpro.VendorRepositoryFoxPro)),
)

var vendorControllerSet = wire.NewSet(controller.VendorControllerInit,
	wire.Bind(new(controller.VendorController), new(*controller.VendorControllerImpl)),
)

var equipmentServiceSet = wire.NewSet(service.EquipmentServiceInit,
	wire.Bind(new(service.EquipmentService), new(*service.EquipmentServiceImpl)),
)

var equipmentRepositorySet = wire.NewSet(foxpro.EquipmentRepositoryInit,
	wire.Bind(new(repository.EquipmentRepository), new(*foxpro.EquipmentRepositoryFoxPro)),
)

var equipmentControllerSet = wire.NewSet(controller.EquipmentControllerInit,
	wire.Bind(new(controller.EquipmentController), new(*controller.EquipmentControllerImpl)),
)

func Init() *Container {
	wire.Build(NewContainer, db, vendorControllerSet, vendorServiceSet, vendorRepositorySet, equipmentServiceSet, equipmentRepositorySet, equipmentControllerSet)
	return nil
}
