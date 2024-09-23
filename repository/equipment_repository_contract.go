package repository

import "github.com/SkiSale/cnswap-legacy-api/domain/dao"

type EquipmentRepository interface {
	Count() int
	FindAllEquipment(offset int, limit int) ([]*dao.Equipment, error)
	FindEquipmentByBarcode(barcode string) (*dao.Equipment, error)
	FindEquipmentByVendor(vendorId int) ([]*dao.Equipment, error)
}
