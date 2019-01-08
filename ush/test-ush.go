package main

import (
	"fmt"
	"strconv"
	"ush"
)

func main() {
	user := ush.CreateUser("Diego", "Hinojosa", "Córdova")
	desc := "Diagnosticado con protruciones lumbares en L3-L4 y L5-S1, sin compromiso nervioso."
	user.SetDescription(desc)
	user.SetAge(30)

	uname := user.GetFullName()
	fmt.Println("Nombre: " + uname)
	fmt.Println("Edad: " + strconv.Itoa(user.GetAge()))
	fmt.Println("Descripción: " + user.GetDescription())
	fmt.Println("Fecha ingreso: " + user.GetAgeStr())
}
