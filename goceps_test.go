package goceps

import (
	"testing"
)

func TestBuscaEndereco(t *testing.T) {

	_, err := BuscaEndereco("07748415-415")
	if err != nil {
		t.Error()
	}
}
