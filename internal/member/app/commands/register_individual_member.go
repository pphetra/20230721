package member_app_commands

import (
	"errors"
	member_domain "taejai/internal/member/domain"
	member_domain_events "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
)

const RegisterIndividualMemberCommandName = "register_individual_member"

type RegisterIndividualMemberCommand struct {
	FirstName         string
	LastName          string
	AddressLine1      string
	AddressLine2      string
	AddressPostalCode string
}

func (c RegisterIndividualMemberCommand) GetCommandName() string {
	return RegisterIndividualMemberCommandName
}

func (c RegisterIndividualMemberCommand) Execute(store shared_app.UnitOfWorkStore) (interface{}, error) {
	memberRepository := store.GetRepository("member").(member_domain.MemberRepository)
	if memberRepository == nil {
		return 0, errors.New("member repository not found")
	}

	// what is the error handling strategy here?
	address, err := member_domain.NewAddress(
		c.AddressLine1,
		c.AddressLine2,
		c.AddressPostalCode,
	)
	if err != nil {
		return 0, err
	}

	member, err := member_domain.NewIndividualMember(
		c.FirstName,
		c.LastName,
		address,
	)
	if err != nil {
		return 0, err
	}
	// should we pass member by value or by reference?
	id, err := memberRepository.Create(&member)
	if err != nil {
		return 0, err
	}

	createdMember, err := memberRepository.GetById(id)
	if err != nil {
		return 0, err
	}

	// publish RegisteredEvent
	event := member_domain_events.NewIndividualMemberRegisteredEvent(*createdMember)
	err = store.AddEventToPublish(event)
	if err != nil {
		return 0, err
	}

	return id, err

}
