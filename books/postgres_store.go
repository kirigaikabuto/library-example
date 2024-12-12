package books

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS books (
		id serial primary key,
		name text
	);`,
}

type postgresStore struct {
	db *sql.DB
}

func NewPostgresStore(cfg PostgresConfig) (Store, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgresStore{db: db}, err
}

func (ps *postgresStore) List() ([]Book, error) {
	var books []Book
	data, err := ps.db.Query("select * from books")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		book := Book{}
		err = data.Scan(&book.Id, &book.Name)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (ps *postgresStore) Create(book *Book) (*Book, error) {
	err := ps.db.QueryRow("insert into books (name) values ($1) RETURNING id", book.Name).Scan(&book.Id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (ps *postgresStore) GetById(id int) (*Book, error) {
	book := &Book{}
	err := ps.db.QueryRow("select * from books where id= $1", id).Scan(&book.Id, &book.Name)
	if err != nil {
		return nil, err
	}
	return book, nil
}
