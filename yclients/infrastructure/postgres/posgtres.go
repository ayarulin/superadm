package postgres
// TODO: extract common logic to pkg/postgres

import (
	"database/sql"
  "superadmin.ru/yclients/infrastructure/postgres/dao"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct { 
  ActiveIntegrations *activeIntegrations
  closeFn func() error
}

func Open(url string) *PostgresDB {
  db, err := sql.Open("postgres", url)

  queries := dao.New(db)
   
  if err == nil {
    err = db.Ping()
  }

  if err != nil {
    log.Fatal("Ошибка подключения к базе данных: ", err)
  }

  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(5)
  db.SetConnMaxLifetime(0)

  log.Println("Connected to the database!")

  return &PostgresDB{
    closeFn: db.Close,
    ActiveIntegrations: &activeIntegrations{queries},
  }
}

func (p *PostgresDB) Close() error {
  return p.closeFn()
}

