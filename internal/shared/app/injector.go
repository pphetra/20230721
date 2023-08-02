package shared_app

type Injector struct {
	CommandExecutor *CommandExecutor
	UnitOfWork      UnitOfWork
	EventBus        EventBus
}
