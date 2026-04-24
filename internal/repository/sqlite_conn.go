package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitSQLiteDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		document_number TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS operation_types (
		id INTEGER PRIMARY KEY,
		description TEXT NOT NULL,
		multiplier INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account_id INTEGER NOT NULL,
		operation_type_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		event_date DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(account_id) REFERENCES accounts(id),
		FOREIGN KEY(operation_type_id) REFERENCES operation_types(id)
	);

	INSERT OR IGNORE INTO operation_types (id, description, multiplier) VALUES 
		(1, 'COMPRA A VISTA', -1),
		(2, 'COMPRA PARCELADA', -1),
		(3, 'SAQUE', -1),
		(4, 'PAGAMENTO', 1);
	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
