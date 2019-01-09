package main

import (
	"fmt"

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
var uibd *ui.DateTimePicker
var uied *ui.DateTimePicker
var uidesc *ui.Entry

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

	fullBox.Append(ui.NewLabel("Fecha de nacimiento:"), false)
	uibd = ui.NewDatePicker()
	fullBox.Append(uibd, false)

	fullBox.Append(ui.NewLabel("Fecha de ingreso:"), false)
	uied = ui.NewDateTimePicker()
	fullBox.Append(uied, false)

	fullBox.Append(ui.NewLabel("Descripción del usuario:"), false)
	uidesc = ui.NewEntry()
	fullBox.Append(uidesc, true)

	saveBtn := ui.NewButton("GUARDAR")
	saveBtn.OnClicked(func(*ui.Button) {
		fmt.Println("Save button pressed!")
		if len(uiname.Text()) == 0 || len(uisn1.Text()) == 0 {
			//ui.MsgBoxError(mainwin,"Usuario no válido", "Debes agregar Nombre y Apellido del usuario.")
			errwin := ui.NewWindow("Error Windows", 640, 480, true)
			errwin.Show()
		} else {
			user := ush.CreateUser(uiname.Text(), uisn1.Text(), uisn2.Text())
			user.SetDescription(uidesc.Text())

			fmt.Println("Name: " + user.GetFullName())
		}
		//fmt.Println("Fecha de ingreso: " + uied.Time().Format("2006-01-02 15:04:05"))
	})
	fullBox.Append(saveBtn, false)
	return fullBox
}

func makeUserList() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	return hbox
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("Registro de usuarios v0.1", 640, 480, true)
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

	//tab.Append("Numbers and Lists", makeNumbersPage())
	//tab.SetMargined(1, true)

	//tab.Append("Data Choosers", makeDataChoosersPage())
	//tab.SetMargined(2, true)
	//statusLbl := ui.NewLabel("[INFO] Programa desarrollado por DOHC 2019")
	//mainwin.Append(statusLbl, false)
	mainwin.Show()
}

func main() {
	/*
		user := ush.CreateUser("Diego", "Hinojosa", "Córdova")
		desc := "Diagnosticado con protruciones lumbares en L3-L4 y L5-S1, sin compromiso nervioso."
		user.SetDescription(desc)
		user.SetAge(30)

		uname := user.GetFullName()
		fmt.Println("Nombre: " + uname)
		fmt.Println("Edad: " + strconv.Itoa(user.GetAge()))
		fmt.Println("Descripción: " + user.GetDescription())
		fmt.Println("Fecha ingreso: " + user.GetAgeStr())
	*/
	ui.Main(setupUI)
}
