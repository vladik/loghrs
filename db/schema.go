package db

import (
	"database/sql"
)

// Init initializes the database schema if tables do not exist
func Init(db *sql.DB) error {
	schema := `
CREATE TABLE IF NOT EXISTS shelves (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS journals (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	shelf_id INTEGER,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (shelf_id) REFERENCES shelves(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS activities (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	journal_id INTEGER NOT NULL,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(journal_id, name),
	FOREIGN KEY (journal_id) REFERENCES journals(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS labels (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS times (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	activity_id INTEGER NOT NULL,
	date TEXT NOT NULL,
	duration INTEGER  NOT NULL,
	note TEXT,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS time_labels (
	time_id INTEGER NOT NULL,
	label_id INTEGER NOT NULL,
	PRIMARY KEY (time_id, label_id),
	FOREIGN KEY (time_id) REFERENCES times(id) ON DELETE CASCADE,
	FOREIGN KEY (label_id) REFERENCES labels(id) ON DELETE CASCADE
);
`
	_, err := db.Exec(schema)
	return err
}
