package goceps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	service := NewService()

	_, err := service.Search("01001-000")
	if err != nil {
		t.Error(err)
		return
	}
}
