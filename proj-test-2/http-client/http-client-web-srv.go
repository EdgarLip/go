package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const server = "http://localhost:8091"

var items = map[string]string{
	"hat": "33.5",
	"shoes": "50",
	"socks": "5",
	"ties": "10",
	"shirt": "20",
	"pants": "30",
	"jacket": "40",
	"dress": "50",
	"scarf": "60",
	"gloves": "70",
	"belt": "80",
	"wallet": "90",
	"watch": "100",
	"ring": "110",
	"earrings": "120",
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

// ListItems performs a GET to /list
func ListItems(server string) (string, error) {
	return httpGet(server + "/list")
}

// AddItem performs a GET to /add?item=...&price=...
func AddItem(server, item, price string) (string, error) {
	params := url.Values{}
	params.Set("item", item)
	params.Set("price", price)
	return httpGet(server + "/add?" + params.Encode())
}

// UpdateItem performs a GET to /update?item=...&price=...
func UpdateItem(server, item, price string) (string, error) {
	params := url.Values{}
	params.Set("item", item)
	params.Set("price", price)
	return httpGet(server + "/update?" + params.Encode())
}

// GetItem performs a GET to /get?item=...
func GetItem(server, item string) (string, error) {
	params := url.Values{}
	params.Set("item", item)
	return httpGet(server + "/get?" + params.Encode())
}

// DeleteItem performs a GET to /delete?item=...
func DeleteItem(server, item string) (string, error) {
	params := url.Values{}
	params.Set("item", item)
	return httpGet(server + "/delete?" + params.Encode())
}

// HttpServerClientDemo demonstrates using all client operations.

func HttpServerClientDemo() {
	for {
		for item, price := range items {
			go ResponseHandler(AddItem(server, item, price))
		}
		for item, _ := range items {
			go ResponseHandler(GetItem(server, item))
		}
		for item, price := range items {
			go ResponseHandler(UpdateItem(server, item, price))
		}

		go ResponseHandler(ListItems(server))

		for item, _ := range items {
			go ResponseHandler(DeleteItem(server, item))
		}
	}
}

func ResponseHandler(body string, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Print(body)
	}
}


