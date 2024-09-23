package foxpro

import (
	"github.com/SkiSale/cnswap-legacy-api/domain/dao"
	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

type VendorRepositoryFoxPro struct {
	table *dbase.File
}

func (r VendorRepositoryFoxPro) Count() int {
	return int(r.table.RowsCount())
}

func (r VendorRepositoryFoxPro) FindAllVendors(offset int, limit int) ([]*dao.Vendor, error) {
	return getSlice[dao.Vendor](r.table, offset, limit)
}

func (r VendorRepositoryFoxPro) FindVendorById(id int) (*dao.Vendor, error) {
	p := func(t *dbase.File) (*dbase.Field, error) {
		return t.NewFieldByName("VENDOR_ID", int32(id))
	}

	return findOneBy[dao.Vendor](r.table, p)
}

func VendorRepositoryInit(db *dbase.Database) *VendorRepositoryFoxPro {
	return &VendorRepositoryFoxPro{
		table: guardTable(db, "vendors"),
	}
}
