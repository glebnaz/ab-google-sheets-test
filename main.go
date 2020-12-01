package main

import (
	"fmt"

	"github.com/glebnaz/ab-google-sheets-test/sheets"
)

const tableID = ""

func main() {
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
