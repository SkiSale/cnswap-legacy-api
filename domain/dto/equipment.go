package dto

import "time"

type EquipmentTransaction struct {
	Type   string
	Values struct {
		Gross float64
		Tax1  float64
		Tax2  float64
	}
}
type EquipmentQuantities struct {
	Registered int
	CheckedIn  int
	Sold       int
}

type EquipmentFlags struct {
	IsNew    bool
	IsRedTag bool
}

type Equipment struct {
	Barcode       string
	VendorId      int32
	ItemTypeID    int32
	Description   string
	ItemSize      string
	Price         float64
	Flags         EquipmentFlags
	Stock         EquipmentQuantities
	TagsPaidCount int32
	DateSold      time.Time
	Transactions  []EquipmentTransaction
	Inventory     bool
	Notes         []string
}
