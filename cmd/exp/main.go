package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "John Smith",
		Age:  24,
		Bio:  `<script>alert("you have been hacked!!!")</script>`,
		Meta: UserMeta{Visits: 4},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
