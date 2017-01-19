package definition_test

import (
	"testing"

	"github.com/cubex/potens-go/definition"
	"github.com/cubex/potens-go/i18n"
)

func TestReadYaml(t *testing.T) {
	def := definition.AppDefinition{}
	err := def.FromConfig("../app-definition.dist.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if i18n.NewTranslatable(def.Name).Get("fr") != "Les clients" {
		t.Error("Failed to read translation")
	}

	if i18n.NewTranslatable(def.Name).Get("en") != "Customers" {
		t.Error("Failed to read translation")
	}

	if i18n.NewTranslatable(def.Name).Get("eeewf") != "Customers" {
		t.Error("Failed to read default")
	}
}
