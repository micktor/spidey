package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/ent"
	"github.com/micktor/spidey/internal/ent/migrate"
)

func Connect(config *config.Config) *ent.Client {
	fmt.Println("creating db connection")

	DBClient, err := ent.Open("mysql", config.Database.Username+
		":"+config.Database.Password+
		"@tcp("+config.Database.Hostname+
		":"+config.Database.Port+")/"+
		config.Database.Database+"?parseTime=True")
	if err != nil {
		fmt.Println("failed opening connection to database: ", err)
	}
	return DBClient
}

func Close(con *ent.Client) {
	fmt.Println("closing db connection")

	if err := con.Close(); err != nil {
		fmt.Println("error closing db connection")
	}
}

func Migrate(con *ent.Client) {
	fmt.Println("running migration")

	// Run the auto migration tool.
	if err := con.Schema.Create(context.Background(),
		migrate.WithGlobalUniqueID(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		fmt.Println("failed creating schema resources: ", err)
	} else {
		fmt.Println("migration successful =)")
	}
}
