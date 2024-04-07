package main

import (
	"fmt"
)

type Point struct {
	x_value int
	y_value int
	name    string
}

type bill struct { // define the struct, also here u create a new type.
	name   string             // the name attribute can hold a string.
	items  map[string]float64 // items can hold a map with keys as strings and values of float64.
	tip    float64            // vit variable can hold float64
	colors []string
	Point
}

func newBill(name string) bill { // this function creates a bill object, very similar to a cunstractor in python.
	b := bill{
		name:   name,
		items:  map[string]float64{}, // initialize the items with empty map.
		tip:    0,
		colors: []string{},
		Point: Point{
			x_value: 10,
			y_value: 20,
		}, // initialize the slice.
	}
	return b
}

func main() {
    bill1 := bill{                                        // example 1 of initialization.    ( manual initialization ) . (it is ok to leave some feilds emtpy)
        name: "table_10",                         // this is long initialization - since i am spesifing the keys and adding them values.
        items: map[string]float64{
            "salat": 10.33,
            "burata": 13.2,
            "juice": 5.6,
        },
        tip: 10.1,
        colors: []string{"red", "blue", "green"},    
        Point: Point{                                                     // this is how u initilize a struct. ( aka embedded struct ) 
            x_value: 10,
            y_value: 20,
            name: "point of table_10",
        },
      }
      
        fmt.Println(bill1)                                 //resault:   {table_10 map[burata:13.2 juice:5.6 salat:10.33] 10.1 [red blue green]}
        fmt.Println(bill1.name)                     //resault:    table_10
        fmt.Println(bill1.Point.name)          //resault:    point of table_10
        fmt.Println(bill1.items["salat"])      //resault:    10.33 
		fmt.Println(bill1.x_value)


	// billList := []bill{}              // create empty slice which it's values will be of type bill ( a struc that i created - few lines up).
	// mybill := newBill("mario's bill") // newBill -> this is the function from the above example ( asme as constructor ).
	// fmt.Println(mybill)               // resault:

	// mybill.tip = 10.2
	// mybill.items["salat"] = 25.8
	// mybill.colors = append(mybill.colors, "white", "yellow")
	// mybill.Point.x_value = 11
	// mybill.Point.name = "mario's point bill"
	// fmt.Println(mybill)
	// fmt.Println(mybill.name)
	// fmt.Println(mybill.Point.name)
	// fmt.Println(mybill.x_value)

	

	// billList = append(billList, mybill)
	// fmt.Println(billList)

}
