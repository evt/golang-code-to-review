package model

import (
	"github.com/jackc/pgx"
)

type ItemRepository struct {
	db *pgx.Conn
}

func NewItemRepository(db *pgx.Conn) *ItemRepository {
	return &ItemRepository{db: db}
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
	_, err := r.db.Exec("UPDATE items SET title = $1 WHERE id = $2", item.Title, item.Id)
	if err != nil {
		return err
	}

	return nil
}
