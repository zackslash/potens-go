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
