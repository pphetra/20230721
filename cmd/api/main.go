package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"

	infra_event_bus "taejai/internal/infra/event_bus"
	infra_mail "taejai/internal/infra/mail"
	infra_unit_of_work "taejai/internal/infra/unit_of_work"
	member_event_handlers "taejai/internal/member/app/event_handlers"
	member_infra_api_handlers "taejai/internal/member/infra/api_handlers"
	shared_app "taejai/internal/shared/app"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initConnectionString() string {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}
	return "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
}
func main() {
	godotenv.Load()

	db, err := sql.Open("postgres", initConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventBus := infra_event_bus.NewGoChannelEventBus()
	unitOfWork := infra_unit_of_work.NewPostgresUnitOfWork(db, eventBus)
	commandExecutor := shared_app.NewCommandExecutor(unitOfWork)
	eventBus.SetCommandExecutor(commandExecutor)
	mailService := infra_mail.NewMailService()

	// register event handers
	eventBus.Subscribe(member_event_handlers.NewSendGreetingMailHandler(mailService))

	app := fiber.New()

	app.Post("/register/individual", member_infra_api_handlers.NewIndividualRegisterHandler(*commandExecutor))

	app.Listen(":3000")

}
