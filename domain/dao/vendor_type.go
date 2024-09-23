package dao

type VendorType struct {
	ID                int32   `dbase:"VENDTYPEID" validate:"required"`
	Description       string  `dbase:"DESCRIPTIO" validate:"length=20"`
	ItemCost          float64 `dbase:"ITEMCOST"`
	BulkItemCost      float64 `dbase:"BULKITEMCO"`
	CommisionRate     float64 `dbase:"COMMISSION"`
	ItemMin           float64 `dbase:"ITEM_MIN"`
	ItemMinCommission float64 `dbase:"ITEM_MIN_C"`
	RefundTax1        bool    `dbase:"REFUND_TAX"`
	RefundTax2        bool    `dbase:"REFUND_TA2"`
	ItemTax1          bool    `dbase:"ITEM_TAX1"`
	ItemTax2          bool    `dbase:"ITEM_TAX2"`
	ListItems         bool    `dbase:"LIST_ITEMS"`
}
