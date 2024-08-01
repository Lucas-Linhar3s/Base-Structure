package domain

import (
	"errors"

	"github.com/jinzhu/copier"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/infrastructure"
)

// Service represents the example service
type Service struct {
	repo IExample
}

// GetService returns the example service
func GetService(repo IExample) *Service {
	return &Service{repo: repo}
}

// GetRepository returns the example repository
func GetRepository(db *database.Database) IExample {
	return newRepository(db)
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
