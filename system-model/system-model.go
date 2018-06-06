package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("System model service ...")
	http.HandleFunc("/systemmodel", getSystemModel)
	log.Fatal(http.ListenAndServe(":9080", nil))
	fmt.Println("Done.")
}

func getSystemModel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get system model request ...")

	json.NewEncoder(w).Encode("Testing K8S, with ISTIO")
}

