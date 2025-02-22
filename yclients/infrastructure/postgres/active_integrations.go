package postgres

import (
	"database/sql"
	"superadmin.ru/yclients/internal/command/core"
)

type activeIntegrations struct {
	db *sql.DB
}

// Create implements core.ActiveIntegrations.
func (a *activeIntegrations) Create(*core.ActiveIntegration) error {
	panic("unimplemented")
}

// Get implements core.ActiveIntegrations.
func (a *activeIntegrations) Get(core.ActiveIntegrationId) (*core.ActiveIntegration, error) {
	panic("unimplemented")
}
