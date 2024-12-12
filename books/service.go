package books

type Service interface {
	List(cmd *ListCommand) ([]Book, error)
	GetById(cmd *GetByIdCommand) (*Book, error)
	Create(cmd *CreateCommand) (*Book, error)
}

type service struct {
	store Store
}

func NewService(store Store) Service {
	return &service{store: store}
}

func (s *service) List(cmd *ListCommand) ([]Book, error) {
	return s.store.List()
}

func (s *service) GetById(cmd *GetByIdCommand) (*Book, error) {
	return s.store.GetById(cmd.Id)

}

func (s *service) Create(cmd *CreateCommand) (*Book, error) {
	return s.store.Create(&Book{
		Name: cmd.Name,
	})
}
