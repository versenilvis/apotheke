package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/verse/apotheke/internal/model"
)

type Store struct {
	db *sql.DB
}

func New(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}
	if err := store.init(); err != nil {
		db.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) init() error {
	schema := `
	CREATE TABLE IF NOT EXISTS commands (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		cmd TEXT NOT NULL,
		cwd TEXT,
		tags TEXT DEFAULT '',
		confirm INTEGER DEFAULT 0,
		frequency INTEGER DEFAULT 0,
		last_used DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_name ON commands(name);
	CREATE INDEX IF NOT EXISTS idx_frequency ON commands(frequency DESC);
	`
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Add(cmd *model.Command) error {
	query := `
	INSERT INTO commands (name, cmd, cwd, tags, confirm, frequency, created_at)
	VALUES (?, ?, ?, ?, ?, 0, CURRENT_TIMESTAMP)
	`
	confirm := 0
	if cmd.Confirm {
		confirm = 1
	}
	_, err := s.db.Exec(query, cmd.Name, cmd.Cmd, cmd.Cwd, cmd.Tags, confirm)
	return err
}

func (s *Store) Remove(name string) error {
	_, err := s.db.Exec("DELETE FROM commands WHERE name = ?", name)
	return err
}

func (s *Store) Get(name string) (*model.Command, error) {
	query := `
	SELECT id, name, cmd, cwd, tags, confirm, frequency, last_used, created_at
	FROM commands WHERE name = ?
	`
	row := s.db.QueryRow(query, name)
	return scanCommand(row)
}

func (s *Store) List(tag string) ([]*model.Command, error) {
	var query string
	var args []interface{}

	if tag != "" {
		query = `
		SELECT id, name, cmd, cwd, tags, confirm, frequency, last_used, created_at
		FROM commands WHERE tags LIKE ? ORDER BY frequency DESC, name ASC
		`
		args = append(args, "%"+tag+"%")
	} else {
		query = `
		SELECT id, name, cmd, cwd, tags, confirm, frequency, last_used, created_at
		FROM commands ORDER BY frequency DESC, name ASC
		`
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commands []*model.Command
	for rows.Next() {
		cmd, err := scanCommandRows(rows)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}
	return commands, rows.Err()
}

func (s *Store) Search(query string) ([]*model.Command, error) {
	sqlQuery := `
	SELECT id, name, cmd, cwd, tags, confirm, frequency, last_used, created_at
	FROM commands 
	WHERE name LIKE ? OR cmd LIKE ?
	ORDER BY 
		CASE WHEN name = ? THEN 0
		     WHEN name LIKE ? THEN 1
		     ELSE 2
		END,
		frequency DESC
	`
	rows, err := s.db.Query(sqlQuery,
		"%"+query+"%", "%"+query+"%",
		query, query+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commands []*model.Command
	for rows.Next() {
		cmd, err := scanCommandRows(rows)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}
	return commands, rows.Err()
}

func (s *Store) IncrementUsage(name string) error {
	query := `
	UPDATE commands 
	SET frequency = frequency + 1, last_used = CURRENT_TIMESTAMP
	WHERE name = ?
	`
	_, err := s.db.Exec(query, name)
	return err
}

func (s *Store) Update(cmd *model.Command) error {
	query := `
	UPDATE commands 
	SET cmd = ?, cwd = ?, tags = ?, confirm = ?
	WHERE name = ?
	`
	confirm := 0
	if cmd.Confirm {
		confirm = 1
	}
	_, err := s.db.Exec(query, cmd.Cmd, cmd.Cwd, cmd.Tags, confirm, cmd.Name)
	return err
}

func scanCommand(row *sql.Row) (*model.Command, error) {
	var cmd model.Command
	var confirm int
	var lastUsed sql.NullTime
	var cwd sql.NullString

	err := row.Scan(
		&cmd.ID, &cmd.Name, &cmd.Cmd, &cwd, &cmd.Tags,
		&confirm, &cmd.Frequency, &lastUsed, &cmd.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	cmd.Confirm = confirm == 1
	if lastUsed.Valid {
		cmd.LastUsed = &lastUsed.Time
	}
	if cwd.Valid {
		cmd.Cwd = &cwd.String
	}
	return &cmd, nil
}

func scanCommandRows(rows *sql.Rows) (*model.Command, error) {
	var cmd model.Command
	var confirm int
	var lastUsed sql.NullTime
	var cwd sql.NullString

	err := rows.Scan(
		&cmd.ID, &cmd.Name, &cmd.Cmd, &cwd, &cmd.Tags,
		&confirm, &cmd.Frequency, &lastUsed, &cmd.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	cmd.Confirm = confirm == 1
	if lastUsed.Valid {
		cmd.LastUsed = &lastUsed.Time
	}
	if cwd.Valid {
		cmd.Cwd = &cwd.String
	}
	return &cmd, nil
}

func (s *Store) GetAll() ([]*model.Command, error) {
	return s.List("")
}

var _ = time.Now
