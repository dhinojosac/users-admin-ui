package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dhinojosac/users-admin-ui/ush"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/boltdb/bolt"
)

var mainwin *ui.Window
var statusLbl *ui.Label

var uirut *ui.Entry
var uiname *ui.Entry
var uiname2 *ui.Entry
var uisn1 *ui.Entry
var uisn2 *ui.Entry
var uiage *ui.Entry
var uigender *ui.Combobox
var uifiscon *ui.Combobox
var uiviscon *ui.Combobox
var uiaudcon *ui.Combobox
var uiprev *ui.Combobox
var uicont *ui.Entry
var uied *ui.DateTimePicker
var uidesc *ui.MultilineEntry

var db *bolt.DB

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("USERS"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func showNewWindows() {
	errwin := ui.NewWindow("Alerta", 320, 160, true)
	vbox := ui.NewVerticalBox()
	acceptBtn := ui.NewButton("Aceptar")
	errwin.OnClosing(func(*ui.Window) bool {
		errwin.Hide()
		return true
	})
	errwin.SetChild(vbox)
	vbox.Append(ui.NewLabel("Realmente desea cerrar la aplicación?"), true)
	vbox.Append(acceptBtn, true)
	errwin.Show()
}

func makeBasicControlsPage() ui.Control {
	fullBox := ui.NewVerticalBox()
	fullBox.SetPadded(true)

	namesBox := ui.NewHorizontalBox()
	namesBox.SetPadded(true)
	fullBox.Append(namesBox, false)

	uiname = ui.NewEntry()
	namesBox.Append(ui.NewLabel(" Nombres * "), false)
	namesBox.Append(uiname, false)
	uiname2 = ui.NewEntry()
	namesBox.Append(uiname2, false)

	surnameBox := ui.NewHorizontalBox()
	surnameBox.SetPadded(true)
	fullBox.Append(surnameBox, false)

	uisn1 = ui.NewEntry()
	surnameBox.Append(ui.NewLabel(" Apellidos * "), false)
	surnameBox.Append(uisn1, false)
	uisn2 = ui.NewEntry()
	surnameBox.Append(uisn2, false)

	diagAgeBox := ui.NewHorizontalBox()
	diagAgeBox.SetPadded(true)
	fullBox.Append(diagAgeBox, false)

	diagAgeBox.Append(ui.NewLabel(" RUT "), false)
	uirut := ui.NewEntry()
	diagAgeBox.Append(uirut, false)

	diagAgeBox.Append(ui.NewLabel("   Sexo "), false)
	uisex := ui.NewCombobox()
	uisex.Append("Hombre")
	uisex.Append("Mujer")
	diagAgeBox.Append(uisex, false)
	diagAgeBox.Append(ui.NewLabel("  Edad"), false)
	uiage := ui.NewEntry()
	diagAgeBox.Append(uiage, false)

	contactBox := ui.NewHorizontalBox()
	contactBox.SetPadded(true)
	fullBox.Append(contactBox, false)
	contactBox.Append(ui.NewLabel("Contacto"), false)
	uicont = ui.NewEntry()
	contactBox.Append(uicont, false)
	contactBox.Append(ui.NewLabel("Previsión"), false)
	uiprev = ui.NewCombobox()
	//Fill prevision
	for i := 0; i < len(ush.PrevisionMap); i++ {
		uiprev.Append(ush.PrevisionMap[i])
	}
	contactBox.Append(uiprev, false)

	//Fill combobox with fisical condition
	uifiscon = ui.NewCombobox()
	for i := 0; i < len(ush.FisconditionMap); i++ {
		uifiscon.Append(ush.FisconditionMap[i])
	}

	//Fill combobox with visual condition
	uiviscon = ui.NewCombobox()
	for i := 0; i < len(ush.VisconditionMap); i++ {
		uiviscon.Append(ush.VisconditionMap[i])
	}

	//Fill combobox with auditive condition
	uiaudcon = ui.NewCombobox()
	for i := 0; i < len(ush.AudconditionMap); i++ {
		uiaudcon.Append(ush.AudconditionMap[i])
	}

	fullBox.Append(ui.NewLabel("Condición física  "), false)
	fullBox.Append(uifiscon, false)
	fullBox.Append(ui.NewLabel("Condición sensorial visual  "), false)
	fullBox.Append(uiviscon, false)
	fullBox.Append(ui.NewLabel("Condición sensorial auditivo  "), false)
	fullBox.Append(uiaudcon, false)

	fullBox.Append(ui.NewLabel("Descripción del usuario"), false)
	uidesc = ui.NewMultilineEntry()
	fullBox.Append(uidesc, true)

	saveBtn := ui.NewButton("GUARDAR")
	saveBtn.OnClicked(func(*ui.Button) {
		fmt.Println("Save button pressed!")
		if len(uiname.Text()) == 0 || len(uisn1.Text()) == 0 {
			ui.MsgBoxError(mainwin, "Usuario no válido", "Debes agregar Nombre y Apellido del usuario.")
			/*
				//test new windows
				showNewWindows()
			*/
		} else {
			user := ush.CreateUser(uiname.Text(), uiname2.Text(), uisn1.Text(), uisn2.Text())
			user.SetDescription(uidesc.Text())
			user.SetFisCondition(uifiscon.Selected())
			ageint, err := strconv.Atoi(uiage.Text())
			if err != nil {

			}
			user.SetAge(ageint)

			err = addUserDB(db, user.GetRut(), user.GetName(), user.GetSurname(), time.Now())
			if err != nil {
				log.Fatal(err)
			}

			//To check data user
			fmt.Println("Name: " + user.GetFullName())
			fmt.Printf("Edad:%d\n", user.GetAge())
			fmt.Println("Edad:" + user.GetAgeStr())
			//fmt.Println("Diag: " + user.GetFisConditionStr())
			fmt.Println("Desc: " + user.GetDescription())
		}
		//fmt.Println("Fecha de ingreso: " + uied.Time().Format("2006-01-02 15:04:05"))
	})
	fullBox.Append(saveBtn, false)
	statusLbl = ui.NewLabel("[INFO] Programa en dasarrollo. dhinojosac 2019")
	fullBox.Append(statusLbl, false)
	return fullBox
}

func makeUserList() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	return hbox
}

func addUserDB(db *bolt.DB, r, n, sn string, date time.Time) error {
	entry := ush.User{rut: r, name: n, surname: sn}
	entryBytes, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not marshal entry json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("USERS")).Put([]byte(date.Format(time.RFC3339)), entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})
	fmt.Println("Added Entry")
	return err
}

func setupUI() {

	//Check database
	var err error
	db, err = setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Create main windows
	mainwin = ui.NewWindow("Registro de usuarios v0.1", 640, 500, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Agregar usuario", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("Lista de usuarios", makeUserList())
	tab.SetMargined(1, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
