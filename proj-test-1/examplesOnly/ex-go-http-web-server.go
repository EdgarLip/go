package examplesOnly

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)


// this program is a cute little http api server - for learning purposes = )
// NOTE: don't do this in real life
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// how to use:  curl http://127.0.0.1:8091/list
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// how to use: curl "http://127.0.0.1:8091/add?item=ties&price=10.0" -i
//
//	OR
//	curl http://127.0.0.1:8091/add?item=ties\&price=10.0  -i
func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusForbidden) // StatusForbidden=403
		fmt.Fprintf(w, "duplicate item: %q\n", item)
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
	} else if item == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Invalid item name, can't be empty\n")
	} else {
		db[item] = dollars(f64)
		fmt.Fprintf(w, "added %s with price %s\n", item, dollars(f64))
	}
}

// how to use: curl "http://127.0.0.1:8091/update?item=asd&price=17.0"
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusForbidden) // StatusForbidden=403
		fmt.Fprintf(w, "Cant update non exiting item : %q\n", item)
		return
	}
	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
	} else {
		db[item] = dollars(f64)
		fmt.Fprintf(w, "updated %s with price %s\n", item, dollars(f64))
	}
}
//how to use: curl http://127.0.0.1:8091/get?'item=socks'
func (db database) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusForbidden) // StatusForbidden=403
		fmt.Fprintf(w, "Cant get non exiting item : %q\n", item)
		return
	}

	fmt.Fprintf(w, "item %s price %s\n", item, dollars(db[item]))
}

//how to use: curl http://127.0.0.1:8091/delete?'item=socks'
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusForbidden) // StatusForbidden=403
		fmt.Fprintf(w, "Cant delete non exiting item : %q\n", item)
		return
	}

	delete(db, item)

	fmt.Fprintf(w, "item %s was deleted !\n", item)

}

func RunHttpServer() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/get", db.get)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8091", nil))
}
