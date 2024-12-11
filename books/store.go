package books

type BookStore interface {
	List() []Book
	GetById(id string) *Book
	Create(book *Book) *Book
}
