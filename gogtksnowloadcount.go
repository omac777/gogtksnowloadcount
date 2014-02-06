package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"	
	"time"
	"fmt"
	"bufio"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
	"bytes"
)

const SHIFTSTARTTIME = string("shiftstarttime")
const SHIFTENDTIME = string("shiftendtime")
const SHIFTSTARTDATE = string("shiftstartdate")
const GUARDNAME = string("guardname")
const LICENSENUMBER = string("licensenumber")
const SHIFTCOMMENT = string("shiftcomment")
const COUNTLOCATION = string("countlocation")
const COUNTFORITEMTYPE = string("countforitemtype")
const SINGLEAXLE = string("singleaxle")
const TANDEMAXLE = string("tandem2axle")
const TRIPLEAXLE = string("tripleaxle")
const COMBOTRUCK = string("combotruck")
const SEMITRAILER = string("semitrailer")

//json stuff has to have fields with names starting with uppercase letters
//this is a golang special wierdness
//to be honest, I'm not happy about this
//We should have complete control over our variable names
type SNLDB struct {
	SNLMap map[string]string `json:"SNLMap"`
	SingleL []time.Time `json:"SingleL"` 
	TandemL []time.Time `json:"TandemL"` 
	TripleL []time.Time `json:"TripleL"` 
	ComboL []time.Time `json:"ComboL"` 
	SemiL []time.Time `json:"SemiL"`
}

func NewSNLDB() *SNLDB {
	f := SNLDB{}
	f.SNLMap = make(map[string]string)
	f.SingleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	f.TandemL = make([]time.Time, 0, 200)
	f.TripleL = make([]time.Time, 0, 200)
	f.ComboL = make([]time.Time, 0, 200)
	f.SemiL = make([]time.Time, 0, 200)
	return &f
}

func (s *SNLDB) testSetAndGetDataFields() {
	s.setShiftStartTime("b");
	s.setShiftEndTime("c");
	s.setShiftStartDate("d");
	s.setGuardName("e");
	s.setLicenseNumber("f");
	s.setShiftComment("g");
	s.setCountLocation("h");
	s.setCountForItemType("i");
	s.singleAxleArrived();
	s.tandemAxleArrived();
	s.tandemAxleArrived();
	s.tripleAxleArrived();
	s.tripleAxleArrived();
	s.tripleAxleArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	fmt.Printf("%v\n", s.getShiftStartTime())
	fmt.Printf("%v\n", s.getShiftEndTime())
	fmt.Printf("%v\n", s.getShiftStartDate())
	fmt.Printf("%v\n", s.getGuardName())
	fmt.Printf("%v\n", s.getLicenseNumber())
	fmt.Printf("%v\n", s.getShiftComment())
	fmt.Printf("%v\n", s.getCountLocation())
	fmt.Printf("%v\n", s.getCountForItemType())
	fmt.Printf("single total: %v\n", s.getSingleAxleTotal())
	fmt.Printf("tandem total: %v\n", s.getTandemAxleTotal())
	fmt.Printf("triple total: %v\n", s.getTripleAxleTotal())
	fmt.Printf("combo total: %v\n" , s.getComboTruckTotal())
	fmt.Printf("semi total: %v\n", s.getSemiTrailerTotal())
}

func (s *SNLDB) debugDataFields() {
	fmt.Printf("%v\n", s.getShiftStartTime())
	fmt.Printf("%v\n", s.getShiftEndTime())
	fmt.Printf("%v\n", s.getShiftStartDate())
	fmt.Printf("%v\n", s.getGuardName())
	fmt.Printf("%v\n", s.getLicenseNumber())
	fmt.Printf("%v\n", s.getShiftComment())
	fmt.Printf("%v\n", s.getCountLocation())
	fmt.Printf("%v\n", s.getCountForItemType())
	fmt.Printf("single total: %v\n", s.getSingleAxleTotal())
	fmt.Printf("tandem total: %v\n", s.getTandemAxleTotal())
	fmt.Printf("triple total: %v\n", s.getTripleAxleTotal())
	fmt.Printf("combo total: %v\n" , s.getComboTruckTotal())
	fmt.Printf("semi total: %v\n", s.getSemiTrailerTotal())
}

