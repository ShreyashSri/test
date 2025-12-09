package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userID            string
	country           string
	createdAt         string
	curriculum        string
	email             string
	fullName          string
	lastLogin         string
	role              string
	signupIP          string
	userTourCompleted bool
}

var users = []User{
	{
		userID:            "1",
		country:           "US",
		createdAt:         "2022-01-01T10:00:00Z",
		curriculum:        "Math",
		email:             "user1@example.com",
		fullName:          "Alice Smith",
		lastLogin:         "2022-06-01T12:00:00Z",
		role:              "student",
		signupIP:          "192.168.1.1",
		userTourCompleted: true,
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
