package main

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type joke struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func extractJoke(resp *http.Response) joke {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
	}

	jokeID := strings.TrimPrefix(doc.Find(".qid").Text(), "#")
	jokeContent := strings.TrimSpace(doc.Find(".quote").Text())

	return joke{jokeID, jokeContent}
}
func getRandomJoke() joke {
	resp, err := http.Get("http://bash.org.pl/random/")
	if err != nil {
		log.Error(err)
	}
	return extractJoke(resp)
}

func getRandomJokeHandler(w http.ResponseWriter, r *http.Request) {

	j := getRandomJoke()
	js := parseJoketoJSON(j)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