func (s *SNLDB) getShiftStartTime() string {
	return string(s.SNLMap[SHIFTSTARTTIME])
}

func (s *SNLDB) getShiftEndTime() string {
	return string(s.SNLMap[SHIFTENDTIME])
}

func (s *SNLDB) getShiftStartDate() string {
	return string(s.SNLMap[SHIFTSTARTDATE])
}

func (s *SNLDB) getGuardName() string {
	return string(s.SNLMap[GUARDNAME])
}

func (s *SNLDB) getLicenseNumber() string {
	return string(s.SNLMap[LICENSENUMBER])
}

func (s *SNLDB) getShiftComment() string {
	return string(s.SNLMap[SHIFTCOMMENT])
}

func (s *SNLDB) getCountLocation() string {
	return string(s.SNLMap[COUNTLOCATION])
}

func (s *SNLDB) getCountForItemType() string {
	return string(s.SNLMap[COUNTFORITEMTYPE])
}

func (s *SNLDB) setShiftStartTime(s_ string) {
	s.SNLMap[SHIFTSTARTTIME] = s_
}

func (s *SNLDB) setShiftEndTime(s_ string) {
	s.SNLMap[SHIFTENDTIME] = s_
}

func (s *SNLDB) setShiftStartDate(s_ string) {
	s.SNLMap[SHIFTSTARTDATE] = s_
}

func (s *SNLDB) setGuardName(s_ string) {
	s.SNLMap[GUARDNAME] = s_
}

func (s *SNLDB) setLicenseNumber(s_ string) {
	s.SNLMap[LICENSENUMBER] = s_
}

func (s *SNLDB) setShiftComment(s_ string) {
	s.SNLMap[SHIFTCOMMENT] = s_
}

func (s *SNLDB) setCountLocation(s_ string) {
	s.SNLMap[COUNTLOCATION] = s_
}

func (s *SNLDB) setCountForItemType(s_ string) {
	s.SNLMap[COUNTFORITEMTYPE] = s_
}

func (s *SNLDB) getSingleAxleTotal() int {
	return len(s.SingleL)
}

func (s *SNLDB) getTandemAxleTotal() int {
	return len(s.TandemL)
}

func (s *SNLDB) getTripleAxleTotal() int {
	return len(s.TripleL)
}

func (s *SNLDB) getComboTruckTotal() int {
	return len(s.ComboL)
}

func (s *SNLDB) getSemiTrailerTotal() int {
	return len(s.SemiL)
}

func (s *SNLDB) singleAxleArrived() {
	s.SingleL = append(s.SingleL, time.Now())
}

func (s *SNLDB) tandemAxleArrived() {
	s.TandemL = append(s.TandemL, time.Now())
}

func (s *SNLDB) tripleAxleArrived() {
	s.TripleL = append(s.TripleL, time.Now())	
}

func (s *SNLDB) comboTruckArrived() {
	s.ComboL = append(s.ComboL, time.Now())	
}

func (s *SNLDB) semiTrailerArrived() {
	s.SemiL = append(s.SemiL, time.Now())	
}

func (s *SNLDB) clear() {
	s.SNLMap = make(map[string]string)
	s.SingleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	s.TandemL = make([]time.Time, 0, 200)
	s.TripleL = make([]time.Time, 0, 200)
	s.ComboL = make([]time.Time, 0, 200)
	s.SemiL = make([]time.Time, 0, 200)
}

