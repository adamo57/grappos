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
	storeTypes := []string{"All", "Wine Shops", "banana"}

	for _, v := range storeTypes  {
		_, err := SearchProductByID(1, v)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSearchBrandByID(t *testing.T) {
	storeTypes := []string{"All", "Wine Shops", "banana"}

	for _, v := range storeTypes  {
		_, err := SearchBrandByID(1, v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
