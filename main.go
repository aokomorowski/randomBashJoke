package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/", getRandomJokeHandler)

	log.Println("Listening on 80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
