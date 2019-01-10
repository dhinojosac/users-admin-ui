package ush

import (
	"strconv"
	"time"
)

var fisconditionMap = map[int]string{
0: "SÍNDROME DOLOROSO DE ORIGEN TRAUMÁTICO"
1 : "SÍNDROME DOLOROSO DE ORIGEN NO TRAUMÁTICO"
2: "ARTROSIS LEVE Y MODERADA DE RODILLA Y CADERA"
3 : "SECUELA DE ACCIDENTE CEREBRO VASCULAR (ACV)"
4: "SECUELAS DE TRAUMATISMO ENCEFALO CRANEANO (TEC)"
5 : "SECUELA DE TRAUMATISMO RAQUIMEDULAR (TRM)"
6: "SECUELA QUEMADURA"
7: "ENFERMEDAD DE PARKINSON"
8 : "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR EN MENOR DE 20 AÑOS CONGÉNITO"
9 : "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR EN MENOR DE 20 AÑOS ADQUIRIDO"
10 : "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR  EN MAYOR DE 20 AÑOS"
11 : "AMPUTACIÓN POR PIE DIABETICO"
12 : "OTROS"
}

var visconditionMap = map[int]string{
	0: "CONGÉNITO"
	1 : "ADQUIRIDO"
	2: "OTROS"

}

var audconditionMap = map[int]string{
		0: "CONGÉNITO"
		1 : "ADQUIRIDO"
		2: "OTROS"
}

type User struct {
	name         string
	name2        string
	surname1     string
	surname2     string
	age          int
	fiscondition int
	viscondition int
	audcondition int
	description  string
	entryDate    time.Time
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

func (u *User) SetFisCondition(d string) {
	u.fiscondition = d
}

func (u *User) GetFisCondition() string {
	return u.fiscondition
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
	return strconv.Itoa(u.age)
}
