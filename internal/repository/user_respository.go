package repository

import "database/sql"

type UserRepository interface{}

type userRepository struct {
	DB_READ  *sql.DB
	DB_WRITE *sql.DB
}

func NewUserRepository(dbRead *sql.DB, dbWrite *sql.DB) UserRepository {
	return &userRepository{DB_READ: dbRead, DB_WRITE: dbWrite}
}
