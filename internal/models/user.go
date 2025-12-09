package models

import "time"

type User struct {
	UserID            string    `json:"userId" firestore:"userId"`
	Country           string    `json:"country" firestore:"country"`
	CreatedAt         time.Time `json:"createdAt" firestore:"createdAt"`
	Curriculum        string    `json:"curriculum" firestore:"curriculum"`
	Email             string    `json:"email" firestore:"email"`
	FullName          string    `json:"fullName" firestore:"fullName"`
	LastLogin         time.Time `json:"lastLogin" firestore:"lastLogin"`
	Role              string    `json:"role" firestore:"role"`
	SignupIP          string    `json:"signupIP" firestore:"signupIP"`
	UserTourCompleted bool      `json:"userTourCompleted" firestore:"userTourCompleted"`
}
