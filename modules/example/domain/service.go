package domain

import (
	"errors"

	"github.com/jinzhu/copier"
	"go.uber.org/dig"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/infrastructure"
)

type repositoryDependencies struct {
	dig.In
	Database *database.Database `name:"DATABASE"`
}
type servicesDependencies struct {
	dig.In
	Repo IExample `name:"EXAMPLE-REPOSITORY"`
}

// Service represents the example service
type Service struct {
	repo IExample
}

// GetService returns the example service
func GetService(dep servicesDependencies) *Service {
	return &Service{repo: dep.Repo}
}

// GetRepository returns the example repository
func GetRepository(dep repositoryDependencies) IExample {
	return newRepository(dep.Database)
}

func (s *Service) Find(req ExampleModel) (string, error) {
	data := infrastructure.ExampleModel{}
	if req.Name == nil || *req.Name == "" {
		return "", errors.New("name is required")
	}

	if err := copier.Copy(&data, req); err != nil {
		return "", err
	}

	return s.repo.Find(data)
}
