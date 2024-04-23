package mapsfunctions

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Basic_map() {
	dict1 := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(dict1)                   // resault:     map[pie:7.99 salad:6.99 soup:4.99 toffee pudding:3.55]
	fmt.Println(dict1["pie"])            // resault:      7.99
	fmt.Println(dict1["nonExistingKey"]) // resault:     0

	v, key_exists := dict1["nonExistingKey"]              // this is how u check is the key exists.
	fmt.Printf("v: %v, key_exists: %v \n", v, key_exists) // v: 0, key_exists: false
	v, key_exists = dict1["salad"]
	fmt.Printf("v: %v, key_exists: %v \n", v, key_exists) // v: 6.99, key_exists: true

	if v, key_exists := dict1["nonExistingKey"]; key_exists { // assign values to v and key_exists , and then check if the value is true/false
		fmt.Println(v) //  this is a very good example.
	} else { //   this is also the concept of the 'Comma ok idiom'
		fmt.Println("key does not exists !")
	}
}

func For_loop_on_map() {
	fmt.Println("------------  for loop over map basic 1 ------------------")
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for key, val := range m {
		fmt.Printf("ranging over a map m ,key: %v , val: %v\n", key, val)
	}
}

func Basic_marshal() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	bs, err := json.Marshal(menu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	json1 := fmt.Sprintf("%s\n", string(bs))
	fmt.Println(json1)
}

type person struct {
	Firts string `json:"first_name,omitempty"`
	Last  string `json:"last_name,omitempty"`
}

func Basic_unmarshal1() {
	jsonCheck := `{"first_name":"edgar","last_name":"lip"}`
	personFromJson := person{}
	err := json.Unmarshal([]byte(jsonCheck), &personFromJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("personFromJson:\n%v\n", personFromJson) // resault:  {edgar lip}
}

func Basic_unmarshal2() {
	jsonCheck := `[{"first_name":"edgar","last_name":"lip"},{"first_name":"nana","last_name":"kushnir"}]`
	personSlice := []person{}
	err := json.Unmarshal([]byte(jsonCheck), &personSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Json1:\n%v\n", personSlice) // resault will be : [{edgar lip} {nana kushnir}]

	for indx, val := range personSlice { // just another way to print it...
		fmt.Printf("Person: %v Full name: %v, %v\n", indx+1, val.Firts, val.Last)
	}
}

func InterfacedMapExampleInAction() {
	dict := make(map[interface{}]interface{}) // Creating a map with keys and values of type interface{}

	dict[1] = "one" // Adding elements to the map with different types of keys and values
	dict["two"] = 2
	dict[true] = false
	dict[3.14] = []int{1, 2, 3}

	fmt.Printf("InterfacedMapExample dict[1]: %v\n", dict[1])         // Output: one
	fmt.Printf("InterfacedMapExample dict['two']: %v\n", dict["two"]) // Output: 2
	fmt.Printf("InterfacedMapExample dict[true]: %v\n", dict[true])   // Output: false
	fmt.Printf("InterfacedMapExample dict[3.14]: %v\n", dict[3.14])   // Output: [1 2 3]
}

func InterfacedMapExampleInActionV2() {
	dict := make(map[interface{}]interface{}) // Creating a map with keys and values of type interface{}
	dict["two"] = 2

	dict["ppl"] = make(map[interface{}]interface{}) // Initialize a nested map under the key "ppl"

	pplMap := dict["ppl"].(map[interface{}]interface{}) // Populate the nested map with a value

	fmt.Printf("NESTED-dict - pplMap value after assertion-1: %v\n", pplMap)                // NESTED-dict - pplMap value after assertion-1: map[]
	fmt.Printf("NESTED-dict - dict['ppl'] value after asserting pplMap: %v\n", dict["ppl"]) // NESTED-dict - dict['ppl'] value after asserting pplMap: map[]

	fmt.Printf("NESTED-dict - mem addr of dict['ppl']: %v\n", reflect.ValueOf(dict["ppl"]).Pointer())
	fmt.Printf("NESTED-dict - mem addr of pplMap: %v\n", reflect.ValueOf(pplMap).Pointer())
	fmt.Printf("NESTED-dict - mem addr of pplMap: %p\n", &pplMap)

	pplMap["edgar"] = "amazing"

	fmt.Printf("NESTED-dict - pplMap after adding some value: %v\n", pplMap) // NESTED-dict - pplMap after adding some value: map[edgar:amazing]

	fmt.Printf("NESTED-dict: dict value: %v\n", dict) // NESTED-dict: dict value: map[ppl:map[edgar:amazing] two:2]
}

func InterfacedMapExampleInActionV3() {
	dict := make(map[interface{}]interface{}) // Creating a map with keys and values of type interface{}
	dict["two"] = 2

	dict["ppl"] = make(map[interface{}]interface{}) // Initialize a nested map under the key "ppl"

	dictPpl := dict["ppl"].(map[interface{}]interface{}) // assert type the nested map to variable with value.
	dictPpl["edgar"] = "amazing"
	dictPpl["student"] = make(map[interface{}]interface{})

	dictPplStudents := dictPpl["student"].(map[interface{}]interface{})
	dictPplStudents["nomi"] = "working"
	fmt.Printf("NESTED-dict x3 -%v\n", dict)
}
