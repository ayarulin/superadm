package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct { 
  db *sql.DB
  ActiveIntegrations *activeIntegrations
}

func Open(host, dbname, user, password string, port int) *PostgresDB {
  connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", connInfo)
   
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
    db: db,
    ActiveIntegrations: &activeIntegrations{db},
  }
}

func (p *PostgresDB) Close() {
  p.db.Close()
}

