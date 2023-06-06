package connection

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

func MysqlConfig(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Println("failed to connect to database")
		return nil, errors.New("failed to create connection to database")
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(200)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
