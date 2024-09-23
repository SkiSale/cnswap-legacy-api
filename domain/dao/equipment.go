package dao

import "time"

type Equipment struct {
	Barcode       string    `dbase:"BARCODE" validate:"required"`
	Owner         int32     `dbase:"OWNER" validate:"required"`
	ItemTypeID    int32     `dbase:"ITYPE" validate:"required"`
	Description   string    `dbase:"DESCRIPTIO" validate:"length=30"`
	ItemSize      string    `dbase:"ITEMSIZE" validate:"length=5"`
	Price         float64   `dbase:"PRICE"`
	IsNew         bool      `dbase:"NEW"`
	IsRedTag      bool      `dbase:"REDTAG"`
	QtyIn         int32     `dbase:"QTY_IN"`
	QtyCheckedIn  int32     `dbase:"QTY_CHECKE"`
	QtySold       int32     `dbase:"QTY_SOLD"`
	TagsPaidCount int32     `dbase:"TAGS_PAID"`
	DateSold      time.Time `dbase:"DATE_SOLD"`
	ValueSold     float64   `dbase:"VALUE_SOLD"`
	ValueTax1     float64   `dbase:"VALUE_TAX1"`
	ValueTax2     float64   `dbase:"VALUE_TAX2"`
	PaidValue     float64   `dbase:"PAID_VALUE"`
	PaidTax1      float64   `dbase:"PAID_TAX1"`
	PaidTax2      float64   `dbase:"PAID_TAX2"`
	Inventory     bool      `dbase:"INVENTORY"`
	Notes         string    `dbase:"NOTES"`
}
