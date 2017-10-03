package grappos

import "testing"

func TestGetLocations(t *testing.T) {
	locations := [4]int{0, 1, 2, 3}
	invalidLocation := -1

	for _, location := range locations {
		_, err := GetLocations(location)

		if err != nil {
			t.Fatal(err)
		}
	}

	_, err := GetLocations(invalidLocation)
	if err == nil {
		t.Fatal("There isn't an error, which is bad")
	}
}

func TestSearchForLocationByPostcode(t *testing.T) {
	postcode := "11238"
	_, err := SearchForLocation(postcode)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchForLocationByCityName(t *testing.T) {
	goodLocation := "San Francisco"
	badLocation := "QWERTY-layout"

	_, err := SearchForLocation(goodLocation)

	if err != nil {
		t.Fatal(err)
	}

	_, err = SearchForLocation(badLocation)
	if err == nil {
		t.Fatal("This should not work")
	}
}
