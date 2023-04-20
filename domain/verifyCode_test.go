package domain

import (
	"testing"
)

func TestVerifyCode(t *testing.T) {
	type test struct {
		uuid        string
		code        int
		errorWanted bool
	}

	tests := []test{
		{
			"86ae27ae-df99-11ed-bf70-2af8bc618b50",
			1234,
			false,
		},
		{
			"86ae27ae-df99-11ed-bf70-2af8bc618b51",
			1234,
			true,
		},
		{
			"86ae27ae-df99-11ed-bf70-2af8bc618b50",
			1235,
			true,
		},
		{
			"86ae27ae-df99-11ed-bf70-2af8bc618b55",
			1235,
			true,
		},
		{
			"uuidBad",
			1111,
			true,
		},
	}

	for _, tt := range tests {
		l := len(tt.uuid) - 5
		t.Run(tt.uuid[l:], func(t *testing.T) {
			_, err := Domainn.VerifyCode(&RequestWithCode{
				RequestId: tt.uuid,
				Code:      tt.code,
			})

			if (err != nil) != tt.errorWanted {
				t.Errorf("uuid: %v, wanted: %v, got %v", tt.uuid, tt.errorWanted, err)
			}
		})
	}
}
