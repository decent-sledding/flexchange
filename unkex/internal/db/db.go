package db

import (
	"errors"
	"database/sql"

	"unkex/config"
)


// Provided a configuration creates a new database connection instance
//
func getDatabase(conf *config.AppConfiguration) (*sql.DB, error) {
	db, err := sql.Open("postgres", GetDsn(conf))

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}


// Gets a DSN (connection string) from provided AppConfiguration
//
// Later we could extend a configuration
// to be able to select different database engine
func getDsn(conf *AppConfiguration) (string, error) {
	dconf = conf.Database

	if dconf.User == "" {
		return nil, errors.New("internal/db: No db user")
	}

	if dconf.Host == "" {
		return nil, errors.New("internal/db: No db host")
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dconf.User,
		dconf.Passowrd,
		dconf.Host,
		dconf.Port,
		dconf.DatabaseName
	)
}