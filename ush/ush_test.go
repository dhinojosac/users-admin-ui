package ush

import "testing"

func TestCreateUser(t *testing.T) {
	user := CreateUser("Diego", "Orlando", "Hinojosa", "Cordova")
	if user.GetName() != "Diego" {
		t.Error("Name incorrect")
	}
	if user.GetFullName() != "Diego Orlando Hinojosa Cordova" {
		t.Error("Name incorrect")
	}

}
