package Controller

import "testing"
import "net/http"

func TestLocation(t *testing.T) {
	res, err := http.Get("http://www.grappos.com/api2/locate.php?1=1&format=json&locate=1")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("Error fetching the data")
	}
}
