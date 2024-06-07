package main

import (
	"encoding/json"
	"io/ioutil"
)

// Pericias representa as habilidades/skills e seus atributos correspondentes
type Pericias map[string]string

// Função para carregar dados de perícias de um arquivo JSON
func carregarPericias(filename string) (Pericias, error) {
	if filename == "" {
		filename = "data/pericias.json"
	}
	var pericias Pericias
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return pericias, err
	}
	err = json.Unmarshal(data, &pericias)
	return pericias, err
}
