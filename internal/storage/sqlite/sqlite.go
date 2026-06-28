package sqlite

import (
	"database/sql"

	"github.com/nios-x/articles-go/internal/config"
	"github.com/nios-x/articles-go/internal/types"
	_ "modernc.org/sqlite"
)

type SqLite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*SqLite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)
	if err != nil {
		return nil, err
	}
	return &SqLite{
		Db: db,
	}, nil
}

func (s *SqLite) CreateUser(name string, email string, age int) (int64, error) {
	smtm, err := s.Db.Prepare("INSERT INTO users(name, email, age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer smtm.Close()
	res, err := smtm.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SqLite) GetUserByID(id int) (*types.User, error) {
	row := s.Db.QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id)

	var user types.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
