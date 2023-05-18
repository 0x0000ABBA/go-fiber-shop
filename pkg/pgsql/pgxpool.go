package pgsql

import (
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx"
)

type Client interface {
}

type pgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var ConnPool *pgx.ConnPool

func NewConfig(username, password, host, port, database string) *pgConfig {
	return &pgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(maxAttempts int, delay time.Duration, cfg *pgConfig, l *log.Logger) (*pgx.ConnPool, error) {
	l.Println("Parsing pgconfig")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Port, cfg.Database,
	)

	l.Println("Connecting using connection string:  ", connStr)
	l.Println("Creating connection pool")

	connCfg, err := pgx.ParseURI(connStr)
	poolCfg := pgx.ConnPoolConfig{
		ConnConfig: connCfg,
	}

	if err != nil {
		return nil, err
	}

	l.Println("Starting to connect to database")

	var pool *pgx.ConnPool

	for maxAttempts > 0 {
		pool, err = pgx.NewConnPool(poolCfg)
		if err != nil {
			l.Printf("Failed to connect to database, waiting for %s, then attempting again. %d attempts left. Error: %s", delay.String(), maxAttempts, err.Error())
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		l.Println("Connected to database")
		return pool, nil
	}
	log.Fatal("Failed to connect to database, no attempts left")
	return nil, err

}

func NewGlClient(maxAttempts int, delay time.Duration, cfg *pgConfig, l *log.Logger) error {
	l.Println("Parsing pgconfig")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Port, cfg.Database,
	)

	l.Println("Connecting using connection string:  ", connStr)
	l.Println("Creating connection pool")

	connCfg, err := pgx.ParseURI(connStr)
	poolCfg := pgx.ConnPoolConfig{
		ConnConfig: connCfg,
	}

	if err != nil {
		return err
	}

	l.Println("Starting to connect to database")

	for maxAttempts > 0 {
		ConnPool, err = pgx.NewConnPool(poolCfg)
		if err != nil {
			l.Printf("Failed to connect to database, waiting for %s, then attempting again. %d attempts left. Error: %s", delay.String(), maxAttempts, err.Error())
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		l.Println("Connected to database")
		return nil
	}
	log.Fatal("Failed to connect to database, no attempts left")
	return err

}