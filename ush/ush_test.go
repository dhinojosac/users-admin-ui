package ush

import (
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	user := CreateUser("Diego", "Orlando", "Hinojosa", "Cordova")
	if user.GetName() != "Diego" {
		t.Error("Name incorrect")
	}

	if user.GetFullName() != "Diego Orlando Hinojosa Cordova" {
		t.Error("Full name incorrect")
	}

	user.SetRut("17000000-5")
	if user.GetRut() != "17000000-5" {
		t.Error("Rut incorrect")
	}

	user.SetAge(30)
	if user.GetAge() != 30 {
		t.Error("Age int incorrect")
	}
	if user.GetAgeStr() != "30" {
		t.Error("Age str incorrect")
	}

	user.SetGender(0)
	if user.GetGender() != 0 {
		t.Error("Gender int incorrect")
	}
	if user.GetGenderStr() != "Hombre" {
		t.Error("Gender str incorrect")
	}

	user.SetGender(1)
	if user.GetGender() != 1 {
		t.Error("Gender int incorrect")
	}
	if user.GetGenderStr() != "Mujer" {
		t.Error("Gender str incorrect")
	}

	user.SetPrevision(1)
	if user.GetPrevision() != 1 {
		t.Error("Prevision int incorrect")
	}
	if user.GetPrevisionStr() != "FONASA B" {
		t.Error("Prevision str incorrect")
	}

	user.SetContact("99055095")
	if user.GetContact() != "99055095" {
		t.Error("Contact str incorrect")
	}

	user.SetDescription("This is a description")
	if user.GetDescription() != "This is a description" {
		t.Error("Description str incorrect")
	}

	now := time.Now()
	user.SetEntryDate(now)
	if user.GetEntryDate() != now {
		t.Error("Entry date time incorrect")
	}
	if user.GetEntryDateStr() != now.Format("2006-01-02 15:04:05") {
		t.Error("Entry date str time incorrect")
	}
}
