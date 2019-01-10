package main

import (
	"fmt"
	"strconv"

	"github.com/dhinojosac/users-admin-ui/ush"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window
var statusLbl *ui.Label

var uiname *ui.Entry
var uiname2 *ui.Entry
var uisn1 *ui.Entry
var uisn2 *ui.Entry
var uifiscon *ui.Combobox
var uiviscon *ui.Combobox
var uiaudcon *ui.Combobox
var uiage *ui.Entry

//var uibd *ui.DateTimePicker
var uied *ui.DateTimePicker
var uidesc *ui.MultilineEntry

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

	nameBox := ui.NewHorizontalBox()
	nameBox.SetPadded(false)
	fullBox.Append(nameBox, false)

	group := ui.NewGroup("Datos Usuario")
	group.SetMargined(true)
	nameBox.Append(group, true)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	uiname = ui.NewEntry()
	entryForm.Append("Nombre (*)", uiname, false)
	uiname2 = ui.NewEntry()
	entryForm.Append("Nombre 2", uiname2, false)
	uisn1 = ui.NewEntry()
	entryForm.Append("Apellido (*)", uisn1, false)
	uisn2 = ui.NewEntry()
	entryForm.Append("Apellido 2", uisn2, false)

	diagAgeBox := ui.NewHorizontalBox()
	diagAgeBox.SetPadded(false)
	fullBox.Append(diagAgeBox, false)
	diagAgeBox.Append(ui.NewLabel("Edad:  "), false)
	uiage := ui.NewEntry()
	diagAgeBox.Append(uiage, false)
	diagAgeBox.Append(ui.NewLabel("    "), false)
	diagAgeBox.Append(ui.NewLabel(" Fecha de ingreso: "), false)
	uied = ui.NewDateTimePicker()
	diagAgeBox.Append(uied, false)

	uifiscon = ui.NewCombobox()
	for _, value := range ush.FisconditionMap {
		uifiscon.Append(value)
	}
	uiviscon = ui.NewCombobox()
	for _, value := range ush.VisconditionMap {
		uiviscon.Append(value)
	}
	uiaudcon = ui.NewCombobox()
	for _, value := range ush.AudconditionMap {
		uiaudcon.Append(value)
	}

	fullBox.Append(ui.NewLabel("Condición física:  "), true)
	fullBox.Append(uifiscon, false)
	fullBox.Append(ui.NewLabel("Condición sensorial visual:  "), true)
	fullBox.Append(uiviscon, false)
	fullBox.Append(ui.NewLabel("Condición sensorial auditivo:  "), true)
	fullBox.Append(uiaudcon, false)

	fullBox.Append(ui.NewLabel("Descripción del usuario:"), false)
	uidesc = ui.NewMultilineEntry()
	fullBox.Append(uidesc, false)

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
			user := ush.CreateUser(uiname.Text(), uisn1.Text(), uisn2.Text())
			user.SetDescription(uidesc.Text())
			user.SetFisCondition(uifiscon.Selected())
			ageint, err := strconv.Atoi(uiage.Text())
			if err != nil {

			}
			user.SetAge(ageint)

			//To check data user
			fmt.Println("Name: " + user.GetFullName())
			fmt.Printf("Edad:%d\n", user.GetAge())
			fmt.Println("Edad:" + user.GetAgeStr())
			fmt.Println("Diag: " + user.GetFisConditionStr())
			fmt.Println("Desc: " + user.GetDescription())
		}
		//fmt.Println("Fecha de ingreso: " + uied.Time().Format("2006-01-02 15:04:05"))
	})
	fullBox.Append(saveBtn, false)
	statusLbl = ui.NewLabel("[INFO] Programa desarrollado por dhinojosac 2019 para la Chanchita <3")
	fullBox.Append(statusLbl, false)
	return fullBox
}

func makeUserList() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	return hbox
}

func setupUI() {
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
