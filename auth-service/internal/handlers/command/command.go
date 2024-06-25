package command

type Command struct {
	RegisterCommandHandler           *RegisterCommandHandler
	LoginCommandHandler              *LoginCommandHandler
	RefreshTokenCommandHandler       *RefreshTokenCommandHandler
	GoogleAuthCommandHandler         *GoogleAuthCommandHandler
	GoogleAuthCallbackCommandHandler *GoogleAuthCallbackCommandHandler
}

func NewCommand(
	registerCommandHandler *RegisterCommandHandler,
	loginCommandHandler *LoginCommandHandler,
	refreshTokenCommandHandler *RefreshTokenCommandHandler,
	googleAuthCommandHandler *GoogleAuthCommandHandler,
	googleAuthCallbackCommandHandler *GoogleAuthCallbackCommandHandler,
) *Command {
	return &Command{
		RegisterCommandHandler:           registerCommandHandler,
		LoginCommandHandler:              loginCommandHandler,
		RefreshTokenCommandHandler:       refreshTokenCommandHandler,
		GoogleAuthCommandHandler:         googleAuthCommandHandler,
		GoogleAuthCallbackCommandHandler: googleAuthCallbackCommandHandler,
	}
}
