package member_domain

import "github.com/stretchr/testify/mock"

type MockMemberRepository struct {
	mock.Mock
}

func (m *MockMemberRepository) GetById(id MemberId) (*Member, error) {
	ret := m.Called(id)
	return ret.Get(0).(*Member), ret.Error(1)
}

func (m *MockMemberRepository) FindByName(name string) ([]*Member, error) {
	ret := m.Called(name)
	return ret.Get(0).([]*Member), ret.Error(1)
}

func (m *MockMemberRepository) Create(member *Member) (MemberId, error) {
	ret := m.Called(member)
	return ret.Get(0).(MemberId), ret.Error(1)
}

func (m *MockMemberRepository) Update(member *Member) error {
	ret := m.Called(member)
	return ret.Error(0)
}

func (m *MockMemberRepository) Delete(member *Member) error {
	ret := m.Called(member)
	return ret.Error(0)
}
