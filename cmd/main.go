package main

import (
	"flag"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/model"
	"github.com/sadihakan/dummy-dump/util"
	"log"
)

var (
	importArg  bool
	exportArg  bool
	sourceType string
	user       string
	path       string
	db         string
	binaryPath string
)

func main() {
	flag.BoolVar(&importArg, "import", false, "Import process")
	flag.BoolVar(&exportArg, "export", false, "Export process")
	flag.StringVar(&sourceType, "source", "", "Source type is: mysql|postgres")
	flag.StringVar(&user, "user", "", "User name")
	flag.StringVar(&path, "path", "", "Import file path")
	flag.StringVar(&db, "db", "", "Database name")
	flag.StringVar(&binaryPath, "binaryPath", "", "Binary path")
	flag.Parse()

	if !model.SOURCE_TYPE(sourceType).IsValid() {
		log.Println("invalid source type")
		return
	}

	if importArg && exportArg {
		log.Fatal("only one operation can be run")
	}

	var dump internal.Dump

	switch sourceType {
	case "postgres":
		dump = internal.Postgres{}
	case "mysql":
		dump = internal.MySQL{}
	case "mssql":
		dump = internal.MSSQL{}
	default:
		panic("")


	}

	if err := dump.Check(); err != nil {
		panic(err)
	}

	binaryPath = internal.CheckBinary(binaryPath, model.SOURCE_TYPE(sourceType), importArg, exportArg)

	if importArg {

		if !util.PathExists(path) {
			panic("Path is not exist")
		}

		if err := dump.Import(binaryPath, user, db, path); err != nil {
			panic(err)
		}
	}

	if exportArg {

		if err := dump.Export(binaryPath, user, db); err != nil {
			panic(err)
		}
	}
}
