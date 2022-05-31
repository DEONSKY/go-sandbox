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
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authHandler    handler.AuthHandler       = handler.NewAuthController(authService, jwtService)
	bookHandler    handler.BookHandler       = handler.NewBookController(bookService, jwtService)
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	root := app.Group("/api", logger.New())

	authRoutes := root.Group("/auth")

	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)

	bookRoutes := root.Group("/books", middleware.Protected())

	bookRoutes.Get("/", bookHandler.All)
	bookRoutes.Post("/", bookHandler.Insert)
	bookRoutes.Get("/:id", bookHandler.FindByID)
	bookRoutes.Put("/:id", bookHandler.Update)
	bookRoutes.Delete("/:id", bookHandler.Delete)

	return app
}
