package main

import "testing"

type testSpec struct {
	input   string
	addr    string
	ar      string
	wantErr error
}

func TestGetSubnetRange(t *testing.T) {
	tests := []testSpec{
		{
			input:   "10.0.0.1/24",
			addr:    "10.0.0.1",
			ar:      "10.0.0.1 - 10.0.0.254",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		got, ar, err := Parse(tt.input)
		if err != tt.wantErr {
			t.Errorf("getSubnetRange(%s) error = %v, wantErr %v", tt.input, err, tt.wantErr)
		}
		if got != tt.addr {
			t.Errorf("getSubnetRange(%s) addr = %v, want %v", tt.input, got, tt.addr)
		}
		if ar != tt.ar {
			t.Errorf("getSubnetRange(%s) ar = %v, want %v", tt.input, ar, tt.ar)
		}
	}
}
