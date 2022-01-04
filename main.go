package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	fmt.Println("Go ORM Tutorial")

	InitialMigration()
	handleRequests()
}
