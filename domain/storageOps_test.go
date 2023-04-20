package domain

import (
	"testing"
)

func TestGetAttempts(t *testing.T) {
	type test struct {
		Name        string
		Uuid        string
		Attempts    int
		errorWanted bool
	}

	tests := []test{
		{
			"Pass",
			"86ae27ae-df99-11ed-bf70-2af8bc618b50",
			1,
			false,
		},
		{
			"Pass",
			"86ae27ae-df99-11ed-bf70-2af8bc618b51",
			5,
			false,
		},
		{
			"Fail",
			"86ae27ae-df99-11ed-bf70-2af8bc618b51",
			4,
			true,
		},
		{
			"Fail",
			"asd123asd123",
			1,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			att, err := Domainn.getAttempts(&RequestWithCode{
				RequestId: tt.Uuid,
				Code:      0,
			})

			if err != nil {
				if tt.errorWanted {
					return
				} else {
					t.Errorf("Got error: %v\n", err)
				}
			}

			if (att != tt.Attempts) != tt.errorWanted {
				t.Errorf("expected: %v, got: %v", tt.Attempts, att)
			}
		})
	}
}

func TestGetCode(t *testing.T) {
	type test struct {
		Name        string
		Uuid        string
		Code        int
		errorWanted bool
	}

	tests := []test{
		{
			"Pass",
			"86ae27ae-df99-11ed-bf70-2af8bc618b50",
			1234,
			false,
		},
		{
			"Fail",
			"86ae27ae-df99-11ed-bf70-2af8bc618b55",
			5,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			code, err := Domainn.getCode(&RequestWithCode{
				RequestId: tt.Uuid,
				Code:      0,
			})

			if err != nil {
				if tt.errorWanted {
					return
				} else {
					t.Errorf("Got error: %v\n", err)
					return
				}
			}

			if (code != tt.Code) != tt.errorWanted {
				t.Errorf("Expected: %v, got %v\n", tt.Code, code)
			}
		})
	}
}

func TestSaveGetCode(t *testing.T) {
	req := &RequestWithCode{
		RequestId: "99ae27ae-df99-11ed-bf70-2af8bc618b99",
		Code:      3210,
	}

	err := Domainn.saveCode(req)
	if err != nil {
		t.Errorf("Got error: %v\n", err)
		return
	}

	code, err := Domainn.getCode(req)
	if err != nil {
		t.Errorf("Got error: %v\n", err)
		return
	}

	if code != req.Code {
		t.Errorf("Expected: %v, Got: %v\n", req.Code, code)
	}
}
