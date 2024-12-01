package repository

import (
	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) error {
	stmt := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY ,
			email VARCHAR(100) UNIQUE NOT NULL,
			name VARCHAR(100) UNIQUE NOT NULL,
			password_hash VARCHAR(100) NOT NULL,
			role_id INTEGER default 1,
			deleted_at TIMESTAMP DEFAULT null,
			notifications boolean default false,
			is_deleted boolean default false,
			device_token varchar(255) default null
		);

		
	
		CREATE TABLE if not exists tokens (
			token VARCHAR(255) PRIMARY KEY,
			revoked BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		

    `

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
