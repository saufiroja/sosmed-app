package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/saufiroja/sosmed-app/account-service/config"
	"github.com/sirupsen/logrus"
)

type PostgresInterface interface {
	DbConnection() *sql.DB
	StartTransaction(ctx context.Context) *sql.Tx
	CommitTransaction(tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx)
}

type Postgres struct {
	*sql.DB
}

func NewPostgres(conf *config.AppConfig, logger *logrus.Logger) PostgresInterface {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	dbname := conf.Postgres.Name
	ssl := conf.Postgres.SSL

	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, ssl)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		logger.Panicf("Error opening database connection: %v\n", err)
	}

	// check connection
	err = db.Ping()
	if err != nil {
		logger.Panicf("Error connecting to database: %v\n", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5)
	db.SetConnMaxIdleTime(5)

	logger.Info("Connected to database")

	return &Postgres{db}
}

func (db *Postgres) DbConnection() *sql.DB {
	return db.DB
}

func (db *Postgres) StartTransaction(ctx context.Context) *sql.Tx {
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	return tx
}

func (db *Postgres) CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func (db *Postgres) RollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(err)
	}
}
