package repository

import (
	"context"

	"superadmin.ru/users/model"
	"superadmin.ru/internal/postgres"
	"superadmin.ru/internal/postgres/dao"
)

type YclientsRegistrations struct {
	db *postgres.DB
}

func NewYclientsRegistrations(db *postgres.DB) *YclientsRegistrations {
	return &YclientsRegistrations{
		db: db,
	}
}

func (r *YclientsRegistrations) Save(ctx context.Context, m *model.YclientsRegistration) error {
	queries := r.db.Queries(ctx)

	id, err := queries.CreateActiveIntegration(ctx,
		dao.CreateActiveIntegrationParams{
			AccountID:    0,
			ExtCompanyID: m.ExtCompanyId,
			CompanyName:  m.CompanyName,
		})

	if err != nil {
		return err
	}

	id = id

	m.ID = id
	return nil
}
