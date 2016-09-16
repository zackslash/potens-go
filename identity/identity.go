package identity

import (
	"encoding/json"
	"io/ioutil"
)

// AppIdentity Auth data for your app
type AppIdentity struct {
	IdentityType string `json:"type"`
	IdentityID   string `json:"identity_id"`
	AppID        string `json:"service_name"`
	PrivateKey   string `json:"private_key"`
}

// FromJSONFile Populates your identity based on your app-identity.json
func (i *AppIdentity) FromJSONFile(jsonFile string) error {
	jsonContent, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonContent, i)
}
