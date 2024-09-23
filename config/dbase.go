package config

import (
	"log"
	"os"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

func ConnectToDBF() *dbase.Database {
	var err error

	p := os.Getenv("DBF_PATH")

	log.Printf("Attempting to load FoxPro database from %s", p)

	db, err := dbase.OpenDatabase(&dbase.Config{
		Filename:   p,
		TrimSpaces: true,
	})

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("dbase.OpenDatabase returned nil. Does the database exist?")
	}

	return db
}
