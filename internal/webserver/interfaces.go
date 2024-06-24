package webserver

import "github.com/desafio-multithreading/internal/dto"

type CEPInterface interface {
	BuscarCEP(cep string, ch chan dto.Endereco)
}

type BrasilAPIService struct {
	Cep          string `json:"cep,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
}

type ViaCEPService struct {
	Cep        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Localidade string `json:"localidade,omitempty"`
	Uf         string `json:"uf,omitempty"`
}
