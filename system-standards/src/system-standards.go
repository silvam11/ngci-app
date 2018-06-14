package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type WorkflowAttributes struct {
  Name string
  Type string
  Required bool
}

func main() {
	log.Printf("Starting system standards ...")
	http.HandleFunc("/data", getData)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting data from System Standards ...")


	json.NewEncoder(w).Encode(WorkflowAttributes{"host", "string", true})
}

