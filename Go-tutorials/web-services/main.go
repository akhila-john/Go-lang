package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type fooHandler struct {
	message string
}
type Products struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Colour string  `json:"colour"`
	Price  float64 `json:"price"`
}

var productList = []Products{
	{ID: 1, Name: "toy", Colour: "pink", Price: 26.99},
	{ID: 2, Name: "chair", Colour: "blue", Price: 78.34},
	{ID: 3, Name: "flower", Colour: "green", Price: 78.90},
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte)(f.message))

}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar calledd"))
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ID {
			highestID = product.ID
		}
	}
	return highestID + 1

}

// they accept a ResponseWriter and a pointer to a Request
// http.ResponseWriter--used for sending responses to any connected HTTP clients. It's also how response headers are set.
//  http.Request--It's how data is retrieved from the web request.
func getProducts(w http.ResponseWriter, r *http.Request) {
	product, err := json.Marshal(productList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	//Write-- writes the data to the connection as part of an HTTP response.
	w.Write((product))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Products
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = json.Unmarshal(bodyBytes, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	productList = append(productList, newProduct)
	pdt := json.Marshal(productList)
	w.Write(pdt)
}

func main() {
	http.Handle("/foo", &fooHandler{message: "foo called"})
	//  http.HandleFunc method, which provides a way to specify how requests to a specific route should be handled.
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/get", getProducts)
	http.HandleFunc("/post", createProduct)
	http.ListenAndServe(":5000", nil)

}
