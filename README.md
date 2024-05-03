# IDENTITY

## Setup Aero

### Install packages
go install github.com/a-h/templ/cmd/templ@latest
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
https://github.com/dmarkham/enumer
go install github.com/fraenky8/tables-to-go@master
go install github.com/anurag925/gomodifytags@latest

### Steps
1. create database `create database aero_development`
2. run `make migrate_up`