package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Revere web gateway service ...")
	http.HandleFunc("/data", getData)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Done.")
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get data request ...")

        var systemModelUrl = "http://system-model-service:8080/systemmodel"
        //var systemModelUrl = "http://localhost:9080/systemmodel"
        fmt.Println("Ready to get model from ", systemModelUrl)
	resp, err := http.Get(systemModelUrl)
	if err != nil {
		fmt.Println("HTTP ERROR: %s ", err)
		panic(err)
        }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("BODY: ", string(body))
	json.NewEncoder(w).Encode(string(body))
}

