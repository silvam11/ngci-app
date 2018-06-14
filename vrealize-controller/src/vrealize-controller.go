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

	awxUsername := os.Getenv("AWX_USERNAME")
	awxPassword := os.Getenv("AWX_PASSWORD")

	awxUrl := os.Getenv("AWX_URL")
	awxEndpoint := awxUrl + "/api/v2/users/"

	log.Println("Ready to install vrealize suite ", awxEndpoint)
	//resp, err := http.Get(awxEnpoint)
	req, err := http.NewRequest("GET", awxEndpoint, nil)
	req.SetBasicAuth(awxUsername, awxPassword)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal("HTTP ERROR: %s ", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Println("BODY: ", string(body))
	json.NewEncoder(w).Encode(string(body))
}