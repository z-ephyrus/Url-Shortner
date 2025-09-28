package db

import (
	"context"
	"database/sql"
	"log"

	"url-shortner/internal/schema"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)
		

	var pg *bun.DB

func Connect(dsn string) *bun.DB{
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	pg = bun.NewDB(sqldb, pgdialect.New())

	if err := sqldb.Ping(); err != nil{
		log.Fatalf("Database Connection Error : %v",err)
	}

	runMigrations()

	return pg
}


func Close()  {
	if pg != nil{
		_ = pg.Close()
	}
}


func runMigrations() {
    ctx := context.Background()
    _, err := pg.NewCreateTable().
        Model((*schema.User)(nil)).
        IfNotExists().
        Exec(ctx)

    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    log.Println("Migrations are complete!")
}
