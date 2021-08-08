package goceps

type Address struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IBGE         string `json:"ibge"`
	GIA          string `json:"gia"`
	DDD          string `json:"ddd"`
	Siafi        string `json:"siafi"`
}

type AddressError struct {
	HaveError bool `json:"erro"`
}
