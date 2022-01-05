package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-sandbox/config"
	"example.com/go-sandbox/controller"
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
	authController controller.AuthController = controller.NewAuthController()
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