type specialAssistant struct {
	// bool isLoadedReport;
	// std::string loadedReportFilename;
	v *gtk.Assistant
	_snldb *SNLDB

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
	SingleLabel *gtk.Label
	TandemLabel *gtk.Label
	TripleLabel *gtk.Label
	ComboLabel *gtk.Label
	SemiLabel *gtk.Label
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
	//tmpblah := sa.readJsonFileSNLDB("snowreport.json")
	//fmt.Println("after returning from readJsonFile %v\n", tmpblah)

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
	sa.singleaxlebutton.Clicked(func() {
		sa._snldb.singleAxleArrived()
		sa.SingleLabel.SetText(strconv.Itoa(sa._snldb.getSingleAxleTotal()));
	})	
	sa.tandemaxlebutton.Clicked(func() {
		sa._snldb.tandemAxleArrived()
		sa.TandemLabel.SetText(strconv.Itoa(sa._snldb.getTandemAxleTotal()));
	})	
	sa.tripleaxlebutton.Clicked(func() {
		sa._snldb.tripleAxleArrived()
		sa.TripleLabel.SetText(strconv.Itoa(sa._snldb.getTripleAxleTotal()));
	})	
	sa.combotruckbutton.Clicked(func() {
		sa._snldb.comboTruckArrived()
		sa.ComboLabel.SetText(strconv.Itoa(sa._snldb.getComboTruckTotal()));
	})	
	sa.semitrailerbutton.Clicked(func() {
		sa._snldb.semiTrailerArrived()
		sa.SemiLabel.SetText(strconv.Itoa(sa._snldb.getSemiTrailerTotal()));
	})	
	
	sa.SingleLabel = gtk.NewLabel("0")
	sa.TandemLabel = gtk.NewLabel("0")
	sa.TripleLabel = gtk.NewLabel("0")
	sa.ComboLabel = gtk.NewLabel("0")
	sa.SemiLabel = gtk.NewLabel("0")

	sa.b10.PackStart(sa.singleaxlebutton, true, true, 1);
	sa.b10.PackStart(sa.SingleLabel, true, true, 1);
	sa.b9.PackStart(sa.b10, true, true, 1);
	sa.b11.PackStart(sa.tandemaxlebutton, true, true, 1);
	sa.b11.PackStart(sa.TandemLabel, true, true, 1);
	sa.b9.PackStart(sa.b11, true, true, 1);
	sa.b12.PackStart(sa.tripleaxlebutton, true, true, 1);
	sa.b12.PackStart(sa.TripleLabel, true, true, 1);
	sa.b9.PackStart(sa.b12, true, true, 1);
	sa.b13.PackStart(sa.combotruckbutton, true, true, 1);
	sa.b13.PackStart(sa.ComboLabel, true, true, 1);
	sa.b9.PackStart(sa.b13, true, true, 1);
	sa.b14.PackStart(sa.semitrailerbutton, true, true, 1);
	sa.b14.PackStart(sa.SemiLabel, true, true, 1);
	sa.b9.PackStart(sa.b14, true, true, 1);
	sa.page4Box.PackStart(sa.b9, true, true, 1);

	sa.page4frame = gtk.NewFrame("" )
	sa.page4frame.Add(sa.page4Box)
	sa.v.AppendPage(sa.page4frame)
	sa.v.SetPageTitle(sa.page4frame, "Truck Type Count")
	sa.v.SetPageType(sa.page4frame, gtk.ASSISTANT_PAGE_CONFIRM)
	sa.v.SetPageComplete(sa.page4frame, true)
}

func (sa *specialAssistant) saveJsonFileSNLDB(myFileName string, f *SNLDB) () {
	fo, err := os.Create(myFileName)
	if err != nil { panic(err) }
	defer fo.Close()
	w := bufio.NewWriter(fo)
 	var myjsonBytes bytes.Buffer
 	enc := json.NewEncoder(&myjsonBytes)
 	err = enc.Encode(f)
	//fmt.Println("myjsonbytes %v\n", myjsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(myjsonBytes.Bytes()) 
	if err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil { 
		log.Fatal(err) 
	}
}

func (sa *specialAssistant) readJsonFileSNLDB(myFileName string) (SNLDB) {
	file, err2 := ioutil.ReadFile(myFileName)
	if err2 != nil {
	   	log.Fatal(err2)
	}
	jsonBytes := bytes.NewBuffer(file)
 	dec := json.NewDecoder(jsonBytes)
 	var res SNLDB
 	err3 := dec.Decode(&res)
	if err3 != nil {
	   	log.Fatal(err3)
	}
 	//fmt.Printf("Read Results: %+v %v\n", res, err3)
	return res
}

