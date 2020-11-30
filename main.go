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
	v, err := table.GetValues(range_)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
