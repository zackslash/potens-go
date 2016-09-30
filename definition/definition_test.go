package definition_test

import (
	"testing"

	"github.com/fortifi/potens-go/definition"
)

func TestReadYaml(t *testing.T) {
	def := definition.AppDefinition{}
	err := def.FromConfig("../app-definition.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
