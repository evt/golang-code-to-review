package mysqlstorage

import (
	"context"
	"database/sql"
	"fmt"
)

type MyStoreStore struct {
	context.Context
	db *sql.DB
}

func MyStoreConstructor(config map[string]string) MyStoreStore {
	db, _ := sql.Open("mysql", "database="+config["db"])
	return MyStoreStore{nil, db}
}

func (s MyStoreStore) Insert(items []string) error {
	stmt, err := s.db.Prepare("INSERT INTO items(id, title) VALUES(?, ?")
	if err != nil {
		return err
	}

	for i, item := range items {
		_, err := stmt.ExecContext(s.Context, i, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *MyStoreStore) Update(id int, title string) error {
	stmt, err := s.db.Prepare(fmt.Sprintf("UPDATE INTO users SET title = '%s' WHERE id = %d", title, id))
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(s.Context)
	if err != nil {
		return err
	}

	return nil
}
