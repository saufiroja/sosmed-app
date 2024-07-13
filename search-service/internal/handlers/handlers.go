package handlers

import "github.com/saufiroja/sosmed/search-service/internal/handlers/query"

type Handler struct {
	Query *query.Query
}

func NewHandler(query *query.Query) *Handler {
	return &Handler{Query: query}
}
