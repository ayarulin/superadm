package core

import "database/sql"

type Repository struct {
  db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
  return &Repository{
    db,
  }
}

func(r *Repository) SaveActiveIntegration(c context.Context, integration *ActiveIntegration) error {
  return nil
}