func (sa *specialAssistant) apply_clicked () {
	println("assistant apply clicked page:", sa.v.GetCurrentPage())
	var whichPage int
	whichPage = sa.v.GetCurrentPage()
	switch (whichPage) {
	case 0:
		sa._snldb.setShiftStartTime(sa.shiftStartTimeEntry.GetText());
		sa._snldb.setShiftEndTime(sa.shiftEndTimeEntry.GetText());
		sa._snldb.setShiftStartDate(sa.shiftStartDateEntry.GetText());
		sa._snldb.setGuardName(sa.guardNameEntry.GetText());
		sa._snldb.setLicenseNumber(sa.guardLicenceNumberEntry.GetText());
		sa._snldb.setShiftComment(sa.guardShiftCommentsEntry.GetText());
		break;
	case 1:
		sa._snldb.setCountLocation(sa.getCountLocation());
		break;
	case 2:
		sa._snldb.setCountForItemType(sa.getCountForItemType());
		break;
	case 3:
		break;
	default:
		break;
	}	
}

func (sa *specialAssistant) getCountLocation() string {
	var selectedRadio string
	if(sa.conroyRadio.GetActive()) {
		selectedRadio = sa.conroyRadio.GetLabel()
	}

	if(sa.michaelRadio.GetActive()) {
		selectedRadio = sa.michaelRadio.GetLabel()
	}

	if(sa.strandherdRadio.GetActive()) {
		selectedRadio = sa.strandherdRadio.GetLabel()
	}

	if(sa.innesRadio.GetActive()) {
		selectedRadio = sa.innesRadio.GetLabel()
	}

	if(sa.clydeRadio.GetActive()) {
		selectedRadio = sa.clydeRadio.GetLabel()
	}
	return selectedRadio
}

func (sa *specialAssistant) getCountForItemType() string {
	var selectedRadio string  
	if(sa.passesRadio.GetActive()) {
		selectedRadio = sa.passesRadio.GetLabel()
	}

	if(sa.ticketsRadio.GetActive()) {
		selectedRadio = sa.ticketsRadio.GetLabel()
	}

	return selectedRadio
}

func (sa *specialAssistant) generateUniqueReportFilename() string {
	var myUniqueFilename string
	var t time.Time
	var myIsoTime string
	t = time.Now()
	myIsoTime = fmt.Sprintf("%d%02d%02dT%02d%02d%02d.%09d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	myUniqueFilename = "snowloadcount" + myIsoTime + ".json";
	return myUniqueFilename
}

func (sa *specialAssistant) close_clicked () {
	println("assistant close clicked page:", sa.v.GetCurrentPage())
	sa.saveJsonFileSNLDB(sa.generateUniqueReportFilename(), sa._snldb)
	gtk.MainQuit()
}

func (sa *specialAssistant) cancel_clicked () {
	println("assistant cancel clicked page:", sa.v.GetCurrentPage())
}

func (sa *specialAssistant) prepare_clicked () {
	println("assistant prepare clicked page:", sa.v.GetCurrentPage())
}

func main() {
	gtk.Init(&os.Args)
	myspecialAssistant := specialAssistant{}
	myspecialAssistant._snldb = NewSNLDB()
	myspecialAssistant._snldb.testSetAndGetDataFields()
	myspecialAssistant._snldb.clear()
	myspecialAssistant.v = gtk.NewAssistant()
	myspecialAssistant.newPage4()
	myspecialAssistant.newPage3()
        myspecialAssistant.newPage2()
	myspecialAssistant.newPage1()

	myspecialAssistant.v.Connect("apply", myspecialAssistant.apply_clicked)
	myspecialAssistant.v.Connect("cancel", myspecialAssistant.cancel_clicked)
	myspecialAssistant.v.Connect("close", myspecialAssistant.close_clicked)
	myspecialAssistant.v.Connect("prepare", myspecialAssistant.prepare_clicked)

	myspecialAssistant.v.SetSizeRequest(640, 480)
	myspecialAssistant.v.ShowAll()
	gtk.Main()
}
