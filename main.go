package main

import (
	"api/middlewares"
	"api/routes/auth"
	"api/routes/birth_record"
	"api/routes/cattle"
	"api/routes/death_record"
	"api/routes/illness_record"
	"api/routes/insemination_record"
	"api/routes/milking_record"
	"api/routes/weight_record"
	"api/utils"

	_ "api/docs"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title						Cattle Farm Management API
// @version					1.0
// @description				A robust REST API for managing cattle farms
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.email				zxselimcan@icloud.com
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
// @host						localhost:4000
// @BasePath					/
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func init() {
	err := godotenv.Load()
	if err != nil {
		panic("ERROR_DOTENV_FILE")
	}
	utils.ConnectSqlite()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	// Add Swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Auth

	auth_api := app.Group("/auth")

	auth_api.Post("/login", auth.Login)
	auth_api.Post("/register", auth.Register)
	auth_api.Post("/email-verification", auth.EmailVerification)

	api := app.Group("/api", middlewares.VerifyJwtKey(false)) // middlewares.VerifyJwtKey(false)

	cattle_api := api.Group("/cattle")
	// get my cattles
	cattle_api.Get("/", cattle.GetMyCattles)
	// get my milkable cattles
	cattle_api.Get("/milkable", cattle.GetMyMilkableCattles)
	// get my inseminated cattles
	cattle_api.Get("/inseminated", insemination_record.GetMyInseminatedCattles)
	// get my pregnant cattles
	cattle_api.Get("/pregnant", insemination_record.GetMyPregnantCattles)
	// get my not pregnant cattles
	cattle_api.Get("/non-pregnant", insemination_record.GetMyNonPregnantCattles)
	// get my dead cattles
	cattle_api.Get("/dead", death_record.GetMyDeadCattles)

	// add new cattle
	cattle_api.Post("/", cattle.NewCattle)

	cattle_owner_only_api := cattle_api.Group("/:cattle_uuid", middlewares.CattleOwnerOrAdminOnly())

	// get my cattle info
	cattle_owner_only_api.Get("/", cattle.GetCattleByUUID)
	// get cattle insemination records
	cattle_owner_only_api.Get("/insemination-records", insemination_record.GetInseminationRecordsByCattleUUID)
	// get cattle birth records
	cattle_owner_only_api.Get("/birth-records", birth_record.GetBirthRecordsByCattleUUID)
	// get cattle illness records
	cattle_owner_only_api.Get("/illness-records", illness_record.GetIllnessRecordsByCattleByUUID)
	// get cattle milking records
	cattle_owner_only_api.Get("/milking-records", milking_record.GetMilkingRecordsByCattleByUUID)
	// get cattle weight records
	cattle_owner_only_api.Get("/weight-records", weight_record.GetWeightRecordsByCattleByUUID)

	// new cattle insemination record
	cattle_owner_only_api.Post("/insemination-records", insemination_record.NewInseminationRecord)
	// new cattle pregnancy record
	cattle_owner_only_api.Post("/insemination-records/new-pregnancy", insemination_record.NewPregnancy)
	// new failed pregnancy record
	cattle_owner_only_api.Post("/insemination-records/failed-pregnancy", insemination_record.FailedPregnancy)
	// new cattle birth record
	cattle_owner_only_api.Post("/birth-records", birth_record.NewBirthRecord)
	// new cattle illness record
	cattle_owner_only_api.Post("/illness-records", illness_record.NewIllnessRecord)
	// new cattle milking record
	cattle_owner_only_api.Post("/milking-records", milking_record.NewMilkingRecord)
	// new cattle weight record
	cattle_owner_only_api.Post("/weight-records", weight_record.NewWeightRecord)
	// new cattle death record
	cattle_owner_only_api.Post("/death-records", death_record.NewDeathRecord)

	if err := app.Listen("127.0.0.1:4000"); err != nil {
		log.Fatal(err)
	}

}
