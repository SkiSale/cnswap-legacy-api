package foxpro_test

import (
	"fmt"
	"testing"

	"github.com/SkiSale/cnswap-legacy-api/repository/foxpro"
	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

func connect() *dbase.Database {
	db, err := dbase.OpenDatabase(&dbase.Config{
		Filename:   "../_data/CnSData.dbc",
		TrimSpaces: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}
func TestVendorById(t *testing.T) {
	r := foxpro.VendorRepositoryInit(connect())

	var tests = []struct {
		id   int
		want string
	}{
		{3100, "Steve Webster"},
		{3870, "Skookum Cycle & Ski Revelstoke"},
	}

	for _, tcase := range tests {
		t.Run(fmt.Sprintf("%d,%s", tcase.id, tcase.want), func(t *testing.T) {
			v, err := r.FindVendorById(tcase.id)
			if err != nil {
				t.Error("got nil", err)
			}

			if v.LastName != tcase.want {
				t.Errorf("got %s wanted %s", v.LastName, tcase.want)
			}
		})
	}
}
