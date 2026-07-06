package app

import "fmt"

type Position struct {
	Name  string
	Level string
}

type User struct {
	Name     string
	Age      int
	Position Position
}

type SignUp struct {
	Name     string
	Email    string
	Password string
}

type Task string

type Tasks struct {
	Title string
	Task  Task
}

type LoggerInfo interface {
	LogInfo()
}

func (u *User) LogInfo() {
	fmt.Printf("Name: %s Age: %d Position: %s Level: %s\n", u.Name, u.Age, u.Position.Name, u.Position.Level)
}

func (s *SignUp) LogInfo() {
	fmt.Printf("Name: %s, Email: %s,  Password: %s\n", s.Name, s.Email, s.Password)
}

func HandleLogInfo(i LoggerInfo) {
	i.LogInfo()
}
