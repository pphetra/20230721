package member_infra_api_handlers

import (
	member_app_commands "taejai/internal/member/app/commands"
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"

	"github.com/gofiber/fiber/v2"
)

func NewIndividualRegisterHandler(commandExecutor shared_app.CommandExecutor) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse request body to json map
		var body map[string]string
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// create command
		cmd := member_app_commands.RegisterIndividualMemberCommand{
			FirstName:         body["first_name"],
			LastName:          body["last_name"],
			AddressLine1:      body["address_line_1"],
			AddressLine2:      body["address_line_2"],
			AddressPostalCode: body["address_postal_code"],
			Email:             body["email"],
		}

		// execute command
		ret, err := commandExecutor.Execute(cmd)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		memberId := ret.(member_domain.MemberId)

		return c.JSON(fiber.Map{
			"member_id": memberId,
		})
	}
}
