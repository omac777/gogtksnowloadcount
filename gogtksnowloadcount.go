package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"	
)

type specialAssistant struct {
	// bool isLoadedReport;
	// std::string loadedReportFilename;
	v *gtk.Assistant

	//fields for page1
	page1frame *gtk.Frame
	page1Box *gtk.VBox
	openReportButton *gtk.Button
	filenameEntry *gtk.Entry
	b0, b1, b2, b3, b4, b5, b6 *gtk.HBox
	shiftStartTimeLabel *gtk.Label
	shiftStartTimeEntry *gtk.Entry
	shiftEndTimeLabel *gtk.Label
	shiftEndTimeEntry *gtk.Entry
	shiftStartDateLabel *gtk.Label
	shiftStartDateEntry *gtk.Entry
	guardNameLabel *gtk.Label
	guardNameEntry *gtk.Entry
	guardLicenseNumberLabel *gtk.Label
	guardLicenceNumberEntry *gtk.Entry
	guardShiftCommentsLabel *gtk.Label
	guardShiftCommentsEntry *gtk.Entry

	vbox *gtk.VBox
	comboboxentry *gtk.ComboBoxEntry
	combobox *gtk.ComboBox

	//fields for page2
	page2frame *gtk.Frame
	page2Box *gtk.VBox
	b7 *gtk.VBox
	conroyRadio, michaelRadio, strandherdRadio, innesRadio, clydeRadio *gtk.RadioButton

	vbox2 *gtk.VBox
	comboboxentry2 *gtk.ComboBoxEntry
	combobox2 *gtk.ComboBox


	//fields for page3
	page3frame *gtk.Frame
	page3Box *gtk.VBox
	b8 *gtk.VBox
	passesRadio, ticketsRadio *gtk.RadioButton
	
	//fields for page4
	page4frame *gtk.Frame
	page4Box *gtk.VBox
	b9 *gtk.VBox
	b10, b11, b12, b13, b14 *gtk.HBox
	singleaxlebutton *gtk.Button
	tandemaxlebutton *gtk.Button
	tripleaxlebutton *gtk.Button
	combotruckbutton *gtk.Button
	semitrailerbutton *gtk.Button
	singleLabel *gtk.Label
	tandemLabel *gtk.Label
	tripleLabel *gtk.Label
	comboLabel *gtk.Label
	semiLabel *gtk.Label
}

