package foxpro

import (
	"github.com/SkiSale/cnswap-legacy-api/domain/dao"
	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

type EquipmentRepositoryFoxPro struct {
	table *dbase.File
}

func (r EquipmentRepositoryFoxPro) Count() int {
	return int(r.table.RowsCount())
}

func (r EquipmentRepositoryFoxPro) FindAllEquipment(offset int, limit int) ([]*dao.Equipment, error) {
	return getSlice[dao.Equipment](r.table, offset, limit)
}

func (r EquipmentRepositoryFoxPro) FindEquipmentByBarcode(barcode string) (*dao.Equipment, error) {
	p := func(t *dbase.File) (*dbase.Field, error) {
		return t.NewFieldByName("BARCODE", barcode)
	}

	return findOneBy[dao.Equipment](r.table, p)
}

func (r EquipmentRepositoryFoxPro) FindEquipmentByVendor(vendorId int) ([]*dao.Equipment, error) {
	p := func(t *dbase.File) (*dbase.Field, error) {
		return t.NewFieldByName("OWNER", int32(vendorId))
	}

	return findBy[dao.Equipment](r.table, p)
}

func EquipmentRepositoryInit(db *dbase.Database) *EquipmentRepositoryFoxPro {
	return &EquipmentRepositoryFoxPro{
		table: guardTable(db, "equipment"),
	}
}
