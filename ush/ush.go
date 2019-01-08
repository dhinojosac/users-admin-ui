package ush

import (
	"time"
)

type User struct {
	name        string
	surname1    string
	surname2    string
	age         int
	birthday    time.Time
	description string
	entryDate   time.Time
}

func CreateUser(n, s1, s2 string) *User {
	nu := new(User)
	nu.name = n
	nu.surname1 = s1
	nu.surname2 = s2
	return nu
}

func (u *User) getName() string {
	return u.name
}
