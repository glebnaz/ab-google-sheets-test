package sheets

import "fmt"

//Range is range for sheets
type Range struct {
	List  string
	Start string
	End   string
}

func (r Range) GetRangeString() (string, error) {
	if len(r.Start) == 0 || len(r.End) == 0 {
		return "", fmt.Errorf("Empty Start or End range field")
	}
	if len(r.List) == 0 {
		return fmt.Sprintf("%s:%s", r.Start, r.End), nil
	}
	return fmt.Sprintf("%s!%s:%s", r.List, r.Start, r.End), nil
}
