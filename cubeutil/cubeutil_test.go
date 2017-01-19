package cubeutil_test

import (
	"testing"

	"github.com/cubex/potens-go/cubeutil"
)

func TestCreateID(t *testing.T) {
	id := cubeutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")
	if id != "kh-w-ekhwlkeh-fklwhelkfh-i-wjg" {
		t.Fail()
	}
}

func TestValidateID(t *testing.T) {
	if cubeutil.ValidateID(cubeutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")) != nil {
		t.Fail()
	}
	if cubeutil.ValidateID(cubeutil.CreateID("abcdef")) != nil {
		t.Fail()
	}
}

func TestValidateGlobalAppID(t *testing.T) {
	if cubeutil.ValidateGlobalAppID("/asd") == nil {
		t.Fail()
	}
	if cubeutil.ValidateGlobalAppID("asd/") == nil {
		t.Fail()
	}
	if cubeutil.ValidateGlobalAppID("/a") == nil {
		t.Fail()
	}
	if cubeutil.ValidateGlobalAppID("d") == nil {
		t.Fail()
	}
	if cubeutil.ValidateGlobalAppID("dwfwe") == nil {
		t.Fail()
	}
	if cubeutil.ValidateGlobalAppID("d/a") != nil {
		t.Fail()
	}
}

func TestRandomAlphaNum(t *testing.T) {
	if len(cubeutil.RandomAlphaNum(1)) != 1 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(10)) != 10 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(100)) != 100 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(1000)) != 1000 {
		t.Fail()
	}
}

func TestAcronym(t *testing.T) {
	if cubeutil.Acronym("A Really Long Acronym Goes Here", 3) != "ARL" {
		t.Error("Limited Long Acronym Failed")
	}
	if cubeutil.Acronym("A Really Long Acronym Goes Here", 0) != "ARLAGH" {
		t.Error("Long Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform System", 3) != "CPS" {
		t.Error("Length Matching Size Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform", 3) != "CPL" {
		t.Error("Extended End Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform", 0) != "CP" {
		t.Error("Short Acronym Failure")
	}
	if cubeutil.Acronym("Cubex", 3) != "CUB" {
		t.Error("Single Word Acronym with Length Failed")
	}
	if cubeutil.Acronym("Cubex", 0) != "C" {
		t.Error("Single Word Acronym Failed")
	}
}

func TestSplitGaID(t *testing.T) {
	vendor, app, err := cubeutil.SplitGaID("cubex/app")
	if err != nil {
		t.Error(err.Error())
	}
	if vendor != "cubex" {
		t.Error("Incorrect vendor ID")
	}
	if app != "app" {
		t.Error("Incorrect app ID")
	}

	vendor, app, err = cubeutil.SplitGaID("cubex")
	if err == nil {
		t.Error("Failed to error on invalid GAID")
	}
	vendor, app, err = cubeutil.SplitGaID("cubex/klhfw/fwejhfew")
	if err == nil {
		t.Error("Failed to error on invalid GAID")
	}
}
