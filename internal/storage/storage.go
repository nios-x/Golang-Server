package storage

import "github.com/nios-x/articles-go/internal/types"

type Storage interface {
	CreateUser(name string, email string, age int) (int64, error)
	GetUserByID(id int) (*types.User, error)
}
