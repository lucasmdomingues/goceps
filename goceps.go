package goceps

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
)

type service struct {
	BaseURL url.URL
}

type Service interface {
	Search(zipcode string) (*Address, error)
}

func NewService() Service {
	return &service{
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "viacep.com.br",
			Path:   "ws/",
		},
	}
}

func (s *service) Search(zipcode string) (*Address, error) {
	exp := regexp.MustCompile(`[^0-9]`)
	zipcode = exp.ReplaceAllString(zipcode, "")

	if len(zipcode) != 8 {
		return nil, errors.New("zipcode must be 8 characters")
	}

	url, err := s.BaseURL.Parse(path.Join(zipcode, "json"))
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var addressError *AddressError

	err = json.Unmarshal(body, &addressError)
	if err != nil {
		return nil, err
	}

	if addressError.HaveError {
		return nil, errors.New("zipcode not found or invalid")
	}

	var address *Address

	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return address, err
}
