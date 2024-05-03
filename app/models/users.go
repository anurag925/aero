//go:generate go run github.com/dmarkham/enumer -type=UserStatus -json -transform=snake
package models

import (
	"database/sql"
	"time"
)

type UserStatus int8

const (
	UserStatusCreated UserStatus = iota
	UserStatusActive
	UserStatusDeleted
)

type User struct {
	applicationModel

	ID             int64          `json:"id" bun:",pk,autoincrement"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	MobileNumber   sql.NullString `json:"mobile_number"`
	Status         UserStatus     `json:"status"`
	Verified       bool           `json:"verified"`
	Iter           sql.NullInt64  `json:"iter"`
	Salt           sql.NullString `json:"salt"`
	PasswordDigest sql.NullString `json:"password_digest"`
	TenantID       int64          `json:"tenant_id"`
}
