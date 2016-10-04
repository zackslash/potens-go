package l10n_test

import (
	"testing"

	"github.com/fortifi/potens-go/l10n"
)

func TestEmpty(t *testing.T) {
	str := &l10n.Translatable{DefaultLanguage: "en", Key: "hello"}
	if str.Get("en") != "" {
		t.Error("Got a translation from never never lang")
	}
}

func TestNoDefault(t *testing.T) {
	str := &l10n.Translatable{Key: "hello"}
	str.Set("en", "Hi")
	if str.Get("fr") != "Hi" {
		t.Error("Unable to use EN default")
	}
}

func TestDefaults(t *testing.T) {
	str := &l10n.Translatable{DefaultLanguage: "en", Key: "hello"}
	str.SetDefaultLanguage("en")
	str.Set("en", "Hello")
	if "Hello" != str.Get("en") {
		t.Error("Unable to do basic lookup")
	}
	if "Hello" != str.Get("en-US") {
		t.Error("Unable to do parent lookup")
	}
	if "Hello" != str.Get("fr") {
		t.Error("Unable to do default lookup")
	}
	str.Set("en-US", "Hey")
	if "Hey" != str.Get("en-US") {
		t.Error("Specific language check failed")
	}
	if "Hello" != str.Get("en-GB") {
		t.Error("Specific language failover check failed")
	}

	str.Set("fr", "Bonjour")
	if "Bonjour" != str.Get("fr") {
		t.Error("Specific language check failed")
	}
}
