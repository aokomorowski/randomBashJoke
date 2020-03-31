package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/", getRandomJokeHandler)

	log.Println("Listening on 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
