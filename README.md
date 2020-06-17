# SQLEX
[![Build Status](https://travis-ci.org/clevergo/sqlex.svg?branch=master)](https://travis-ci.org/clevergo/sqlex)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/sqlex/badge.svg?branch=master)](https://coveralls.io/github/clevergo/sqlex?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/clevergo.tech/sqlex?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/sqlex)](https://goreportcard.com/report/github.com/clevergo/sqlex)
[![Release](https://img.shields.io/github/release/clevergo/sqlex.svg?style=flat-square)](https://github.com/clevergo/sqlex/releases)

SQLEX is extensions for database/sql and jmoiron/sqlx.

## Transaction

```go
sqlex.Transact(db, func(tx *sql.Tx) error {
    return nil
})

sqlex.TransactContext(ctx, txOpts, db, func(tx *sql.Tx) error {
    return nil
})

sqlex.Transactx(db, func(tx *sqlx.Tx) error {
    return nil
})

sqlex.TransactContext(ctx, txOpts, db, func(tx *sqlx.Tx) error {
    return nil
})
```
