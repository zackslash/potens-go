package definition_test

import (
	"github.com/fortifi/potens-go/definition"
	"testing"
)

func TestReadYaml(t *testing.T) {
	def := definition.AppDefinition{}
	err := def.FromConfig("../app-definition.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
