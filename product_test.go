package grappos

import (
	"testing"
)

func TestProduct(t *testing.T) {
	_, err := GetProducts("Ic-769313195")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBadProduct(t *testing.T) {
	_, err := GetProducts("IP-0122011238")
	if err == nil {
		t.Fatal("Should not get in here")
	}
}
