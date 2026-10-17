package app

import "vyshkat/internal/park"

type Container struct {
	Park *park.Park
}

func NewContainer(service park.ServiceCenter) *Container {
	return &Container{Park: park.NewPark(service)}
}
