// Package driver helps to test TCP keepalive support when lib/pq is used with a custom driver.
package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/lib/pq"
)

// keepAlive is interval between keep-alive probes in seconds for an active Postgres connection.
const keepAlive = 5 * time.Second

func main() {
	var (
		pgHost     = flag.String("host", "localhost", "Postgres host")
		pgPort     = flag.Uint("port", 5432, "Postgres port")
		pgDatabase = flag.String("database", "postgres", "Postgres database name")
		pgUser     = flag.String("user", "postgres", "Postgres user")
		pgPassword = flag.String("password", "", "Postgres password")
	)
	flag.Parse()

	sql.Register("postgres-test", &pqDriver{})

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=3 application_name=testdriver sslmode=disable binary_parameters=yes", *pgHost, *pgPort, *pgDatabase, *pgUser, *pgPassword)
	db, err := sql.Open("postgres-test", dsn)
	if err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("postgres ping failed: %v", err)
	}
	db.Close()
}

type pqDriver struct{}

func (d *pqDriver) Open(name string) (driver.Conn, error) {
	return pq.DialOpen(
		&dialer{net.Dialer{KeepAlive: keepAlive}},
		name,
	)
}

type dialer struct {
	d net.Dialer
}

func (d dialer) Dial(network, address string) (net.Conn, error) {
	return d.d.Dial(network, address)
}

func (d dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}

func (d dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return d.d.DialContext(ctx, network, address)
}
