package datasource

import (
	"context"
	"log"
	"project/datasource/pg"
	"project/ent"
	"project/ent/migrate"
)

var client *ent.Client

type DataSource interface {
	Open() *ent.Client
	Close(*ent.Client)
}

func init() {
	Pg()
}

func Pg() {
	var pg DataSource = &pg.DataSourcePG{}
	client = pg.Open()
}

func Client() *ent.Client {
	//if err := client.Debug().Schema.Create(context.Background(),migrate.WithGlobalUniqueID(false)); err != nil {
	if err := client.Schema.Create(context.Background(),migrate.WithGlobalUniqueID(false)); err != nil {
		log.Fatal(err)
	}
	return client
}
