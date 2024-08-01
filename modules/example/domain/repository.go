package domain

import (
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/database"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/modules/example/infrastructure"
)

type repository struct {
	pg infrastructure.PGExample
}

func newRepository(db *database.Database) *repository {
	return &repository{
		pg: infrastructure.PGExample{
			Db: db,
		},
	}
}

// Find implements IExample.
func (r *repository) Find(req infrastructure.ExampleModel) (string, error) {
	return r.pg.Find(req)
}
