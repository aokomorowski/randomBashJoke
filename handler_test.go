package main

import (
	"net/http"
	"testing"
)

func TestExtractJoke(t *testing.T) {
	resp, _ := http.Get("http://bash.org.pl/400561/")

	got := extractJoke(resp)
	want := joke{"400561", `<krosnoPL> został u was akumulator do kamery

<dwP0L> znalazlem

<dwP0L> li-ion

<dwP0L> Panasonic

<dwP0L> CGR-V610

<krosnoPL> ten;]

<krosnoPL> takze przechowajcie go w odpowiednich warunkach

<dwP0L> kapuje

<dwP0L> idzie do woreczka i lodówka :>

<krosnoPL> no i głaszcz go codziennie`}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
