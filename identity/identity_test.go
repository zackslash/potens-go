package identity_test

import (
	"testing"

	"github.com/fortifi/potens-go/identity"
)

func TestReadYaml(t *testing.T) {
	ident := identity.AppIdentity{}
	err := ident.FromJSONFile("../app-identity.json")
	if err != nil {
		t.Fatal(err)
	}
}
