package books

type ListCommand struct {
}

func (cmd *ListCommand) Exec(service Service) (interface{}, error) {
	return service.List(cmd)
}

type CreateCommand struct {
	Name string `json:"name,omitempty"`
}

func (cmd *CreateCommand) Exec(service Service) (interface{}, error) {
	return service.Create(cmd)
}

type GetByIdCommand struct {
	Id int `json:"id"`
}

func (cmd *GetByIdCommand) Exec(service Service) (interface{}, error) {
	return service.GetById(cmd)
}
