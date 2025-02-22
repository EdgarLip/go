package structfunctions

import (
	"fmt"
	"math"
	//"reflect"
)

// Define a struct
type Person struct {
	Name string
	Age  int
	City string
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

func (p Person) String() string {
	/*
		this function will decide how this object will be printed.
		fmt.Println(Person)
	*/
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

// ***   example 3   ***
type Bill struct {
    Name  string
    Items map[string]float64
    Tip   float64
}


// getEmailOfPerson gets a doiman string like gmail.com 
// returns the full mail address of a person
func (b Bill) GetBillName() string {
	return fmt.Sprintf("%v  --  %v  --  %v", b.Name, b.Items, b.Tip)
}

func (b *Bill) GetBillNameWithPointer() string {
	return fmt.Sprintf("%v  --  %v  --  %v", b.Name, b.Items, b.Tip)
	//return fmt.Sprintf("%v.%v.%v", b.Name, reflect.ValueOf(b.Items).MapKeys(), b.Tip)
}

func (b Bill) UpdateTip(tip float64) {
    //(*b).tip = tip                                               // this will work because u change the order of go's operator priorites.
     b.Tip = tip
}

func (b *Bill) UpdateTipWithPointer(tip float64) {
    //(*b).tip = tip                                               // this will work because u change the order of go's operator priorites.
     b.Tip = tip
}



