package command

type Command struct {
	InsertUserCommandHandler *InsertUserCommandHandler
}

func NewCommand(
	insertUserCommandHandler *InsertUserCommandHandler,
) *Command {
	return &Command{
		InsertUserCommandHandler: insertUserCommandHandler,
	}
}
