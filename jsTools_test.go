package main

import (
	"reflect"
	"testing"
)

func TestParseJoketoJSON(t *testing.T) {
	j := joke{"123123", "lorem ipsum"}
	got, _ := parseJokeToJSON(j)
	want := []byte(`{"id":"123123","content":"lorem ipsum"}`)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
