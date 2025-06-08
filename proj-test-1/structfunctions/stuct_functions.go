package structfunctions

import (
	"fmt"
	"math"
	//"strings"
	//"reflect"
)

type Point struct { // new struct type.
	X_value int
	Y_value int
	Name    string
}

type Bill struct { // define the struct, also here u create a new type.
	Name    string             // the name attribute can hold a string.
	Items   map[string]float64 // items can hold a map with keys as strings and values of float64.
	Tip     float64            // vit variable can hold float64
	Colors  []string
	SubBill *Bill //
	Point         // this is called embeded struct  ( struct within a struct )
}

// Define a struct
type Person struct {
	Name string
	Age  int
	City string
}

func PrintBillData() {
	// long struct initialization
	bill1 := Bill{
		Name: "table_10",
		Items: map[string]float64{
			"salat":  10.33,
			"burata": 13.2,
			"juice":  5.6,
		},
		Tip:     10.1,
		Colors:  []string{"red", "blue", "green"},
		SubBill: &Bill{Name: "Child Bill", Tip: 19.2},
		Point: Point{
			X_value: 10,
			Y_value: 20,
			Name:    "point of table_10",
		},
	}

	fmt.Println(bill1)      //resault:   {table_10 map[burata:13.2 juice:5.6 salat:10.33] 10.1 [red blue green]}
	fmt.Println(bill1.Name) //resault:    table_10
	fmt.Println(bill1.SubBill)
	fmt.Println(bill1.SubBill.Name)   // resault:   Child Bill
	fmt.Println(bill1.Items["salat"]) //resault:    10.33
	fmt.Println(bill1.Point.Name)     //resault:    point of table_10
	fmt.Println(bill1.X_value)        //resault:    10
}

func WorkWithMapOfBills() {
	//bills1 := map[string]*bill{}
	//bills2 := map[string]bill{}

}

// Function that modifies a copy of the struct
func modifyPerson(p Person) {
	p.Name = "John Doe"
	p.Age = 30
	p.City = "New York"
}

// Function that modifies the original struct using a pointer
func modifyPersonByPointer(p *Person) {
	p.Name = "Jane Smith"
	p.Age = 25
	p.City = "San Francisco"
}

/*
this function will decide how this object will be printed.
fmt.Println(Person)
*/
func (p Person) String() string {
	return fmt.Sprintf("name: %v, Age: %v", p.Name, p.Age)
}

// getEmailOfPerson gets a doiman string like gmail.com
// returns the full mail address of a person
func (p Person) GetEmailOfPerson(domain string) string {
	return fmt.Sprintf("%v.%v.%v@%v", p.Name, p.Age, p.City, domain)
}

func (p *Person) GetEmailOfPersonWithPointer(domain string) string {
	return fmt.Sprintf("%v.%v.%v@%v", p.Name, p.Age, p.City, domain)
}

func ModifyPersonInAction() {
	fmt.Println("---- Running ModifyPersonInAction ---- ")
	person1 := Person{
		Name: "testPerson",
		Age:  28,
		City: "test city",
	}
	fmt.Println("Person: ", person1)
	modifyPerson(person1)

	// Person remains unchanged because modifyPerson operates on a copy
	fmt.Println("Person after modifyPerson:", person1)

	// Call modifyPersonByPointer - passing a pointer to the struct
	modifyPersonByPointer(&person1)

	// Person is modified because modifyPersonByPointer modifies the original struct
	fmt.Println("Person after modifyPersonByPointer:", person1)

	fmt.Println("---- Running ModifyPersonInAction ---- ")
}

// ***   example 1   ***
func (p Person) personSayHiReceiver() {
	fmt.Printf("HI, name is: %v\n", p.Name)
}

func (p *Person) personSayHiReceiverByPointer() {
	fmt.Printf("HI, name is: %v\n", p.Name)
}

func PersonSayHiInAction() {
	person1 := Person{
		Name: "Person1NonPointer",
		Age:  28,
		City: "test city",
	}

	person2 := &Person{
		Name: "Person2Pointer",
		Age:  28,
		City: "test city",
	}

	person1.personSayHiReceiver()
	person1.personSayHiReceiverByPointer()

	person2.personSayHiReceiver()
	person2.personSayHiReceiverByPointer()
}

//   ***   end of example 1   ***

// ***   example 2   ***
type circle struct {
	radius float64
}
type shape interface { // create new type "shape" which implemets 2 methods
	area() float64            //   which ever struct implemets those methods
	areaWithPointer() float64 // he will be also ( bbecomes also ) of type "shape".
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) areaWithPointer() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("area", s.areaWithPointer())
}

func InfoInAction() {
	c := circle{5}
	//info(c) // not wroking
	info(&c) // working !
}

//   ***   end of example 2   ***

// this funtion overrides how the Bill object is being printed.
func (b Bill) String() string {
	return fmt.Sprintf("%v  --  %v  --  %v", b.Name, b.Items, b.Tip)
}

// this funtion overrides how the Point object is being printed.
func (p Point) String() string {
	return fmt.Sprintf("%v  --  %v  --  %v", p.Name, p.X_value, p.Y_value)
}

// ***   example 3   ***

// getEmailOfPerson gets a doiman string like gmail.com
// returns the full mail address of a person
func (b Bill) GetBillName() string {
	return fmt.Sprintf("%v  --  %v  --  %v", b.Name, b.Items, b.Tip)
}

func (b *Bill) GetBillNameWithPointer() string {
	return fmt.Sprintf("%v  --  %v  --  %v", b.Name, b.Items, b.Tip)
	//return fmt.Sprintf("%v.%v.%v", b.Name, reflect.ValueOf(b.Items).MapKeys(), b.Tip)
}

// this will work because u change the order of go's operator priorites.
// but this will not update the tip of the struct that u need to.
// since this is a value receiver.
func (b Bill) UpdateTip(tip float64) {
	//(*b).tip = tip
	b.Tip = tip
}

// this will work because u change the order of go's operator priorites.
// and it will cahnge the tip of the bill - since this is a pointer receiver.
func (b *Bill) UpdateTipWithPointer(tip float64) {
	//(*b).tip = tip
	b.Tip = tip
}

//line struct example.
type Line struct {
	Begin, End Point
}

type Path []Point

// example how to calculate a distance between 2 points with line object.
func (l Line) Distance() float64 {
	return math.Hypot(float64(l.End.X_value)-float64(l.Begin.X_value), float64(l.End.Y_value)-float64(l.Begin.Y_value))
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}
