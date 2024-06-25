package handlers

import "github.com/saufiroja/sosmed-app/auth-service/internal/handlers/command"

type Handler struct {
	Command *command.Command
}

func NewHandler(command *command.Command) *Handler {
	return &Handler{
		Command: command,
	}
}
