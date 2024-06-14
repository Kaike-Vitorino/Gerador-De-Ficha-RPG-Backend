package character_data

import (
	"encoding/json"
	"os"
)

// Atributos representa os atributos de um character_logic
type Atributos struct {
	Forca        []int `json:"Forca"`
	Agilidade    []int `json:"Agilidade"`
	Inteligencia []int `json:"Inteligencia"`
	Empatia      []int `json:"Empatia"`
}

// Função para carregar dados de atributos de um arquivo JSON
func CarregarAtributos(filename string) (Atributos, error) {
	if filename == "" {
		filename = "data/atributos.json"
	}
	var atributos Atributos
	data, err := os.ReadFile(filename)
	if err != nil {
		return atributos, err
	}
	err = json.Unmarshal(data, &atributos)
	return atributos, err
}
