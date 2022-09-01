package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whomami struct {
	Name  string
	Role  string
	State string
}

func main() {
	request1()
}

func whoami(response http.ResponseWriter, r *http.Request) {
	who := []whomami{
		whomami{Name: "Sindhu",
			Role:  "Associate Software Engineer",
			State: "AP",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit : homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Sindhu"

	fmt.Fprintf(response, "Some details about Sindhu..... ")
	fmt.Println("Endpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoami)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
