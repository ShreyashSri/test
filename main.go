package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID            string
	Country           string
	CreatedAt         string
	Curriculum        string
	Email             string
	FullName          string
	LastLogin         string
	Role              string
	SignupIP          string
	UserTourCompleted bool
}

var users = []User{
	{
		UserID:            "1",
		Country:           "US",
		CreatedAt:         "2022-01-01T10:00:00Z",
		Curriculum:        "Math",
		Email:             "user1@example.com",
		FullName:          "Alice Smith",
		LastLogin:         "2022-06-01T12:00:00Z",
		Role:              "student",
		SignupIP:          "192.168.1.1",
		UserTourCompleted: true,
	},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8081")
}
