package main

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}
func parseJoketoJSON(j joke) ([]byte, error) {

	js, jsonerr := JSONMarshal(j, true)
	if jsonerr != nil {
		log.Error(jsonerr)
	}
	return js, jsonerr
}
