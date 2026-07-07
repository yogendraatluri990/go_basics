package main

import (
	"fmt"

	basicsApp "github.com/yogendraatluri990/go_basics.git/app"
)

func main() {
	tasks := make(chan basicsApp.TaskEvent)
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

	go basicsApp.HandleLogInfo(signup, tasks)

	tasks <- basicsApp.TaskEvent{
		Title: basicsApp.TaskUserSignup,
		Task:  basicsApp.TaskPending,
	}

	signUpEvents := <-tasks
	fmt.Printf("TaskEvents: -> Task.Title: %s  and Task.Task: %s\n", signUpEvents.Title, signUpEvents.Task)
	if signUpEvents.Task == basicsApp.TaskDone && signUpEvents.Title == basicsApp.TaskSignupComplete {
		tasks <- basicsApp.TaskEvent{
			Title: basicsApp.TaskUser,
			Task:  basicsApp.TaskPending,
		}
		go basicsApp.HandleLogInfo(user, tasks)

	}

	userProfilTask := <-tasks
	fmt.Printf("TaskEvents: -> Task.Title: %s  and Task.Task: %s\n", userProfilTask.Title, userProfilTask.Task)
}
