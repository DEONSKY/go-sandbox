package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/DEONSKY/go-sandbox/docs"
	"github.com/joho/godotenv"

	"github.com/DEONSKY/go-sandbox/config"
	router "github.com/DEONSKY/go-sandbox/routes"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {

	defer config.CloseDatabaseConnection(db)

	fmt.Println("Go ORM Tutorial")

	//r := gin.Default()

	app := router.New()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
