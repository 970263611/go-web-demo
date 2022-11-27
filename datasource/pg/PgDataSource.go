package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"project/ent"
	"project/file"
	"time"
)

type DataSourcePG struct{}

func (pg DataSourcePG) Open() *sql.DB {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		file.GetEnvParam().DatasourceHost, file.GetEnvParam().DatasourcePort, file.GetEnvParam().DatasourceUser, file.GetEnvParam().DatasourcePassword, file.GetEnvParam().DatasourceTable)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func (pg DataSourcePG) Close(db *ent.Client) {
	db.Close()
}
