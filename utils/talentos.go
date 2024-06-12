package utils

import (
	"encoding/json"
	"io/ioutil"
)

// Talentos representa os talentos de um personagem
type Talentos struct {
	TalentosGerais map[string][]string `json:"TalentosGerais"`
}

// Função para carregar dados de talentos de um arquivo JSON
func carregarTalentos(filename string) (Talentos, error) {
	var talentos Talentos
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return talentos, err
	}
	err = json.Unmarshal(data, &talentos)
	if err != nil {
		return talentos, err
	}
	return talentos, nil
}
