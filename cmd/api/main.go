package main

import (
	"github.com/gofiber/fiber/v2"

	"taejai/internal/infra"
	infra_eventbus_watermill "taejai/internal/infra/event_bus/watermill"
	infra_mail "taejai/internal/infra/mail"
	infra_unit_of_work "taejai/internal/infra/unit_of_work"
	member_event_handlers "taejai/internal/member/app/event_handlers"
	member_infra_api_handlers "taejai/internal/member/infra/api_handlers"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"

	member_infra "taejai/internal/member/infra"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	db, err := infra.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventBus := infra_eventbus_watermill.NewGoChannelEventBus()
	unitOfWork := infra_unit_of_work.NewUnitOfWork(db, eventBus)

	shared_app.ServiceRegistry.Register("mail_service", infra_mail.NewNoOPMailService())
	shared_domain.RepositoryRegistry.Register("member", member_infra.NewMemberRepositoryFactory())

	// register event handers
	eventBus.Subscribe(member_event_handlers.NewIndividualMemberRegisteredHandler(unitOfWork))
	eventBus.Subscribe(member_event_handlers.NewGreetingMailSendHandler(unitOfWork))

	app := fiber.New()

	app.Post("/register/individual", member_infra_api_handlers.NewIndividualRegisterHandler(unitOfWork))

	app.Listen(":3000")

}
