package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type joke struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func extractJoke(resp *http.Response) (joke, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
		return joke{}, err
	}

	jokeID := strings.TrimPrefix(doc.Find(".qid").Text(), "#")
	jokeContent := strings.TrimSpace(doc.Find(".quote").Text())

	return joke{jokeID, jokeContent}, nil
}
func getRandomJoke() (joke, error) {
	resp, err := http.Get("http://bash.org.pl/random/")
	if err != nil {
		log.Error(err)
		return joke{}, err
	}
	return extractJoke(resp)
}

func getRandomJokeHandler(w http.ResponseWriter, _ *http.Request) {

	j, err := getRandomJoke()
	if err != nil {
		errMessage := fmt.Sprintf("Can't get a joke from Bash.org: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(errMessage)
		w.Write([]byte(errMessage))
		return
	}

	js, err := parseJokeToJSON(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}
