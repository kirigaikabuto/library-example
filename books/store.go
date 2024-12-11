package books

type Store interface {
	List() ([]Book, error)
	GetById(id int) (*Book, error)
	Create(book *Book) (*Book, error)
}
