# Postgres keepalive

This repository shows a few approaches to control
[TCP keepalive](https://www.ibm.com/support/pages/ibm-aix-tcp-keepalive-probes)
in `lib/pq` based on https://github.com/lib/pq/pull/999.

Custom driver `sql.Open("postgres-custom", dsn)`.

```sh
$ go build ./cmd/driver
$ strace -e trace=network ./driver

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```

Ubiquitous `sql.Open("postgres", dsn)`.

```sh
$ go build ./cmd/sqlopen
$ strace -e trace=network ./sqlopen

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```

Rarely used `pq.Open(dsn)`.

```sh
$ go build ./cmd/open
$ strace -e trace=network ./open

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```

Connector `sql.OpenDB(connector)`.

```sh
$ go build ./cmd/connector
$ strace -e trace=network ./connector

setsockopt(3, SOL_SOCKET, SO_KEEPALIVE, [1], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPINTVL, [5], 4) = 0
setsockopt(3, SOL_TCP, TCP_KEEPIDLE, [5], 4) = 0
```
