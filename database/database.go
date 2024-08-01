package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
)

type Database struct {
	db                 *sql.DB
	transactionTimeout int
	Builder            sq.StatementBuilderType
}

func Open(c *config.Config, dir string, driver string) (database *Database, err error) {
	var db *sql.DB

	if driver == "mysql" || driver == "postgres" {
		driverConfig := stdlib.DriverConfig{
			ConnConfig: pgx.ConnConfig{
				RuntimeParams: map[string]string{
					//Verificar
					"application_name": "github.com/Lucas-Linhar3s/Base-Structure-Golang",
					"DateStyle":        "ISO",
					"IntervalStyle":    "iso_8601",
					// TODO:
					"search_path": "public",
				},
			},
		}
		stdlib.RegisterDriverConfig(&driverConfig)

		db, err = sql.Open("pgx", driverConfig.ConnectionString(
			c.Data.Db.User.Nick+
				"://"+
				c.Data.Db.User.Username+
				":"+
				c.Data.Db.User.Password+
				"@"+
				c.Data.Db.User.Hostname+
				":"+
				c.Data.Db.User.Port+
				"/"+
				c.Data.Db.User.Name))
		if err != nil {
			return nil, err
		}
	} else if driver == "sqlite" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path := filepath.Join(currentDir, dir, c.Data.Db.User.Dsn)

		// Verifica se o arquivo do banco de dados existe
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, errors.New("banco de dados sqlite não encontrado")
		}

		db, err = sql.Open("sqlite3", path)
		if err != nil {
			return nil, err
		}
	} else {
		panic("unknown db driver")
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.Data.Db.User.MaxIdle)
	db.SetMaxOpenConns(c.Data.Db.User.MaxConn)
	db.SetConnMaxLifetime(time.Second * 60)

	return &Database{
		db:                 db,
		transactionTimeout: c.Data.Db.User.TransactionTimeout,
		Builder:            sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db),
	}, nil
}

func (p *Database) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
