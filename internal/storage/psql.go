package storage

import (
	//"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"

	"github.com/keenfury/library_api/config"
)

type (
	DB struct{}
)

var PsqlDB *sqlx.DB

func init() {
	var err error
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", config.PostgresUser, config.PostgresPass, config.PostgresDB, config.PostgresHost)

	PsqlDB, err = sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Panicln("Could not connect with connection string:", connectionStr)
	}
}
