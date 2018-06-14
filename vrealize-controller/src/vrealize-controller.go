package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
  log.Println("Starting vrealize controller ...")
  http.HandleFunc("/data", getData)
  http.HandleFunc("/install", installSuite)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting data from System Standards ...")

	systemStandardsServiceUrl := os.Getenv("SYSTEM_STANDARDS_SERVICE_URL")

	var systemStandardsUrl = systemStandardsServiceUrl + "/data"
	log.Println("Ready to get data from ", systemStandardsUrl)
	resp, err := http.Get(systemStandardsUrl)
	if err != nil {
		log.Fatal("HTTP ERROR: %s ", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Println("BODY: ", string(body))
	json.NewEncoder(w).Encode(string(body))
}

func installSuite(w http.ResponseWriter, r *http.Request) {
	log.Println("Installing vRealize suite ...")

	awxUrl := os.Getenv("AWX_URL")

	log.Println("Ready to install vrealize suite ", awxUrl)
	resp, err := http.Get(awxUrl)
	if err != nil {
		log.Fatal("HTTP ERROR: %s ", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Println("BODY: ", string(body))
	json.NewEncoder(w).Encode(string(body))
}