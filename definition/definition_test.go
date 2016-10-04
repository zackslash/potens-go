package definition_test

import (
	"testing"

	"github.com/fortifi/potens-go/definition"
	"github.com/fortifi/potens-go/l10n"
)

func TestReadYaml(t *testing.T) {
	def := definition.AppDefinition{}
	err := def.FromConfig("../app-definition.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if l10n.NewTranslatable(def.Name).Get("fr") != "Les clients" {
		t.Error("Failed to read translation")
	}

	if l10n.NewTranslatable(def.Name).Get("en") != "Customers" {
		t.Error("Failed to read translation")
	}

	if l10n.NewTranslatable(def.Name).Get("eeewf") != "Customers" {
		t.Error("Failed to read default")
	}
}
