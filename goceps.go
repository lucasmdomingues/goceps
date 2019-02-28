package goceps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Endereco struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Unidade     string `json:"unidade"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
}

func BuscaEndereco(cep string) (*Endereco, error) {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var endereco *Endereco

	err = json.Unmarshal(body, &endereco)
	if err != nil {
		return nil, err
	}

	return endereco, err
}
