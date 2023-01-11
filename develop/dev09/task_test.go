package main

import (
	"testing"
)

const (
	correct     = "https://www.google.com"
	wrongURL    = "htps://www.google.com"
	wrongDomain = "https://www.gog228le.com"
)

func TestPath(t *testing.T) {
	t.Run("CorrectURL", func(t *testing.T) {
		if _, err := download(correct); err != nil {
			t.Errorf("error: %v]n", err)
		}
	})
	t.Run("WrongURL", func(t *testing.T) {
		if _, err := download(wrongURL); err != nil {
			t.Errorf("error: %v\n", err)
		}
	})
	t.Run("wrongDomain", func(t *testing.T) {
		if _, err := download(wrongDomain); err != nil {
			t.Errorf("error: %v\n", err)
		}
	})
}
