package main

import (
	"fmt"
	"os"
	"time"

	"github.com/desafio-multithreading/internal/convert"
	"github.com/desafio-multithreading/internal/dto"
	"github.com/desafio-multithreading/internal/webserver"
)

func main() {
	for _, cep := range os.Args[1:] {
		ch := make(chan dto.Endereco)

		brasilAPIService := webserver.BrasilAPIService{}
		viaCEPService := webserver.ViaCEPService{}

		go brasilAPIService.BuscarCEP(cep, ch)
		go viaCEPService.BuscarCEP(cep, ch)

		select {
		case endereco := <-ch:
			enderecoJSON := convert.ConvertToJSON(endereco)
			fmt.Println(enderecoJSON)
		case endereco := <-ch:
			enderecoJSON := convert.ConvertToJSON(endereco)
			fmt.Println(enderecoJSON)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
		}
	}
}
