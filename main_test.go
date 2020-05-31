package main

import "testing"

func TestIntro(t *testing.T) {
	got := Intro()
	want := "Welcome to Remo :)\n\n" +
		"Remo is an app to demo Github releases\n" +
		"Author: Kyle\n" +
		"Date: June 2020\n" +
		"Email: nvietthu@gmail.com\n" +
		"Github: github.com/kyledakid\n"
	if got != want {
		t.Errorf("Got: %q\n Want: %q", got, want)
	}
}

func TestReleaseInfo(t *testing.T) {
}
