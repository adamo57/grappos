package Controller

import "testing"

func TestGetLocations(t *testing.T) {
	locations := [4]int{0, 1, 2, 3}
	invalidLocation := -1

	for _, location := range locations {
		_, err := GetLocations(location)

		if err != nil {
			t.Fatal("Error")
		}
	}

	_, err := GetLocations(invalidLocation)
	if err == nil {
		t.Fatal("There isn't an error, which is bad")
	}
}

func TestSearchForLocation(t *testing.T) {
	goodLocation := "11238"
	badLocation := "11"

	_, err := SearchForLocation(goodLocation)

	if err != nil {
		t.Fatal("Invalid location")
	}

	_, err = SearchForLocation(badLocation)
	if err == nil {
		t.Fatal("This should not work")
	}
}
