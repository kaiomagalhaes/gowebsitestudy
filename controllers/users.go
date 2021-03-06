package controllers

import (
	"fmt"
	"gowebsitestudy/models"
	"gowebsitestudy/views"
	"net/http"
)

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "users/new"),
		LoginView:   views.NewView("bootstrap", "users/login"),
		UserService: us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Printf("login email %s", form.Email, form.Password)
	fmt.Printf("login password %s", form.Password)
	user := u.UserService.Authenticate(form.Email, form.Password)
	fmt.Fprintln(w, user)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.UserService.Create(user); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, user)
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	models.UserService
}
