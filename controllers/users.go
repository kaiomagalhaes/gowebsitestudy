package controllers

import (
	"fmt"
	"gowebsitestudy/views"
	"net/http"
)

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	parseForm(r, &form)

	fmt.Fprintln(w, form)
}

type Users struct {
	NewView *views.View
}
