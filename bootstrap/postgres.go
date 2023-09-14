package bootstrap

import (
	"database/sql"
	"github.com/SawitProRecruitment/UserService/commons"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	Db *sql.DB
}

func NewPostgresClient(dsn string, maxOpenConn int, maxIdleConn int) *PostgresClient {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(commons.ErrorConnectionToDatabase.Error())
	}
	// Set the maximum number of open and idle connections
	db.SetMaxOpenConns(maxOpenConn) // Maximum number of open connections
	db.SetMaxIdleConns(maxIdleConn) // Maximum number of idle connections

	return &PostgresClient{
		Db: db,
	}
}
