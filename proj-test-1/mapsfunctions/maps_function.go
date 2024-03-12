package mapsfunctions

import "fmt"

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

	v, key_exists := dict1["nonExistingKey"]            // this is how u check is the key exists.
	fmt.Printf("v: %v, key_exists: %v ", v, key_exists) // v: 0, key_exists: false
	v, key_exists = dict1["salad"]
	fmt.Printf("v: %v, key_exists: %v ", v, key_exists) // v: 6.99, key_exists: true

	if v, key_exists := dict1["nonExistingKey"]; key_exists { // assign values to v and key_exists , and then check if the value is true/false
		fmt.Println(v) //  this is a very good example.
	} else { //   this is also the concept of the 'Comma ok idiom'
		fmt.Println("key does not exists !")
	}
}
