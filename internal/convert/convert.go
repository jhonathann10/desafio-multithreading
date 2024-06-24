package convert

import (
	"encoding/json"

	"github.com/desafio-multithreading/internal/dto"
)

func ConvertToJSON(endereco dto.Endereco) string {
	json, err := json.Marshal(endereco)
	if err != nil {
		return ""
	}
	return string(json)
}
