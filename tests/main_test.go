package main

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid_data", 100, 10, 10, false},
	{"invalid_data", 100, 0, 0, true},
	{"expect-5", 50, 10, 5, false},	
	{"expect-fraction", -1, -777, 0.0012870013, false},	

}

func TestDivision(t *testing.T){
	for _,tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got )
		}
	}
}
