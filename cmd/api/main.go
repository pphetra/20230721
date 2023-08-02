package main

import (
	"github.com/gofiber/fiber/v2"

	"taejai/internal/infra"
	infra_eventbus_watermill "taejai/internal/infra/event_bus/watermill"
	infra_unit_of_work "taejai/internal/infra/unit_of_work"
	member_event_handlers "taejai/internal/member/app/event_handlers"
	member_infra_api_handlers "taejai/internal/member/infra/api_handlers"
	shared_app "taejai/internal/shared/app"

	_ "taejai/internal/infra/mail"
	_ "taejai/internal/member/infra"

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
	eventBus.SetCommandExecutor(unitOfWork.GetCommandExecutor())

	injector := shared_app.Injector{
		UnitOfWork: unitOfWork,
		EventBus:   eventBus,
	}

	// register event handers
	eventBus.Subscribe(member_event_handlers.NewSendGreetingMailHandler())

	app := fiber.New()

	app.Post("/register/individual", member_infra_api_handlers.NewIndividualRegisterHandler(injector))

	app.Listen(":3000")

}
