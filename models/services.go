package models

import (
	"context"
	"fmt"
	"github.com/daniel-z-johnson/funcss/conf"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Services struct {
	FunCSS FunCSSService
}

// TODO implement this
func NewServices(configFile string) *Services {
	appConfig, err := conf.LoadFunCSSConf(configFile)
	if err != nil {
		panic(err)
	}
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DB.Host, appConfig.DB.Port, appConfig.DB.User, appConfig.DB.Password, appConfig.DB.Name)
	config, err := pgxpool.ParseConfig(dbString)
	if err != nil {
		panic(err)
	}
	config.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		fmt.Println("Before Acquire")
		return true
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		fmt.Println("After Connect")
		return nil
	}
	config.AfterRelease = func(conn *pgx.Conn) bool {
		fmt.Println("AfterRelease")
		return true
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	return &Services{
		FunCSS: NewFunCSSService(pool),
	}
}
