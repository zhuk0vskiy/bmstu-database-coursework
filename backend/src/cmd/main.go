package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/config"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/pkg/logger"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/pkg/monitoring"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/tui"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
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
	m := monitoring.New(c.Monitoring.Url)

	db, err := newConn(ctx, &c.Database)
	if err != nil {
		l.Fatalf("failed to connect to database: %v", err)
	}

	//atexit.Register(func() {
	//	_ = monitoring.DecUsersOnline()
	//})
	//atexit.Exit()
	//atexit.Fatal()
	err = m.IncUsersOnline()
	if err != nil {
		l.Errorf("failed to increment users online: %v", err)
	}
	defer func() {
		_ = m.DecUsersOnline()
	}()
	tui.Run(db, c, l, m)
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
