package postgres

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"superadmin.ru/infrastructure/postgres/dao"
)

type DB struct {
	*sql.DB
}

func Open(url string) *DB {
	db, err := sql.Open("postgres", url)

	if err == nil {
		err = db.Ping()
	}

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	log.Println("Connected to the database!")

	return &DB{
		db,
	}
}

func (db *DB) Queries(ctx context.Context) *dao.Queries {
	tx, ok := fromContext(ctx)

	if ok {
		return dao.New(tx)
	}

	return dao.New(db)
}
