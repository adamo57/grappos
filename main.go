package main

import (
	"fmt"
	"hvíla/Controller"
)

func main() {
	fmt.Println(Controller.GetLocations())
	fmt.Println(Controller.SearchForLocation("11238"))
}
