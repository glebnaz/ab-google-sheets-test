package main

import (
	"fmt"
	"os"

	"github.com/glebnaz/ab-google-sheets-test/sheets"
)

func main() {
	tableID := os.Getenv(sheets.KeyID)
	range_ := sheets.Range{
		Start: "A1",
		End:   "C3",
		List:  "",
	}
	table, err := sheets.InitSheet(tableID)
	if err != nil {
		panic(err)
	}
	dataTable, err := table.GetValues(range_)
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
