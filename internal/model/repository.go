package model

import (
	"github.com/jackc/pgx"
)

type ItemRepository struct {
	db *pgx.Conn
}

type Config struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func NewItemRepository(cfg Config) (*ItemRepository, error) {
	db, err := pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	})

	if err != nil {
		return nil, err
	}

	return &ItemRepository{db}, nil
}

func (r *ItemRepository) Add(items []Item) error {
	rows := make([][]interface{}, 0, len(items))

	for _, item := range items {
		rows = append(rows, []interface{}{item.Title})
	}

	_, err := r.db.CopyFrom([]string{"items"}, []string{"id", "title"}, pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) Update(item Item) error {
	_, err := r.db.Exec("UPDATE item SET title = $1 WHERE id = $2", item.Title, item.Id)
	if err != nil {
		return err
	}

	return nil
}
