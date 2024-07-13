package query

type Query struct {
	GetUserByUsernameQueryHandler *GetUserByUsernameQueryHandler
	GetAllUsersQueryHandler       *GetAllUsersQueryHandler
}

func NewQuery(
	getUserByUsernameQueryHandler *GetUserByUsernameQueryHandler,
	getAllUsersQueryHandler *GetAllUsersQueryHandler,
) *Query {
	return &Query{
		GetUserByUsernameQueryHandler: getUserByUsernameQueryHandler,
		GetAllUsersQueryHandler:       getAllUsersQueryHandler,
	}
}
