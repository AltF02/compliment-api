package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)


func getCompliment(w http.ResponseWriter, r *http.Request)  {
	complimentsJson, err := ioutil.ReadFile("compliments.json")
	if err != nil {
		fmt.Println(err)
	}

	var compliments []string
	_ = json.Unmarshal(complimentsJson, &compliments)
	rand.Seed(time.Now().Unix())
	_, err = fmt.Fprintf(w, compliments[rand.Intn(len(compliments))])
	if err != nil {
		fmt.Println(err)
	}
}


func homeLink(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the compliment api, please visit github.com/DankDumpster/Compliment-api for more info")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/compliment", getCompliment).Methods("GET")
	fmt.Println("Server running on http://localhost:8090/")

	Port := ":8090"
	if os.Getenv("PORT") != "" {
		Port = ":" + os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(Port, router))
}