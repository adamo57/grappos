package Controller

import (
	"net/http"
	"testing"
)

func TestProduct(t *testing.T) {
	res, err := http.Get("http://www.grappos.com/api2/subscriber.php?1=1&format=json&uid=Ic-769313195")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("Error fetching the data")
	}
}
