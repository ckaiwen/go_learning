package main

import "fmt"

type People interface {
	ReturnName() string
}

type Role interface {
	People
	ReturnRole()string
}
type Student struct {
	Name string
	Role string
}

func (s *Student)ReturnRole() string {
	return s.Role
}
func (s *Student)ReturnName()string  {
	return s.Name
}

func main() {
   student1:=Student{Name: "cfk",Role: "学生"}
   var role Role
   role= &student1
   fmt.Println(role.ReturnName())
   fmt.Println(role.ReturnRole())
}
