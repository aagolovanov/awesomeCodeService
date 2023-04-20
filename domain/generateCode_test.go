package domain

import (
	"testing"
)

func TestNumberVerify(t *testing.T) {
	trueNumbers := []string{
		"+79998887766", "+7-999-888-77-66", "+7 (999) 888-77-66",
	}

	falseNumbers := []string{
		"789", "123765", "1-1-1", "+982",
	}

	for _, n := range trueNumbers {
		if !verifyNumber(n) {
			t.Errorf("verifyNumber(%q) = false, want true", n)
		}
	}

	for _, n := range falseNumbers {
		if verifyNumber(n) {
			t.Errorf("verifyNumber(%q) = true, want false", n)
		}
	}
}

func TestGenerateCode(t *testing.T) {

	// FIXME mock storage and other
	d := Domain{
		Storage: nil,
		Logg:    nil,
		Config:  nil,
	}

	r, err := d.GenerateCode(&RequestGenerate{Number: "+7 (999) 888-77-66"})
	if err != nil {
		t.Error(err)
	}

	matchUuid := verifyUuid(r.RequestId)

	if !(1000 <= r.Code && r.Code <= 9999) {
		t.Errorf("Expected code between 1000;9999, actual %v", r.Code)
	}

	if !matchUuid {
		t.Errorf("Bad uuid: %v", r.RequestId)
	}
}
