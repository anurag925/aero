package models

type SchemaMigrations struct {
	Version int `json:"version"`
	Dirty   int `json:"dirty"`
}
