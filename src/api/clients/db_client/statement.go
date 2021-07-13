package db_client

type Statement struct {
	Name       string
	Query      string
	Args       []interface{}
	InsertedId int64
}

func NewStatement(name string, query string, arg ...interface{}) Statement {
	return Statement {
		Name: name,
		Query: query,
		Args: arg,
	}
}

func (s *Statement) GetLastInsertedId() int64 {
	return s.InsertedId
}
