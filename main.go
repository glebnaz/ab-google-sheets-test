package main

import (
	"fmt"

	"github.com/glebnaz/ab-google-sheets-test/sheets"
)

func main() {
	range_ := sheets.Range{
		Start: "A1",
		End:   "C3",
		List:  "",
	}
	table, err := sheets.InitSheet("15y-MClMIO6wtf6o2cj0lfa7Kn4GZmbok1yOBqWwDuDc")
	if err != nil {
		panic(err)
	}
	v, err := table.GetValues(range_)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
