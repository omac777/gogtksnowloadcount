package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"	
	"time"
	//"container/list"
	"fmt"
	"bufio"
	//"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
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

type SNLDB struct {
	snlMap map[string]string
	singleL []time.Time //single axle date/time-stamps
	tandemL []time.Time //tandem axle date/time-stamps
	tripleL []time.Time //triple axle date/time-stamps
	comboL []time.Time //combo truck date/time-stamps
	semiL []time.Time //semi-trailer date/time-stamps

	// singleL *list.List //single axle date/time-stamps
	// tandemL *list.List //tandem axle date/time-stamps
	// tripleL *list.List //triple axle date/time-stamps
	// comboL *list.List //combo truck date/time-stamps
	// semiL *list.List //semi-trailer date/time-stamps
}

func NewSNLDB() *SNLDB {
	f := SNLDB{}
	f.snlMap = make(map[string]string)
	f.singleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	f.tandemL = make([]time.Time, 0, 200)
	f.tripleL = make([]time.Time, 0, 200)
	f.comboL = make([]time.Time, 0, 200)
	f.semiL = make([]time.Time, 0, 200)

	// f.singleL = new(list.List)
	// f.tandemL = new(list.List)
	// f.tripleL = new(list.List)
	// f.comboL = new(list.List)
	// f.semiL = new(list.List)
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
	return string(s.snlMap[SHIFTSTARTTIME])
}

func (s *SNLDB) getShiftEndTime() string {
	return string(s.snlMap[SHIFTENDTIME])
}

func (s *SNLDB) getShiftStartDate() string {
	return string(s.snlMap[SHIFTSTARTDATE])
}

func (s *SNLDB) getGuardName() string {
	return string(s.snlMap[GUARDNAME])
}

func (s *SNLDB) getLicenseNumber() string {
	return string(s.snlMap[LICENSENUMBER])
}

func (s *SNLDB) getShiftComment() string {
	return string(s.snlMap[SHIFTCOMMENT])
}

func (s *SNLDB) getCountLocation() string {
	return string(s.snlMap[COUNTLOCATION])
}

func (s *SNLDB) getCountForItemType() string {
	return string(s.snlMap[COUNTFORITEMTYPE])
}

func (s *SNLDB) setShiftStartTime(s_ string) {
	s.snlMap[SHIFTSTARTTIME] = s_
}

func (s *SNLDB) setShiftEndTime(s_ string) {
	s.snlMap[SHIFTENDTIME] = s_
}

func (s *SNLDB) setShiftStartDate(s_ string) {
	s.snlMap[SHIFTSTARTDATE] = s_
}

func (s *SNLDB) setGuardName(s_ string) {
	s.snlMap[GUARDNAME] = s_
}

func (s *SNLDB) setLicenseNumber(s_ string) {
	s.snlMap[LICENSENUMBER] = s_
}

func (s *SNLDB) setShiftComment(s_ string) {
	s.snlMap[SHIFTCOMMENT] = s_
}

func (s *SNLDB) setCountLocation(s_ string) {
	s.snlMap[COUNTLOCATION] = s_
}

func (s *SNLDB) setCountForItemType(s_ string) {
	s.snlMap[COUNTFORITEMTYPE] = s_
}

func (s *SNLDB) getSingleAxleTotal() int {
	return len(s.singleL)
}

func (s *SNLDB) getTandemAxleTotal() int {
	return len(s.tandemL)
}

func (s *SNLDB) getTripleAxleTotal() int {
	return len(s.tripleL)
}

func (s *SNLDB) getComboTruckTotal() int {
	return len(s.comboL)
}

func (s *SNLDB) getSemiTrailerTotal() int {
	return len(s.semiL)
}

func (s *SNLDB) singleAxleArrived() {
	//s.singleL.PushBack(time.Now())
	s.singleL = append(s.singleL, time.Now())
	// l := len(s.singleL)
	// c := cap(s.singleL)
	// if (l < c) {
	// 	if(l == 0) {
	// 		s.singleL[0]= time.Now()
	// 	} else {
	// 		s.singleL[l-1]= time.Now()			
	// 	}
	// } else {
	// 	s.singleL = append(s.singleL, time.Now())
	// }
}

func (s *SNLDB) tandemAxleArrived() {
	//s.tandemL.PushBack(time.Now())
	s.tandemL = append(s.tandemL, time.Now())
	// l := len(s.tandemL)
	// c := cap(s.tandemL)
	// if (l < c) {
	// 	if(l == 0) {
	// 		s.tandemL[0]= time.Now()
	// 	} else {
	// 		s.tandemL[l-1]= time.Now()			
	// 	}
	// } else {
	// 	s.tandemL = append(s.tandemL, time.Now())
	// }
}

func (s *SNLDB) tripleAxleArrived() {
	//s.tripleL.PushBack(time.Now())
	s.tripleL = append(s.tripleL, time.Now())	
	// l := len(s.tripleL)
	// c := cap(s.tripleL)
	// if (l < c) {
	// 	if(l == 0) {
	// 		s.tripleL[0]= time.Now()
	// 	} else {
	// 		s.tripleL[l-1]= time.Now()			
	// 	}
	// } else {
	// 	s.tripleL = append(s.tripleL, time.Now())
	// }
}

func (s *SNLDB) comboTruckArrived() {
	//s.comboL.PushBack(time.Now())
	s.comboL = append(s.comboL, time.Now())	
	// l := len(s.comboL)
	// c := cap(s.comboL)
	// if (l < c) {
	// 	if(l == 0) {
	// 		s.comboL[0]= time.Now()
	// 	} else {
	// 		s.comboL[l-1]= time.Now()			
	// 	}
	// } else {
	// 	s.comboL = append(s.comboL, time.Now())
	// }

}

