package sheets

import (
	gsheets "google.golang.org/api/sheets/v4"
)

//Sheet is table from google sheets
//this interface implemented
type Sheet interface {
	GetAllValues() (map[string]string, error)
	GetRangeValues(range_ Range) (map[string]string, error)
}

type sheet struct {
	//id of table, you can find id in url when editing table
	//https://docs.google.com/spreadsheets/d/15y-MClMIO6wtf6o2cj0lfa7Kn4GZmbok1yOBqWwDuDc/edit#gid=0
	//for example in this url id is: 15y-MClMIO6wtf6o2cj0lfa7Kn4GZmbok1yOBqWwDuDc
	id string
	//srv is Service from google lib api
	srv *gsheets.Service
}
