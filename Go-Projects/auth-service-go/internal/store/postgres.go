package store

import (
	"auth-service-go/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Function fetches configuration files & try to connect with postgreSQL
func NewPostgres(cfg *config.Config) (*pgxpool.Pool, error) {

	// Build Database Source name: includes (user, pass, host, port, etc..)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Create context with 5 sec timeout, if taking much longer time connection attempt cancel.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Creating Connection Pool, not single connection, to handel Reuse DB connections, handel concurrent request, much faster than opening new connection everytime..
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// This will actively check if DB reachable, are cred correct, is network ok?
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	// log request, that w have connected to DB
	log.Println("Connected to PostgreSQL")
	return pool, nil
}