func (s *SNLDB) semiTrailerArrived() {
	//s.semiL.PushBack(time.Now())
	s.semiL = append(s.semiL, time.Now())	
	// l := len(s.semiL)
	// c := cap(s.semiL)
	// if (l < c) {
	// 	if(l == 0) {
	// 		s.semiL[0]= time.Now()
	// 	} else {
	// 		s.semiL[l-1]= time.Now()			
	// 	}
	// } else {
	// 	s.semiL = append(s.semiL, time.Now())
	// }
}

func (s *SNLDB) clear() {
	s.snlMap = make(map[string]string)
	s.singleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	s.tandemL = make([]time.Time, 0, 200)
	s.tripleL = make([]time.Time, 0, 200)
	s.comboL = make([]time.Time, 0, 200)
	s.semiL = make([]time.Time, 0, 200)

	// s.singleL = new(list.List)
	// s.tandemL = new(list.List)
	// s.tripleL = new(list.List)
	// s.comboL = new(list.List)
	// s.semiL = new(list.List)
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
	sa.singleaxlebutton.Clicked(func() {
		sa._snldb.singleAxleArrived()
		sa.singleLabel.SetText(strconv.Itoa(sa._snldb.getSingleAxleTotal()));
	})	
	sa.tandemaxlebutton.Clicked(func() {
		sa._snldb.tandemAxleArrived()
		sa.tandemLabel.SetText(strconv.Itoa(sa._snldb.getTandemAxleTotal()));
	})	
	sa.tripleaxlebutton.Clicked(func() {
		sa._snldb.tripleAxleArrived()
		sa.tripleLabel.SetText(strconv.Itoa(sa._snldb.getTripleAxleTotal()));
	})	
	sa.combotruckbutton.Clicked(func() {
		sa._snldb.comboTruckArrived()
		sa.comboLabel.SetText(strconv.Itoa(sa._snldb.getComboTruckTotal()));
	})	
	sa.semitrailerbutton.Clicked(func() {
		sa._snldb.semiTrailerArrived()
		sa.semiLabel.SetText(strconv.Itoa(sa._snldb.getSemiTrailerTotal()));
	})	
	
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

//snldb struct's marshalling hasn't been implemented yet
//so this function is still a buggy skeleton
func (sa *specialAssistant) saveJsonFileSNLDB(myFileName string, f *SNLDB) () {
	fo, err := os.Create(myFileName)
	if err != nil { panic(err) }
	defer fo.Close()
	w := bufio.NewWriter(fo)
	//jsonBytes, err := json.Marshal(mySNLDB)
	jsonBytes, err := json.Marshal([]interface{}{
		f.snlMap,
		f.singleL,
		f.tandemL,
		f.tripleL,
		f.comboL,
		f.semiL,
	})

	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(jsonBytes) 
	if err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil { 
		log.Fatal(err) 
	}
}

//snldb struct's marshalling hasn't been implemented yet
//so this function is still a buggy skeleton
type A struct {
	A map[string]string
}

type B struct {
	B []time.Time
}
	
type C struct {
	C []time.Time
}

type D struct {
	D []time.Time
}

type E struct {
	E []time.Time
}

type F struct {
	F []time.Time
}

func (sa *specialAssistant) readJsonFileSNLDB(myFileName string) (SNLDB) {
	type Any struct {
	 	*A
	 	*B
	 	*C
	 	*D
	 	*E
	 	*F
	}
	var res []Any
	
	//var myS *SNLDB
	//myS = NewSNLDB()
	var myTestSNLDB SNLDB
	// input, err := os.Open(myFileName)
	// if err != nil {
	//   	log.Fatal(err)
	// }

	file, err2 := ioutil.ReadFile(myFileName)
	fmt.Printf("thefilecontents:", file)
	err := json.Unmarshal(file, &res)
	fmt.Printf("Results: %v\n", res, err, err2)

	//err = json.Unmarshal(data, &res)
        // myjsondecoder := json.NewDecoder(input)
	// for {
	//  	err := myjsondecoder.Decode(&myTestSNLDB)
	//  	if err != nil {
	//   		if err == io.EOF {
	//   			break
	//   		}
	//   		log.Fatal(err)
	//   	}
	// }
	return myTestSNLDB
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


func (sa *specialAssistant) close_clicked () {
	println("assistant close clicked page:", sa.v.GetCurrentPage())
	//sa._snldb.debugDataFields()
	//snldb struct's marshalling hasn't been implemented yet
	//so this function is still a buggy skeleton
	// b, err := json.Marshal(sa._snldb.snlMap)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)

	// b, err = json.Marshal(sa._snldb.singleL)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
	// b, err = json.Marshal(sa._snldb.tandemL)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
	// b, err = json.Marshal(sa._snldb.tripleL)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
	// b, err = json.Marshal(sa._snldb.comboL)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
	// b, err = json.Marshal(sa._snldb.semiL)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
	sa.saveJsonFileSNLDB("snowreport.json", sa._snldb)
	tmpblah := sa.readJsonFileSNLDB("snowreport.json")
	fmt.Println("sssssss %v\n", tmpblah)
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


// type A struct {
// 	A string
// }
// type B struct {
// 	B int64
// }
// type C struct {
// 	C float64
// }

// func main() {
// 	data, err := json.Marshal([]interface{}{
// 		A{"foo"},
// 		B{42},
// 		C{1.234},
// 	})
// 	fmt.Printf("%s %v\n", data, err)
// 	err = json.Unmarshal(data, &res)
// 	fmt.Printf("%+v %v\n", res, err)
