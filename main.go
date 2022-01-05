package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-sandbox/config"
	"example.com/go-sandbox/controller"
	"example.com/go-sandbox/repository"
	"example.com/go-sandbox/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := gin.Default()
	myRouter.GET("/users", AllUsers)
	myRouter.POST("/user/", NewUser)
	myRouter.DELETE("/user/:name", DeleteUser)
	myRouter.PUT("/user/:name", UpdateUser)
	myRouter.Run()
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)

	fmt.Println("Go ORM Tutorial")

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()

	//handleRequests()

}
