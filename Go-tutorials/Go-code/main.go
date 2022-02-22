package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	fmt.Println("hello world")

	var i int = 3
	fmt.Println(i)

	name := "anne"
	fmt.Println(name)

	u := models.User{
		ID:        3,
		Firstname: "dane",
		Lastnsme:  "thomas",
	}
	fmt.Println(u)

	port := 3000
	_, err := webServer(port, 4)
	fmt.Println(err)
}

func webServer(port int, entries int) (int, error) {
	fmt.Println("server started at", port)
	fmt.Println("server started at", entries)
	return port, nil
}
