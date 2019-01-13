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
	13: "NO APLICA",
}

var VisconditionMap = map[int]string{
	0: "CONGÉNITO",
	1: "ADQUIRIDO",
	2: "OTROS",
	3: "NO APLICA",
}

var AudconditionMap = map[int]string{
	0: "CONGÉNITO",
	1: "ADQUIRIDO",
	2: "OTROS",
	3: "NO APLICA",
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

//User type
type User struct {
	Rut          string `json:"rut"`
	Name         string `json:"name"`
	Name2        string `json:"name2"`
	Surname1     string `json:"surname"`
	Surname2     string `json:"surname2"`
	Age          int
	Gender       int
	Fiscondition int
	Viscondition int
	Audcondition int
	Prevition    int
	Contact      string
	Description  string
	EntryDate    time.Time
}

func CreateUser(n1, n2, s1, s2 string) *User {
	nu := new(User)
	nu.Name = n1
	nu.Name2 = n2
	nu.Surname1 = s1
	nu.Surname2 = s2
	nu.EntryDate = time.Now()
	return nu
}

//Rut
func (u *User) SetRut(d string) {
	u.Rut = d
}

func (u *User) GetRut() string {
	return u.Rut
}

//Name
func (u *User) SetName(d string) {
	u.Name = d
}

func (u *User) GetName() string {
	return u.Name
}

//Second Name
func (u *User) SetName2(d string) {
	u.Name2 = d
}

func (u *User) GetName2() string {
	return u.Name2
}

//SurName
func (u *User) SetSurName(d string) {
	u.Surname1 = d
}

func (u *User) GetSurName() string {
	return u.Surname1
}

//Second Surname
func (u *User) SetSurName2(d string) {
	u.Surname2 = d
}

func (u *User) GetSurName2() string {
	return u.Surname2
}

//Age
func (u *User) SetAge(a int) {
	u.Age = a
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) GetAgeStr() string {
	return strconv.Itoa(u.Age)
}

//Gender
func (u *User) SetGender(a int) {
	u.Gender = a
}

func (u *User) GetGender() int {
	return u.Gender
}

func (u *User) GetGenderStr() string {
	return GenderMap[u.Gender]
}

//Fisical condition
func (u *User) SetFisCondition(d int) {
	u.Fiscondition = d
}

func (u *User) GetFisCondition() int {
	return u.Fiscondition
}

func (u *User) GetFisConditionStr() string {
	return FisconditionMap[u.Fiscondition]
}

//Visual condition
func (u *User) SetVisualCondition(d int) {
	u.Viscondition = d
}

func (u *User) GetVisualCondition() int {
	return u.Viscondition
}

func (u *User) GetVisualConditionStr() string {
	return VisconditionMap[u.Viscondition]
}

//Auditive condition
func (u *User) SetAuditiveCondition(d int) {
	u.Audcondition = d
}

func (u *User) GetAuditiveCondition() int {
	return u.Audcondition
}

func (u *User) GetAuditiveConditionStr() string {
	return AudconditionMap[u.Audcondition]
}

//Prevision
func (u *User) SetPrevision(d int) {
	u.Prevition = d
}

func (u *User) GetPrevision() int {
	return u.Prevition
}

func (u *User) GetPrevisionStr() string {
	return PrevisionMap[u.Prevition]
}

//Contact
func (u *User) SetContact(d string) {
	u.Contact = d
}

func (u *User) GetContact() string {
	return u.Contact
}

//Description
func (u *User) SetDescription(d string) {
	u.Description = d
}

func (u *User) GetDescription() string {
	return u.Description
}

func (u *User) GetFullName() string {
	return u.Name + " " + u.Name2 + " " + u.Surname1 + " " + u.Surname2
}

//Entry time
func (u *User) SetEntryDate(t time.Time) {
	u.EntryDate = t
}

func (u *User) GetEntryDate() time.Time {
	return u.EntryDate
}

func (u *User) GetEntryDateStr() string {
	return u.EntryDate.Format("2006-01-02 15:04:05")
}
