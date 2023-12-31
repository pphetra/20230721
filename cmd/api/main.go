package main

import (
	"database/sql"
	"fmt"
	"os"

	infra_event_bus "taejai/internal/infra/event_bus"
	infra_unit_of_work "taejai/internal/infra/unit_of_work"

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

	db, err := sql.Open("postgres", initConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventBus := infra_event_bus.NewChannelEventBus()
	unitOfWork := infra_unit_of_work.NewPostgresUnitOfWork(db, eventBus)

	// TODO use unitOfWork
	fmt.Println(unitOfWork)

}
