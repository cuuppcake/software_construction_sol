package main

import "vyshkat/internal/park"

func main() {
	container := park.NewContainer(park.RegularService{})
	park.RunDemo(container.Park)
}
