package fortutil_test

import (
	"github.com/fortifi/potens-go/fortutil"
	"testing"
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
}
