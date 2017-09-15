package main

import (
	"fmt"
	"grappos_api/Controller"
)

func main() {
	fmt.Println(Controller.GetLocations(1))
	fmt.Println(Controller.SearchForLocation("11238"))
	fmt.Println(Controller.GetProducts("Ic-769313195"))
	fmt.Println(Controller.SearchCoordinates("43.048633", "-76.155551"))
	fmt.Println(Controller.SearchBrandByID(300, ""))
	fmt.Println(Controller.SearchProductByID(7411250, ""))
}
