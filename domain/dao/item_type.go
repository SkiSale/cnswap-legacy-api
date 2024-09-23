package dao

type ItemType struct {
	ID           int32  `dbase:"ITYPEID"`
	Description  string `dbase:"ITDESCRIPT" validate:"required,length=15"`
	TaxableRate1 bool   `dbase:"TAXABLE1"`
	TaxableRate2 bool   `dbase:"TAXABLE2"`
}
