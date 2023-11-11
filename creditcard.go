package main

//https://go.dev/doc/tutorial/call-module-code

import (
	"fmt"
	"strconv"
	"strings"

	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"

	c "example.com/cred/creditvalidty"
)

type requestBody struct {
	Value string `json: value`
}

type responseBody struct {
	Status string `json: status`
}

func main() {
	fmt.Println("Welcome to Credit Card Validator.")

	r := mux.NewRouter()
	// curl -X GET -d '{"value":"8,9,9,2,7,3,9,8,7,9,3"}'  "http://127.0.0.1:8089/cred/"
	// {"Status":"OK"}

	// 	curl -X GET -d '{"value":"0,9,9,2,7,3,9,8,7,9,3"}'  "http://127.0.0.1:8089/cred/"
	// {"Status":"FAIL"}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	r.HandleFunc("/cred/", handleGetWithRequestBody).Methods("GET")

	r.HandleFunc("/credhello/", handler).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8089", nil)

	// cardNumberGiven := &[]int{7, 9, 9, 2, 7, 3, 9, 8, 7, 9, 3}

	// err := c.CheckCard(cardNumberGiven)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(" Card is Valid")
	// }

}
func stringToIntList(input string, delimiter string) ([]int, error) {
	// Split the string into a slice of substrings
	strValues := strings.Split(input, delimiter)
	fmt.Println(delimiter, "teju")

	// Convert each substring to an integer
	var intValues []int
	for _, str := range strValues {
		// Convert string to int
		intValue, err := strconv.Atoi(str)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		// Append to the list
		intValues = append(intValues, intValue)
	}

	fmt.Println("jjjj", intValues)
	return intValues, nil

}

func handleGetWithRequestBody(w http.ResponseWriter, r *http.Request) {

	fmt.Println("teju", r.Body)

	// Decode the JSON request body
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	CredNumber, _ := stringToIntList(req.Value, ",")

	fmt.Println(req.Value, CredNumber)

	err = c.CheckCard(&CredNumber)

	var resp1 responseBody
	if err != nil {
		resp1.Status = "OK"
	} else {
		resp1.Status = "FAIL"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp1)

}