func (sa *specialAssistant) newPage1() {
	sa.page1Box = gtk.NewVBox(false, 1)
	sa.b0 = gtk.NewHBox(false, 1)
	sa.b1 = gtk.NewHBox(false, 1)
	sa.b2 = gtk.NewHBox(false, 1)
	sa.b3 = gtk.NewHBox(false, 1)
	sa.b4 = gtk.NewHBox(false, 1)
	sa.b5 = gtk.NewHBox(false, 1)
	sa.b6 = gtk.NewHBox(false, 1)
	sa.openReportButton = gtk.NewButtonWithLabel("Open report")
	sa.filenameEntry = gtk.NewEntry()
	sa.shiftStartTimeLabel = gtk.NewLabel("Shift Start Time:")
	sa.shiftStartTimeEntry = gtk.NewEntry()
	sa.shiftEndTimeLabel = gtk.NewLabel("Shift End Time:")
	sa.shiftEndTimeEntry = gtk.NewEntry()
	sa.shiftStartDateLabel = gtk.NewLabel("Shift Start Date:")
	sa.shiftStartDateEntry = gtk.NewEntry()
	sa.guardNameLabel = gtk.NewLabel("Guard Name:")
	sa.guardNameEntry = gtk.NewEntry()
	sa.guardLicenseNumberLabel = gtk.NewLabel("Guard License #:")
	sa.guardLicenceNumberEntry = gtk.NewEntry()
	sa.guardShiftCommentsLabel = gtk.NewLabel("Shift Comments:")
	sa.guardShiftCommentsEntry = gtk.NewEntry()

	sa.b0.PackStart(sa.openReportButton, true, true, 1);
	sa.b0.PackStart(sa.filenameEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b0, true, true, 1);

	sa.b1.PackStart(sa.shiftStartTimeLabel, true, true, 1);
	sa.b1.PackStart(sa.shiftStartTimeEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b1, true, true, 1);

	sa.b2.PackStart(sa.shiftEndTimeLabel, true, true, 1);
	sa.b2.PackStart(sa.shiftEndTimeEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b2, true, true, 1);
	
	sa.b3.PackStart(sa.shiftStartDateLabel, true, true, 1);
	sa.b3.PackStart(sa.shiftStartDateEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b3, true, true, 1);
	
	sa.b4.PackStart(sa.guardNameLabel, true, true, 1);
	sa.b4.PackStart(sa.guardNameEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b4, true, true, 1);
	
	sa.b5.PackStart(sa.guardLicenseNumberLabel, true, true, 1);
	sa.b5.PackStart(sa.guardLicenceNumberEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b5, true, true, 1);
	
	sa.b6.PackStart(sa.guardShiftCommentsLabel, true, true, 1);
	sa.b6.PackStart(sa.guardShiftCommentsEntry, true, true, 1);
	sa.page1Box.PackStart(sa.b6, true, true, 1);
	
	sa.page1frame = gtk.NewFrame("")
	sa.page1frame.Add(sa.page1Box)
	sa.v.AppendPage(sa.page1frame)
	sa.v.SetPageTitle(sa.page1frame, "Guard Shift")
	sa.v.SetPageType(sa.page1frame, gtk.ASSISTANT_PAGE_CONFIRM)
	sa.v.SetPageComplete(sa.page1frame, true)
}

func (sa *specialAssistant) newPage2() {
	sa.page2Box = gtk.NewVBox(false, 1)
	sa.b7 = gtk.NewVBox(false, 1)
	sa.conroyRadio = gtk.NewRadioButtonWithLabel(nil, "Conroy")
	sa.michaelRadio = gtk.NewRadioButtonWithLabel(sa.conroyRadio.GetGroup(), "Micheal")
	sa.strandherdRadio = gtk.NewRadioButtonWithLabel(sa.conroyRadio.GetGroup(), "Strandherd")
	sa.innesRadio = gtk.NewRadioButtonWithLabel(sa.conroyRadio.GetGroup(), "Innes")
	sa.clydeRadio = gtk.NewRadioButtonWithLabel(sa.conroyRadio.GetGroup(), "Clyde")

	sa.b7.PackStart(sa.conroyRadio, true, true, 1);
	sa.b7.PackStart(sa.michaelRadio, true, true, 1);
	sa.b7.PackStart(sa.strandherdRadio, true, true, 1);
	sa.b7.PackStart(sa.innesRadio, true, true, 1);
	sa.b7.PackStart(sa.clydeRadio, true, true, 1);
	sa.page2Box.PackStart(sa.b7, true, true, 1);
	
	sa.page2frame = gtk.NewFrame("" )
	sa.page2frame.Add(sa.page2Box)
	sa.v.AppendPage(sa.page2frame)
	sa.v.SetPageTitle(sa.page2frame, "Counting At Location:")
	sa.v.SetPageType(sa.page2frame, gtk.ASSISTANT_PAGE_CONFIRM)
	sa.v.SetPageComplete(sa.page2frame, true)
}

func (sa *specialAssistant) newPage3() {
	sa.page3Box = gtk.NewVBox(false, 1)
	sa.b8 = gtk.NewVBox(false, 1)

	sa.passesRadio = gtk.NewRadioButtonWithLabel(nil, "Passes")
	sa.ticketsRadio = gtk.NewRadioButtonWithLabel(sa.passesRadio.GetGroup(), "Tickets")

	sa.b8.PackStart(sa.passesRadio, true, true, 1);
	sa.b8.PackStart(sa.ticketsRadio, true, true, 1);
	sa.page3Box.PackStart(sa.b8, true, true, 1);

	sa.page3frame = gtk.NewFrame("" )
	sa.page3frame.Add(sa.page3Box)
	sa.v.AppendPage(sa.page3frame)
	sa.v.SetPageTitle(sa.page3frame, "Counting Totals For:")
	sa.v.SetPageType(sa.page3frame, gtk.ASSISTANT_PAGE_CONFIRM)
	sa.v.SetPageComplete(sa.page3frame, true)
}

