package sheets

import (
	"fmt"

	gsheets "google.golang.org/api/sheets/v4"
)

//Sheet is table from google sheets
//this interface implemented
type Sheet interface {
	GetValues(range_ Range) ([][]string, error)
}

type sheet struct {
	//id of table, you can find id in url when editing table
	//https://docs.google.com/spreadsheets/d/15y-MClMIO6wtf6o2cj0lfa7Kn4GZmbok1yOBqWwDuDc/edit#gid=0
	//for example in this url id is: 15y-MClMIO6wtf6o2cj0lfa7Kn4GZmbok1yOBqWwDuDc
	id string
	//srv is Service from google lib api
	srv *gsheets.Service
}

//GetValues retrun data from table by [][]string
//range is Range type to get Values
func (s *sheet) GetValues(range_ Range) ([][]string, error) {
	var data [][]string
	r, err := range_.GetRangeString()
	if err != nil {
		return nil, err
	}

	rsp, err := s.srv.Spreadsheets.Values.Get(s.id, r).Do()
	if err != nil {
		return nil, err
	}
	if len(rsp.Values) == 0 {
		return nil, fmt.Errorf("Empty Values")
	}

	for _, row := range rsp.Values {
		var d []string
		for _, col := range row {
			d = append(d, fmt.Sprintf("%v", col))
		}
		data = append(data, d)
	}
	return data, nil
}
