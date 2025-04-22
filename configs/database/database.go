package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database interface {
    Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
    Get(dest interface{}, query string, args ...interface{}) error
    Exec(query string, args ...interface{}) (sql.Result, error)
    NamedExec(query string, arg interface{}) (sql.Result, error)
	MustBegin() *sqlx.Tx
	Ping() error
	SetMaxOpenConns(n int)
	SetMaxIdleConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Close() error
}

var DB *sqlx.DB  

func Connect() error {
	cfg := LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	DB.SetMaxOpenConns(10) 
	DB.SetMaxIdleConns(2)  
	DB.SetConnMaxLifetime(time.Minute)

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connected using MySQL with sqlx!")
	return nil
}
