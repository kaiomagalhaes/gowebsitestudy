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

func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "users/new"),
		UserService: us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := &models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.UserService.Create(user); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, user)
}

type Users struct {
	NewView *views.View
	models.UserService
}
