package handlers

import (
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers/command"
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers/query"
)

type Handler struct {
	Command *command.Command
	Query   *query.Query
}

func NewHandler(command *command.Command, query *query.Query) *Handler {
	return &Handler{
		Command: command,
		Query:   query,
	}
}
