package grappos

import (
	"testing"
)

func TestSearchCoordinates(t *testing.T) {
	_, err := SearchCoordinates("0.0", "0.0")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchProductByID(t *testing.T) {
	_, err := SearchProductByID(1, "banana")
	if err != nil {
		t.Fatal(err)
	}
}
