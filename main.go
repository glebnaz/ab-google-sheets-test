package main

import (
	"fmt"
	"os"

	"github.com/glebnaz/ab-google-sheets-test/sheets"
)

func main() {
	tableID := os.Getenv(sheets.KeyID)
	diapason := sheets.Range{
		Start: os.Getenv(sheets.KeyStart),
		End:   os.Getenv(sheets.KeyEnd),
		List:  os.Getenv(sheets.KeyList),
	}
	table, err := sheets.InitSheet(tableID)
	if err != nil {
		panic(err)
	}
	dataTable, err := table.GetValues(diapason)
	if err != nil {
		panic(err)
	}
	for _, row := range dataTable {
		for _, col := range row {
			fmt.Printf(" %s ", col)
		}
		fmt.Printf("\n")
	}
}
