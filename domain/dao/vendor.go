package dao

type Vendor struct {
	ID                   int32  `dbase:"VENDOR_ID" validate:"required"`
	VendorTypeID         int32  `dbase:"VENDTYPE" validate:"required"`
	LastName             string `dbase:"LASTNAME" validate:"length=30"`
	FirstName            string `dbase:"FIRSTNAME" validate:"length=30"`
	AddressLineOne       string `dbase:"ADDRESS1" validate:"length=30"`
	AddressLineTwo       string `dbase:"ADDRESS2" validate:"length=30"`
	AddressCity          string `dbase:"CITY" validate:"length=25"`
	AddressProvince      string `dbase:"PROV" validate:"length=3"`
	AddressPostcode      string `dbase:"POSTAL" validate:"length=7"`
	ContactName          string `dbase:"CONTACT" validate:"length=40"`
	Phone                string `dbase:"PHONE" validate:"length=18"`
	Email                string `dbase:"EMAIL" validate:"length=40"`
	NotificationsEnabled bool   `dbase:"EMAIL_NOTI"`
	DeductTags           bool   `dbase:"DEDUCT_TAG"`
}
