package model

import "time"

type User struct {
	FirstName string
	LastName  string
	Birthdate time.Time
	IsActive  bool
}

func (u *User) NewModified(user *User) {
	user.FirstName = "Morteza"
}

func Deactive(user *User) {
	user.IsActive = false
}
