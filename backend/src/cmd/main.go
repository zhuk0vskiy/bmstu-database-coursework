package main

import (
	"app/src/config"
	"app/src/pkg/logger"
	"app/src/tui"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

//func main() {
//	cfg, err := config.ReadConfig()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	pool, err := newConn(context.Background(), &cfg.DBConfig)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	app := app.NewApp(pool, cfg)
//	termui := tui.NewTUI(app)
//
//	err = termui.Run()
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

func main() {
	ctx := context.Background()
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := newConn(ctx, &c.Database)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(1)

	// Create logger

	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)

	l := logger.New(c.Logger.Level, loggerFile)

	tui.Run(db, c, l)
}

func newConn(ctx context.Context, cfg *config.DatabaseConfig) (pool *pgxpool.Pool, err error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.Postgres.Driver,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("подключение к БД: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("пинг БД: %w", err)
	}

	return pool, nil
}
