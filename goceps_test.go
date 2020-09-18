package goceps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	_, err := Search("01001-000")
	if err != nil {
		t.Error(err)
		return
	}
}