func (sa *specialAssistant) newPage4() {
	sa.page4Box = gtk.NewVBox(false, 1)
	sa.b9 = gtk.NewVBox(false, 1)

	sa.b10 = gtk.NewHBox(false, 1)
	sa.b11 = gtk.NewHBox(false, 1)
	sa.b12 = gtk.NewHBox(false, 1)
	sa.b13 = gtk.NewHBox(false, 1)
	sa.b14 = gtk.NewHBox(false, 1)

	sa.singleaxlebutton = gtk.NewButtonWithLabel("singleaxle")
	sa.tandemaxlebutton = gtk.NewButtonWithLabel("tandemaxle")
	sa.tripleaxlebutton = gtk.NewButtonWithLabel("tripleaxle")
	sa.combotruckbutton = gtk.NewButtonWithLabel("combotruck")
	sa.semitrailerbutton = gtk.NewButtonWithLabel("semitrailer")

	sa.singleLabel = gtk.NewLabel("0")
	sa.tandemLabel = gtk.NewLabel("0")
	sa.tripleLabel = gtk.NewLabel("0")
	sa.comboLabel = gtk.NewLabel("0")
	sa.semiLabel = gtk.NewLabel("0")

	sa.b10.PackStart(sa.singleaxlebutton, true, true, 1);
	sa.b10.PackStart(sa.singleLabel, true, true, 1);
	sa.b9.PackStart(sa.b10, true, true, 1);
	sa.b11.PackStart(sa.tandemaxlebutton, true, true, 1);
	sa.b11.PackStart(sa.tandemLabel, true, true, 1);
	sa.b9.PackStart(sa.b11, true, true, 1);
	sa.b12.PackStart(sa.tripleaxlebutton, true, true, 1);
	sa.b12.PackStart(sa.tripleLabel, true, true, 1);
	sa.b9.PackStart(sa.b12, true, true, 1);
	sa.b13.PackStart(sa.combotruckbutton, true, true, 1);
	sa.b13.PackStart(sa.comboLabel, true, true, 1);
	sa.b9.PackStart(sa.b13, true, true, 1);
	sa.b14.PackStart(sa.semitrailerbutton, true, true, 1);
	sa.b14.PackStart(sa.semiLabel, true, true, 1);
	sa.b9.PackStart(sa.b14, true, true, 1);
	sa.page4Box.PackStart(sa.b9, true, true, 1);

	sa.page4frame = gtk.NewFrame("" )
	sa.page4frame.Add(sa.page4Box)
	sa.v.AppendPage(sa.page4frame)
	sa.v.SetPageTitle(sa.page4frame, "Truck Type Count")
	sa.v.SetPageType(sa.page4frame, gtk.ASSISTANT_PAGE_CONFIRM)
	sa.v.SetPageComplete(sa.page4frame, true)
}

func main() {
	gtk.Init(&os.Args)
	myspecialAssistant := specialAssistant{}
	myspecialAssistant.v = gtk.NewAssistant()
	myspecialAssistant.newPage4()
	myspecialAssistant.newPage3()
        myspecialAssistant.newPage2()
	myspecialAssistant.newPage1()
	myspecialAssistant.v.Connect("cancel", gtk.MainQuit)
	myspecialAssistant.v.Connect("close", gtk.MainQuit)
	myspecialAssistant.v.SetSizeRequest(640, 480)
	myspecialAssistant.v.ShowAll()
	gtk.Main()
}
