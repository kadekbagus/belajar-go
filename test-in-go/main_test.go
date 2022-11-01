package main

import "testing"

var tests = []struct {
	name     string
	devidend float32
	divisor  float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 10.0, 10.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.devidend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but gone one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDevide(t *testing.T) {
	_, err := devide(10.0, 1.0)
	if err != nil {
		t.Error("got an error when we should not have")
	}
}

func TestBadDevide(t *testing.T) {
	_, err := devide(10.0, 0)
	if err == nil {
		t.Error("Dig not got an error when we should not have")
	}
}
