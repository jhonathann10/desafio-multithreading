package dto

type Endereco struct {
	APIName    string `json:"api_name,omitempty"`
	Cep        string `json:"cep,omitempty"`
	Estado     string `json:"estado,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
}
