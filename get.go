package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(URL string) string {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(string(body))
	return string(body)
}
