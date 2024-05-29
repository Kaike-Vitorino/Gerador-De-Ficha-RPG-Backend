package main

import (
	"encoding/json"
	"io/ioutil"
)

// Talentos representa os talentos e as classes que podem ter eles
type Talentos map[string][]string

// Função para carregar dados de talentos de um arquivo JSON
func carregarTalentos(filename string) (Talentos, error) {
	if filename == "" {
		filename = "data/talentos.json"
	}
	var talentos Talentos
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return talentos, err
	}
	err = json.Unmarshal(data, &talentos)
	return talentos, err
}
