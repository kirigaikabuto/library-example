package books

type Service interface {
	List(cmd *ListCommand) ([]Book, error)
	GetById(cmd *GetByIdCommand) (*Book, error)
	Create(cmd *CreateCommand) (*Book, error)
}
