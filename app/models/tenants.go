//go:generate go run github.com/dmarkham/enumer -type=TenantStatus -json -transform=snake

package models

import (
	"database/sql"
	"time"
)

type TenantStatus int8

const (
	TenantStatusCreated TenantStatus = iota
	TenantStatusActive
	TenantStatusDeleted
)

type Tenant struct {
	applicationModel

	ID         int64          `json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Name       string         `json:"name"`
	Domain     sql.NullString `json:"domain"`
	Status     int8           `json:"status"`
	SchemaName sql.NullString `json:"schema_name"`
	DataSource sql.NullString `json:"data_source"`
	Settings   sql.NullString `json:"settings"`
}
