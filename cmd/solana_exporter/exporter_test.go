package main

import (
	"reflect"
	"testing"
)

func TestLastEpochCredits(t *testing.T) {
	marker := []uint64{alpenglowMarker, alpenglowMarker, alpenglowMarker}
	tests := []struct {
		name    string
		credits [][]uint64
		want    []uint64
		wantOK  bool
	}{
		{"empty", nil, nil, false},
		{"plain", [][]uint64{{1063, 100, 90}}, []uint64{1063, 100, 90}, true},
		{"marker last", [][]uint64{{1063, 100, 90}, marker}, []uint64{1063, 100, 90}, true},
		{"marker then new epoch", [][]uint64{{1063, 100, 90}, marker, {1064, 200, 100}}, []uint64{1064, 200, 100}, true},
		{"marker only", [][]uint64{marker}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := lastEpochCredits(tt.credits)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("entry = %v, want %v", got, tt.want)
			}
		})
	}
}
