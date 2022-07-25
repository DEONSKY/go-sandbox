package router

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/handler"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/DEONSKY/go-sandbox/utils/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                    = config.SetupDatabaseConnection()
	envVariables    config.EnvironmentVariables = config.LoadEnvVariables()
	bookRepository  repository.BookRepository   = repository.NewBookRepository(db)
	issueRepository repository.IssueRepository  = repository.NewIssueRepository(db)
	issueService    service.IssueService        = service.NewIssueService(issueRepository)
	issueHandler    handler.IssueHandler        = handler.NewIssueHandler(issueService)
	jwtService      service.JWTService          = service.NewJWTService(envVariables)
	bookService     service.BookService         = service.NewBookService(bookRepository)
	authHandler     handler.AuthHandler         = handler.NewAuthController(jwtService)
	bookHandler     handler.BookHandler         = handler.NewBookController(bookService, jwtService)
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "[${time}] ${status} ${latency} ${method} ${path} - ${pid} - ${locals:requestid}\n",
	}))

	app.Get("/docs/*", swagger.HandlerDefault)

	root := app.Group("/api")

	authRoutes := root.Group("/auth")

	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)

	/*
		bookRoutes := root.Group("/books", middleware.Protected())

		bookRoutes.Get("/", bookHandler.All)
		bookRoutes.Post("/", bookHandler.Insert)
		bookRoutes.Get("/:id", bookHandler.FindByID)
		bookRoutes.Put("/:id", bookHandler.Update)
		bookRoutes.Delete("/:id", bookHandler.Delete)
	*/

	issueRoutes := root.Group("/issue", middleware.Protected())
	issueRoutes.Post("/", issueHandler.InsertIssue)
	issueRoutes.Get("/", issueHandler.GetIssues)
	issueRoutes.Get("/kanban/", issueHandler.GetIssuesKanban)
	issueRoutes.Put("/add-issue-dependency/:issue_id/:dependent_issue_id", issueHandler.InsertDependentIssueAssociation)
	issueRoutes.Put("/assignie-user/:issue_id/:user_id", issueHandler.AssignieIssueToUser)

	issueCommentRoutes := root.Group("/issue-comment", middleware.Protected())
	issueCommentRoutes.Post("/", handler.AddIssueComment)

	projectRoutes := root.Group("/project", middleware.Protected())
	projectRoutes.Post("/", handler.InsertProject)
	projectRoutes.Get("/sidenav-options/", handler.GetProjectsByUserId)

	subjectRoutes := root.Group("/subject", middleware.Protected())
	subjectRoutes.Post("/", handler.InsertSubject)
	subjectRoutes.Put("/:subject_id/:user_id", handler.InsertUserToSubject)
	subjectRoutes.Get("/user-options/:subject_id", handler.GetSubjectsUsersOptions)

	app.Use(func(c *fiber.Ctx) error {
		return utils.ReturnErrorResponse(404, "Service not exists", []string{})
	})
	return app
}
