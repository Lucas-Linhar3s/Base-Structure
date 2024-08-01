package application

// ExampleReq represents the example request
type ExampleReq struct {
	Name *string `copier:"Name" json:"name"`
}

// ExampleRes represents the example response
type ExampleRes struct {
	Data string `json:"data"`
}
