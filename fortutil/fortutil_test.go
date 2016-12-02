package fortutil_test

import (
	"testing"

	"github.com/fortifi/potens-go/fortutil"
)

func TestCreateID(t *testing.T) {
	id := fortutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")
	if id != "kh-w-ekhwlkeh-fklwhelkfh-i-wjg" {
		t.Fail()
	}
}

func TestValidateID(t *testing.T) {
	if fortutil.ValidateID(fortutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")) != nil {
		t.Fail()
	}
	if fortutil.ValidateID(fortutil.CreateID("abcdef")) != nil {
		t.Fail()
	}
}

func TestValidateGlobalAppID(t *testing.T) {
	if fortutil.ValidateGlobalAppID("/asd") == nil {
		t.Fail()
	}
	if fortutil.ValidateGlobalAppID("asd/") == nil {
		t.Fail()
	}
	if fortutil.ValidateGlobalAppID("/a") == nil {
		t.Fail()
	}
	if fortutil.ValidateGlobalAppID("d") == nil {
		t.Fail()
	}
	if fortutil.ValidateGlobalAppID("dwfwe") == nil {
		t.Fail()
	}
	if fortutil.ValidateGlobalAppID("d/a") != nil {
		t.Fail()
	}
}

func TestRandomAlphaNum(t *testing.T) {
	if len(fortutil.RandomAlphaNum(1)) != 1 {
		t.Fail()
	}
	if len(fortutil.RandomAlphaNum(10)) != 10 {
		t.Fail()
	}
	if len(fortutil.RandomAlphaNum(100)) != 100 {
		t.Fail()
	}
	if len(fortutil.RandomAlphaNum(1000)) != 1000 {
		t.Fail()
	}
}

func TestAcronym(t *testing.T) {
	if fortutil.Acronym("A Really Long Acronym Goes Here", 3) != "ARL" {
		t.Error("Limited Long Acronym Failed")
	}
	if fortutil.Acronym("A Really Long Acronym Goes Here", 0) != "ARLAGH" {
		t.Error("Long Acronym Failed")
	}
	if fortutil.Acronym("Fortifi Technologies Ltd", 3) != "FTL" {
		t.Error("Length Matching Size Acronym Failed")
	}
	if fortutil.Acronym("Fortifi Technologies", 3) != "FTE" {
		t.Error("Extended End Acronym Failed")
	}
	if fortutil.Acronym("Fortifi Technologies", 0) != "FT" {
		t.Error("Short Acronym Failure")
	}
	if fortutil.Acronym("Fortifi", 3) != "FOR" {
		t.Error("Single Word Acronym with Length Failed")
	}
	if fortutil.Acronym("Fortifi", 0) != "F" {
		t.Error("Single Word Acronym Failed")
	}
}

func TestSplitGaID(t *testing.T) {
	vendor, app, err := fortutil.SplitGaID("fortifi/app")
	if err != nil {
		t.Error(err.Error())
	}
	if vendor != "fortifi" {
		t.Error("Incorrect vendor ID")
	}
	if app != "app" {
		t.Error("Incorrect app ID")
	}

	vendor, app, err = fortutil.SplitGaID("fortifi")
	if err == nil{
		t.Error("Failed to error on invalid GAID")
	}
	vendor, app, err = fortutil.SplitGaID("fortifi/klhfw/fwejhfew")
	if err == nil{
		t.Error("Failed to error on invalid GAID")
	}
}
