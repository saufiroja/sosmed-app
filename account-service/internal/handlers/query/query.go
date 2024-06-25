package query

type Query struct {
	GetAccountByUsernameAndEmailHandlerQuery *GetAccountByUsernameAndEmailHandlerQuery
}

func NewQuery(
	getAccountByUsernameAndEmailHandlerQuery *GetAccountByUsernameAndEmailHandlerQuery,
) *Query {
	return &Query{
		GetAccountByUsernameAndEmailHandlerQuery: getAccountByUsernameAndEmailHandlerQuery,
	}
}
