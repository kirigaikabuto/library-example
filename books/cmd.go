package books

type ListCommand struct {
}

func (cmd *ListCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).List(cmd)
}

type CreateCommand struct {
	Name string `json:"name,omitempty"`
}

func (cmd *CreateCommand) Exec(pub interface{}) (interface{}, error) {
	return pub.(Service).Create(cmd)
}

type GetByIdCommand struct {
	Id int `json:"id"`
}

func (cmd *GetByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetById(cmd)
}
