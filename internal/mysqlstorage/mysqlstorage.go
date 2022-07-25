package mysqlstorage

import (
	"context"
	"database/sql"
	"fmt"
)

type MyStoreStore struct {
	context.Context
	Dbh *sql.DB
}

func MyStoreConstructor(config map[string]string) MyStoreStore {
	db, _ := sql.Open("mysql", "database="+config["db"])
	return MyStoreStore{nil, db}
}

func (s MyStoreStore) Insert(items []string) {
	stmt, err := s.Dbh.Prepare("INSERT INTO items(id, title) VALUES(?, ?")
	if err != nil {
		panic(err)
	}

	for i, item := range items {
		stmt.ExecContext(s.Context, i, item)
	}
}

func (s *MyStoreStore) Update(id int, title string) {
	stmt, err := s.Dbh.Prepare(fmt.Sprintf("UPDATE INTO users SET title = '%s' WHERE id = %d", title, id))
	if err != nil {
		panic(err)
	}

	stmt.ExecContext(s.Context)
}
