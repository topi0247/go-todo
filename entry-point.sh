set -o errexit

sql-migrate up
go run main.go
