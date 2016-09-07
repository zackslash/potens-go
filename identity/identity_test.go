package identity_test

import (
	"github.com/fortifi/potens-go/identity"
	"testing"
)

func TestReadYaml(t *testing.T) {
	ident := identity.AppIdentity{}
	err := ident.FromJsonFile("../app-identity.json")
	if err != nil {
		t.Fatal(err)
	}
}
