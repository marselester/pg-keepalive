# Postgres keepalive

This repository shows three approaches to control TCP keepalive in `lib/pq` based on
https://github.com/lib/pq/pull/999.

Custom driver `sql.Open("postgres-test", dsn)`.

```sh
$ GOOS=linux go build ./cmd/driver
$ strace -e trace=network ./driver

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```

Ubiquitous `sql.Open("postgres", dsn)`.

```sh
$ GOOS=linux go build ./cmd/open
$ strace -e trace=network ./open

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```

Connector `sql.OpenDB(connector)`.

```sh
$ GOOS=linux go build ./cmd/connector
$ strace -e trace=network ./connector

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```
