package pg

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"project/ent"
	"time"
)

type DataSourcePG struct{}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pgadmin"
	dbname   = "postgres"
)

func (pg DataSourcePG) Open() *ent.Client {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entSql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func (pg DataSourcePG) Close(db *ent.Client) {
	db.Close()
}
