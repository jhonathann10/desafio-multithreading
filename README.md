# desafio-multithreading

## Execução
- `go run cmd/server/main.go <CEP>`

- Exemplo:
```sh
jhonathann-virtuoso@PC-002668 ~/go/src/github.com/desafio-multithreading (main)$ go run cmd/server/main.go 88162033
{"api_name":"ViaCEP","cep":"88162-033","estado":"SC","cidade":"Biguaçu","bairro":"Jardim Janaína","logradouro":"Servidão Olívio João Virtuoso"}
jhonathann-virtuoso@PC-002668 ~/go/src/github.com/desafio-multithreading (main)$ go run cmd/server/main.go 88162033
{"api_name":"BrasilAPI","cep":"88162033","estado":"SC","cidade":"Biguaçu","bairro":"Jardim Janaína","logradouro":"Servidão Olívio João Virtuoso"}
```