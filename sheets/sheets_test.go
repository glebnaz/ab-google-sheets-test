package sheets

import (
	"os"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func initTest(t *testing.T) *assert.Assertions {
	return assert.New(t)
}

//TestInitSheets testing init function for tables
func TestInitSheets(t *testing.T) {
	assert := initTest(t)

	idSheet := os.Getenv(KeyID)
	if len(idSheet) == 0 {
		t.Fatal("Empty idSheet")
	}

	t.Run("Empty ID sheet", func(t *testing.T) {
		_, err := InitSheet("")
		assert.NotNil(err)
	})

	t.Run("Positive Init Test", func(t *testing.T) {
		_, err := InitSheet(idSheet)
		assert.Nil(err)
	})
}

func TestSheets(t *testing.T) {
	assert := initTest(t)

	idSheet := os.Getenv(KeyID)
	if len(idSheet) == 0 {
		t.Fatal("Empty idSheet")
	}

	diaposon := sheets.Range{
		Start: os.Getenv(KeyStart),
		End:   os.Getenv(KeyEnd),
		List:  os.Getenv(KeyList),
	}

	t.Run("Positive Sheets Test", func(t *testing.T) {
		table, err := InitSheet(idSheet)
		assert.Nil(err)
		table.GetValues()
		assert.Nil(err)
	})
}
