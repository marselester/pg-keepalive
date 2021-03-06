// Package open helps to test TCP keepalive support when pq.Open is called directly.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func main() {
	var (
		pgHost      = flag.String("host", "localhost", "Postgres host")
		pgPort      = flag.Uint("port", 5432, "Postgres port")
		pgDatabase  = flag.String("database", "postgres", "Postgres database name")
		pgUser      = flag.String("user", "postgres", "Postgres user")
		pgPassword  = flag.String("password", "", "Postgres password")
		pgKeepAlive = flag.Int("keepalives_interval", 5, "interval between keep-alive probes in seconds for an active Postgres connection")
	)
	flag.Parse()

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=3 application_name=testopen keepalives_interval=%d sslmode=disable binary_parameters=yes", *pgHost, *pgPort, *pgDatabase, *pgUser, *pgPassword, *pgKeepAlive)
	c, err := pq.Open(dsn)
	if err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}
	defer c.Close()

	txn, err := c.Begin()
	if err != nil {
		log.Fatalf("postgres query failed: %v", err)
	}
	txn.Rollback()
}
