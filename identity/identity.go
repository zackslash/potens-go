package identity

import (
	"encoding/json"
	"io/ioutil"
)

type AppIdentity struct {
	IdentityType string `json:"type"`
	IdentityID   string `json:"identity_id"`
	AppID        string `json:"service_name"`
	PrivateKey   string `json:"private_key"`
}

func (i *AppIdentity) FromJsonFile(jsonFile string) error {
	jsonContent, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonContent, i)
}
