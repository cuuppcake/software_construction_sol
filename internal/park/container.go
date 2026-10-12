package park

type Container struct {
	Park *Park
}

func NewContainer(service ServiceCenter) *Container {
	return &Container{Park: NewPark(service)}
}
