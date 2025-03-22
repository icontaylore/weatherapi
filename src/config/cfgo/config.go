package cfgo

import (
	"bot_notif/src/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

func Cfg() (*models.Config, error) {
	if err := godotenv.Load("../../config/.env"); err != nil {
		log.Fatal(err)
	}
	newCfg := &models.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB"),
		DBuser:   os.Getenv("POSTGRES_USER"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
	if newCfg == nil {
		log.Fatal("nil config")
	}
	return newCfg, nil
}

func ParseCfgDefaultSql(config *models.Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(config.DBuser),
		config.Password,
		config.Host,
		config.Port,
		config.DBname,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("ошибка с подкл к базе")
	}
	return db
}

func ParseCfg(config *models.Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(config.DBuser),
		config.Password,
		config.Host,
		config.Port,
		config.DBname,
	)

	conn, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return conn, err
}

func StartDb(parseCfg *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), parseCfg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("коннект к базе\n")

	if _, err = conn.Exec(context.Background(), ";"); err != nil {
		return nil, err
	}

	return conn, nil
}
