package grappos

import (
	"testing"
)

func TestProduct(t *testing.T) {
	_, err := GetProducts("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBadProduct(t *testing.T) {
	_, err := GetProducts("")
	if err == nil {
		t.Fatal("Should not get in here")
	}
}
