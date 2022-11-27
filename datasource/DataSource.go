package datasource

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	"log"
	"project/datasource/pg"
	"project/ent"
	"project/ent/migrate"
)

var db *sql.DB
var client *ent.Client

type DataSource interface {
	Open() *sql.DB
	Close(*ent.Client)
}

func init() {
	Pg()
}

func Pg() {
	var pg DataSource = &pg.DataSourcePG{}
	db = pg.Open()
	drv := entSql.OpenDB(dialect.Postgres, db)
	client = ent.NewClient(ent.Driver(drv))
}

func Client() *ent.Client {
	//if err := client.Debug().Schema.Create(context.Background(),migrate.WithGlobalUniqueID(false)); err != nil {
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(false)); err != nil {
		log.Fatal(err)
	}
	return client
}

func NativeSqlQuery(sql string, args ...any) ([]map[string]string, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0, 10)
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]string, 10)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}
