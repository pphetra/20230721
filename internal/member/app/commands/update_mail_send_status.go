package member_app_commands

import (
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
	"time"
)

const UpdateMailSendCommandName = "UpdateMailSendCommand"

type UpdateMailSendCommand struct {
	MemberId member_domain.MemberId
	Time     time.Time
}

func (c UpdateMailSendCommand) GetCommandName() string {
	return UpdateMailSendCommandName
}

func (c UpdateMailSendCommand) Execute(store shared_app.UnitOfWorkStore) (interface{}, error) {
	memberRepo := store.GetRepository("member").(member_domain.MemberRepository)

	member, err := memberRepo.GetById(c.MemberId)
	if err != nil {
		return nil, err
	}

	member.MailSend = &c.Time

	err = memberRepo.Update(member)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
