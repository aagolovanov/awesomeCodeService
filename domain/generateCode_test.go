package domain

import (
	"context"
	"strconv"
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

	type test struct {
		Name      string
		Number    string
		wantError bool
	}

	tests := []test{
		{
			"Pass",
			"+79998887766",
			false,
		},
		{
			"Fail",
			"+7123451",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := Domainn.GenerateCode(&RequestGenerate{
				Number: tt.Number,
			})

			if (err != nil) != tt.wantError {
				t.Errorf("Got error %v\n", err)
				return
			} else if err != nil {
				return
			}

			data, err := Domainn.Storage.GetAllData(context.Background(), result.RequestId)
			code, _ := strconv.Atoi(data["code"])
			if err != nil {
				t.Errorf("Got error %v\n", err)
				return
			}
			if code != result.Code {
				t.Errorf("Generated: %v In DB: %v", result.Code, code)
				return
			}
		})
	}

}
