package member_app_commands_test

import (
	"errors"
	member_app_commands "taejai/internal/member/app/commands"
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterIndividualCommand_SuccessWithPublishEvent(t *testing.T) {
	mockUnitOfWorkStore := &shared_app.MockUnitOfWorkStore{}
	mockUnitOfWork := &shared_app.MockUnitOfWork{
		UnitOfWorkStore: mockUnitOfWorkStore,
	}
	mockMemberRepository := &member_domain.MockMemberRepository{}

	memberId := member_domain.MemberId(1)
	member := member_domain.Member{
		Id:    memberId,
		Name1: "John",
		Name2: "Doe",
		Type:  member_domain.Individual,
		Address: member_domain.Address{
			Line1:      "123 Main St",
			Line2:      "Apt 4",
			PostalCode: "12345",
		},
	}

	command := member_app_commands.RegisterIndividualMemberCommand{
		FirstName:         member.Name1,
		LastName:          member.Name2,
		AddressLine1:      member.Address.Line1,
		AddressLine2:      member.Address.Line2,
		AddressPostalCode: member.Address.PostalCode,
	}

	mockUnitOfWorkStore.On("GetRepository", "member").Return(mockMemberRepository)
	mockMemberRepository.On("Create", mock.Anything).Return(memberId, nil)
	mockMemberRepository.On("GetById", memberId).Return(&member, nil)
	mockUnitOfWorkStore.On("AddEventToPublish", mock.Anything).Return(nil)

	executor := shared_app.NewCommandExecutor(mockUnitOfWork)
	retId, err := executor.Execute(command)
	assert.NoError(t, err)
	assert.Equal(t, memberId, retId)

	mockUnitOfWorkStore.AssertExpectations(t)
	mockUnitOfWork.AssertExpectations(t)
	mockMemberRepository.AssertExpectations(t)
}

func TestRegisterIndividualCommand_FailedWithoutPublishEvent(t *testing.T) {
	mockUnitOfWorkStore := &shared_app.MockUnitOfWorkStore{}
	mockUnitOfWork := &shared_app.MockUnitOfWork{
		UnitOfWorkStore: mockUnitOfWorkStore,
	}
	mockMemberRepository := &member_domain.MockMemberRepository{}

	memberId := member_domain.MemberId(1)
	member := member_domain.Member{
		Id:    memberId,
		Name1: "John",
		Name2: "Doe",
		Type:  member_domain.Individual,
		Address: member_domain.Address{
			Line1:      "123 Main St",
			Line2:      "Apt 4",
			PostalCode: "12345",
		},
	}

	command := member_app_commands.RegisterIndividualMemberCommand{
		FirstName:         member.Name1,
		LastName:          member.Name2,
		AddressLine1:      member.Address.Line1,
		AddressLine2:      member.Address.Line2,
		AddressPostalCode: member.Address.PostalCode,
	}

	mockUnitOfWorkStore.On("GetRepository", "member").Return(mockMemberRepository)
	mockMemberRepository.On("Create", mock.Anything).Return(member_domain.MemberId(0), errors.New("error"))

	executor := shared_app.NewCommandExecutor(mockUnitOfWork)
	_, err := executor.Execute(command)
	assert.Error(t, err)

	mockUnitOfWorkStore.AssertExpectations(t)
	mockUnitOfWork.AssertExpectations(t)
	mockMemberRepository.AssertExpectations(t)
}
