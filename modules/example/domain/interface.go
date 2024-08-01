package domain

import "github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/infrastructure"

// IExample represents the example domain
type IExample interface {
	Find(req infrastructure.ExampleModel) (string, error)
}
