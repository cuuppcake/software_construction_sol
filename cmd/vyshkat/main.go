package main

import (
	"vyshkat/internal/app"
	"vyshkat/internal/park"
)

func main() {
	container := app.NewContainer(park.RegularService{})
	app.RunDemo(container.Park)
}
