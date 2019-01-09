package ush

import (
	"time"
)

type User struct {
	name        string
	name2       string
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
	nu.entryDate = time.Now()
	return nu
}

func (u *User) SetDescription(d string) {
	u.description = d
}

func (u *User) GetDescription() string {
	return u.description
}

func (u *User) SetName(n string) {
	u.name = n
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetFullName() string {
	return u.name + " " + u.name2 + " " + u.surname1 + " " + u.surname2
}

func (u *User) SetAge(a int) {
	u.age = a
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetAgeStr() string {
	return u.entryDate.Format("2006-01-02 15:04:05")
}
