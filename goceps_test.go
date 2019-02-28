package goceps

import (
	"testing"
)

func TestBuscaEndereco(t *testing.T) {

	cep := "07748415"

	_, err := BuscaEndereco(cep)
	if err != nil {
		t.Error()
	}
}
