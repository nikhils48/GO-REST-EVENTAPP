package database

import "database/sql"

type UserModel struct {
	DB *sql.DB
}
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
