package foxpro

import (
	"errors"
	"fmt"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

func marshallRow[M any](r *dbase.Row) (*M, error) {
	m := new(M)
	if err := r.ToStruct(m); err != nil {
		return nil, err
	}
	return m, nil
}

func guardTable(db *dbase.Database, tn string) *dbase.File {
	t, ok := db.Tables()[tn]
	if !ok {
		panic(fmt.Sprintf("Could not load table %s from db", tn))
	}
	return t
}

type predicateFactory func(table *dbase.File) (*dbase.Field, error)

func findOneBy[M any](t *dbase.File, p predicateFactory) (*M, error) {
	f, err := p(t)
	if err != nil {
		return nil, err
	}

	records, err := t.Search(f, false)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	return marshallRow[M](records[0])
}

func findBy[M any](t *dbase.File, p predicateFactory) ([]*M, error) {
	f, err := p(t)
	if err != nil {
		return nil, err
	}

	records, err := t.Search(f, false)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	slice := make([]*M, 0)
	for _, row := range records {
		m, err := marshallRow[M](row)
		if err != nil {
			continue
		}
		slice = append(slice, m)
	}
	return slice, nil
}

func getSlice[M any](t *dbase.File, offset int, length int) ([]*M, error) {
	slice := make([]*M, 0)

	o := uint32(offset)
	l := uint32(length)

	if uint32(offset) > t.RowsCount() {
		return nil, errors.New("offset out of range")
	}

	t.GoTo(o)
	for !t.EOF() {
		row, err := t.Next()
		if err != nil {
			continue
		}

		if row.Position > (o + l - 1) {
			break
		}

		m, err := marshallRow[M](row)
		if err != nil {
			continue
		}
		slice = append(slice, m)
	}

	return slice, nil
}
