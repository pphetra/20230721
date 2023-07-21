package member_app_commands

import shared_app "taejai/internal/shared/app"

const SendGreetingMailCommandName = "SendGreetingMailCommand"

type SendGreetingMailCommand struct {
	MemberId string
	FullName string
	Email    string
}

// implement Command interface
func (c SendGreetingMailCommand) GetCommandName() string {
	return SendGreetingMailCommandName
}

// implement Command interface
func (c SendGreetingMailCommand) Execute(store shared_app.UnitOfWorkStore) (interface{}, error) {
	// TODO: implement this
	return nil, nil
}
