package main

import (
	"time"

	basicsApp "github.com/yogendraatluri990/go_basics.git/app"
)

func main() {

	user := &basicsApp.User{
		Name: "Yogendra",
		Age:  30,
		Position: basicsApp.Position{
			Name:  "Software Engineer",
			Level: "Senior",
		},
	}

	signup := &basicsApp.SignUp{
		Name:     "Yogendra",
		Email:    "yoge123@gmail.com",
		Password: "password123",
	}

	go basicsApp.HandleLogInfo(signup)
	time.Sleep(time.Millisecond)

	go basicsApp.HandleLogInfo(user)
	time.Sleep(time.Minute * 2)
}
