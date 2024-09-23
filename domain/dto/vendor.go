package dto

type VendorAddress struct {
	LineOne  string `json:"lineOne"`
	LineTwo  string `json:"lineTwo"`
	City     string `json:"city"`
	Province string `json:"province"`
	Postcode string `json:"postcode"`
}

type Vendor struct {
	ID                   int32         `json:"id"`
	VendorTypeID         int32         `json:"vendorTypeId"`
	LastName             string        `json:"lastName"`
	FirstName            string        `json:"firstName"`
	ContactName          string        `json:"contactName"`
	Address              VendorAddress `json:"address"`
	Phone                string        `json:"phone"`
	Email                string        `json:"email"`
	NotificationsEnabled bool          ``
	DeductTags           bool          ``
}
