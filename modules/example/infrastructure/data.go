package infrastructure

import (
	"fmt"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
)

// PGExample represents the example repository
type PGExample struct {
	Db *database.Database
}

// Find example
func (pg *PGExample) Find(req ExampleModel) (string, error) {
	return fmt.Sprintf("Bem Vindo ao modulo example: %s", *req.Name), nil
}
