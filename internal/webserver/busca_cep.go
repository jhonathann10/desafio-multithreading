package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/desafio-multithreading/internal/dto"
)

func GetRequest(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return io.ReadAll(req.Body)
}

func (bas *BrasilAPIService) BuscarCEP(cep string, ch chan dto.Endereco) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	res, err := GetRequest(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	err = json.Unmarshal(res, &bas)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	ch <- dto.Endereco{
		APIName:    "BrasilAPI",
		Cep:        bas.Cep,
		Estado:     bas.State,
		Cidade:     bas.City,
		Bairro:     bas.Neighborhood,
		Logradouro: bas.Street,
	}
}

func (vcs *ViaCEPService) BuscarCEP(cep string, ch chan dto.Endereco) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := GetRequest(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	err = json.Unmarshal(res, &vcs)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	ch <- dto.Endereco{
		APIName:    "ViaCEP",
		Cep:        vcs.Cep,
		Estado:     vcs.Uf,
		Cidade:     vcs.Localidade,
		Bairro:     vcs.Bairro,
		Logradouro: vcs.Logradouro,
	}
}
