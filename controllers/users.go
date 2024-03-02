package controllers

import (
	"fmt"
	"net/http"

	"github.com/progsamdev/coursescalhoun/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "user created: %+v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("name_test")
	if err != nil {
		u.Templates.New.Execute(w, r, nil)
		return
	}
	fmt.Fprintf(w, "cookie value: %s\n", cookie.Value)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "name_test",
		Value:    "value_test",
		Path:     "/signin",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "User authenticated: %+v", user)
}
