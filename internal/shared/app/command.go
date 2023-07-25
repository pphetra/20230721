package shared_app

type Command interface {
	GetCommandName() string
	Execute(UnitOfWorkStore) (interface{}, error)
}

type CommandExecutor struct {
	uow UnitOfWork
}

func NewCommandExecutor(uow UnitOfWork) *CommandExecutor {
	return &CommandExecutor{
		uow: uow,
	}
}

// command will execute in a transaction
// if command is successful, it will commit the transaction
// we inject publish function to command so that command can publish event
func (c *CommandExecutor) Execute(command Command) (interface{}, error) {
	return c.uow.DoInTransaction(command.Execute)
}
