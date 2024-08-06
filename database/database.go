package database

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/config"
	"github.com/Lucas-Linhar3s/Base-Structure-Golang/pkg/log"
)

type datbaseDependencies struct {
	dig.In
	Config *config.Config `name:"CONFIG"`
	Logger *log.Logger    `name:"LOGGER"`
}

type Database struct {
	db                 *sql.DB
	transactionTimeout int
	Builder            sq.StatementBuilderType
}

func NewDatabase(dep datbaseDependencies) *Database {
	db, err := open(dep.Config)
	if err != nil {
		dep.Logger.Fatal("failed to open database", zap.Error(err))
	}

	return db
}

func open(c *config.Config) (database *Database, err error) {
	var db *sql.DB

	if c.Data.Db.User.Driver == "mysql" || c.Data.Db.User.Driver == "postgres" {
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
	} else if c.Data.Db.User.Driver == "sqlite" {
		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path := filepath.Join(currentDir, "", c.Data.Db.User.Dsn)

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
