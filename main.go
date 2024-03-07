package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const portNumber string = ":8000"

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Hello world from yedu")
	case "POST":
		d := json.NewDecoder(r.Body)

		c := struct {
			Card *string `json:"card"`
		}{}

		err := d.Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(*c.Card) == 16 {

			la := luhnAlgorithm(*c.Card)

			if la {
				fmt.Fprint(w, "The card is valid ")
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)

			}
		} else {
			http.Error(w, "Invalid card", http.StatusBadRequest)

		}

	default:
		fmt.Fprint(w, "does not exist")
	}
}

func luhnAlgorithm(cn string) bool {
	log.Println("Running algorithm")

	digits := []int{}
	for i := 0; i < len(cn); i++ {
		digits = append(digits, int(cn[i]-'0'))
	}
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] = digits[i] * 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum = sum + digits[i]
	}
	if sum%10 == 0 {
		return true

	} else {
		return false
	}
}
func main() {

	log.Println("Running programme")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(portNumber, mux)
	if err != nil {
		log.Println("An error happened", err)
		log.Fatal(err)
	}
}
