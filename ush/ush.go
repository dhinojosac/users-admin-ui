package ush

import (
	"strconv"
	"time"
)

var FisconditionMap = map[int]string{
	0:  "SÍNDROME DOLOROSO DE ORIGEN TRAUMÁTICO",
	1:  "SÍNDROME DOLOROSO DE ORIGEN NO TRAUMÁTICO",
	2:  "ARTROSIS LEVE Y MODERADA DE RODILLA Y CADERA",
	3:  "SECUELA DE ACCIDENTE CEREBRO VASCULAR (ACV)",
	4:  "SECUELAS DE TRAUMATISMO ENCEFALO CRANEANO (TEC)",
	5:  "SECUELA DE TRAUMATISMO RAQUIMEDULAR (TRM)",
	6:  "SECUELA QUEMADURA",
	7:  "ENFERMEDAD DE PARKINSON",
	8:  "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR EN MENOR DE 20 AÑOS CONGÉNITO",
	9:  "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR EN MENOR DE 20 AÑOS ADQUIRIDO",
	10: "OTRO DÉFICIT SECUNDARIO CON COMPROMISO NEUROMUSCULAR  EN MAYOR DE 20 AÑOS",
	11: "AMPUTACIÓN POR PIE DIABETICO",
	12: "OTROS",
}

var VisconditionMap = map[int]string{
	0: "CONGÉNITO",
	1: "ADQUIRIDO",
	2: "OTROS",
}

var AudconditionMap = map[int]string{
	0: "CONGÉNITO",
	1: "ADQUIRIDO",
	2: "OTROS",
}

var GenderMap = map[int]string{
	0: "Hombre",
	1: "Mujer",
}

var PrevisionMap = map[int]string{
	0: "FONASA A",
	1: "FONASA B",
	2: "FONASA C",
	3: "FONASA D",
	4: "PRAIS",
}

type User struct {
	rut          string
	name         string
	name2        string
	surname1     string
	surname2     string
	age          int
	gender       int
	fiscondition int
	viscondition int
	audcondition int
	prevision    int
	contact      string
	description  string
	entryDate    time.Time
}

func CreateUser(n1, n2, s1, s2 string) *User {
	nu := new(User)
	nu.name = n1
	nu.name2 = n2
	nu.surname1 = s1
	nu.surname2 = s2
	nu.entryDate = time.Now()
	return nu
}

//Rut
func (u *User) SetRut(d string) {
	u.rut = d
}

func (u *User) GetRut() string {
	return u.rut
}

//Name
func (u *User) SetName(d string) {
	u.name = d
}

func (u *User) GetName() string {
	return u.name
}

//Second name
func (u *User) SetName2(d string) {
	u.name2 = d
}

func (u *User) GetName2() string {
	return u.name2
}

//Surname
func (u *User) SetSurname(d string) {
	u.surname1 = d
}

func (u *User) GetSurname() string {
	return u.surname1
}

//Second surname
func (u *User) SetSurname2(d string) {
	u.surname2 = d
}

func (u *User) GetSurname2() string {
	return u.surname2
}

//Age
func (u *User) SetAge(a int) {
	u.age = a
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetAgeStr() string {
	return strconv.Itoa(u.age)
}

//Gender
func (u *User) SetGender(a int) {
	u.gender = a
}

func (u *User) GetGender() int {
	return u.gender
}

func (u *User) GetGenderStr() string {
	return GenderMap[u.gender]
}

//Fisical condition
func (u *User) SetFisCondition(d int) {
	u.fiscondition = d
}

func (u *User) GetFisCondition() int {
	return u.fiscondition
}

func (u *User) GetFisConditionStr() string {
	return FisconditionMap[u.fiscondition]
}

//Visual condition
func (u *User) SetVisualCondition(d int) {
	u.viscondition = d
}

func (u *User) GetVisualCondition() int {
	return u.viscondition
}

func (u *User) GetVisualConditionStr() string {
	return VisconditionMap[u.viscondition]
}

//Auditive condition
func (u *User) SetAuditiveCondition(d int) {
	u.audcondition = d
}

func (u *User) GetAuditiveCondition() int {
	return u.audcondition
}

func (u *User) GetAuditiveConditionStr() string {
	return AudconditionMap[u.audcondition]
}

//Prevision
func (u *User) SetPrevision(d int) {
	u.prevision = d
}

func (u *User) GetPrevision() int {
	return u.prevision
}

func (u *User) GetPrevisionStr() string {
	return PrevisionMap[u.prevision]
}

//Contact
func (u *User) SetContact(d string) {
	u.contact = d
}

func (u *User) GetContact() string {
	return u.contact
}

//Description
func (u *User) SetDescription(d string) {
	u.description = d
}

func (u *User) GetDescription() string {
	return u.description
}

func (u *User) GetFullName() string {
	return u.name + " " + u.name2 + " " + u.surname1 + " " + u.surname2
}

//Entry time
func (u *User) SetEntryDate(t time.Time) {
	u.entryDate = t
}

func (u *User) GetEntryDate() time.Time {
	return u.entryDate
}

func (u *User) GetEntryDateStr() string {
	return u.entryDate.Format("2006-01-02 15:04:05")
}
