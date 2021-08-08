package goceps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	service := NewService()

	t.Run("should be able return a address passing zipcode with special characters", func(t *testing.T) {
		_, err := service.Search("01001-000")
		assert.Nil(t, err)
	})

	t.Run("should be able return a address passing zipcode without special characters", func(t *testing.T) {
		_, err := service.Search("01001000")
		assert.Nil(t, err)
	})

	t.Run("should be able return a address passing a invalid zipcode", func(t *testing.T) {
		_, err := service.Search("99999999")
		assert.NotNil(t, err)
	})
}
