package structfunctions

import "fmt"

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
