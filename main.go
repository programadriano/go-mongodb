package main

import (
	// Standard library packages
	"fmt"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/programadriano/go-mongodb/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	// Create a new user
	r.POST("/user", uc.CreateUser)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	fmt.Println("Running in port:3030")
	http.ListenAndServe("localhost:3030", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://admin:Rfle5Bmbooqr@localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	fmt.Println("s:", s)
	// Deliver session
	return s
}
