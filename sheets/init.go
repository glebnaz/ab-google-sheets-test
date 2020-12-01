package sheets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

//InitSheet initializes a new table service by
//fetching access data from the "credentials.json" file
func InitSheet(id string) (Sheet, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("Empty id sheet")
	}

	client, err := serviceClient("credentials.json")
	if err != nil {
		return nil, err
	}
	service, err := sheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)
	}
	sheets := sheet{
		srv: service,
		id:  id,
	}
	return &sheets, nil
}

func serviceClient(credentialFile string) (*http.Client, error) {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		return &http.Client{}, err
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return &http.Client{}, err
	}
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			sheets.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(oauth2.NoContext)
	return client, nil
}
