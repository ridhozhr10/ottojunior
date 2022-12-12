package cliaction

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/pkg/database"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// MigrateAction cli action
func MigrateAction(c *cli.Context) error {
	conf := database.Config{
		DBHost: c.String("db-host"),
		DBUser: c.String("db-user"),
		DBPass: c.String("db-pass"),
		DBPort: c.String("db-port"),
		DBName: c.String("db-name"),
	}
	db, err := database.ConnectGORM(conf)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return db.AutoMigrate(
		&model.User{},
		&model.Topup{},
		&model.Transaction{},
		&model.Balance{},
	)
}
