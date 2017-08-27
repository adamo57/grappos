package main

import (
	"fmt"
	"hvíla/Controller"
)

func main() {
	fmt.Println(Controller.GetLocations(1))
	fmt.Println(Controller.SearchForLocation("11238"))
	fmt.Println(Controller.GetProducts(18))
}
