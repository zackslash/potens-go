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
