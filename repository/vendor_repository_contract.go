package repository

import "github.com/SkiSale/cnswap-legacy-api/domain/dao"

type VendorRepository interface {
	Count() int
	FindAllVendors(offset int, limit int) ([]*dao.Vendor, error)
	FindVendorById(id int) (*dao.Vendor, error)
}
