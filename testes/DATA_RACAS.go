package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Raca representa as informacoes de uma raca no sistema
type Raca struct {
	AtributoChave     string   `json:"atributo_chave"`
	TalentoAscendente string   `json:"talento_ascendente"`
	ProfissoesTipicas []string `json:"profissoes_tipicas"`
	Idades            Idades   `json:"idades"`
}

// Idades representa as faixas etarias para uma raca
type Idades struct {
	Jovem  *[2]int `json:"Jovem"`
	Adulto [2]int  `json:"Adulto"`
	Idoso  *[2]int `json:"Idoso"`
}

// PersonagemRacas armazena todas as racas e suas infos
type PersonagemRacas struct {
	Racas     []string
	RacasInfo map[string]Raca
}

// Função para carregar dados de raças e idades de um arquivo JSON unificado
func carregarRacas(filename string) (map[string]Raca, error) {
	var racas map[string]Raca
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return racas, err
	}
	err = json.Unmarshal(data, &racas)
	return racas, err
}

// NewPersonagemRacas inicializa PersonagemRacas com valores carregados do arquivo JSON unificado.
func NewPersonagemRacas(classes []string, filename string) (*PersonagemRacas, error) {
	if filename == "" {
		filename = "data/racas.json"
	}
	racas, err := carregarRacas(filename)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar raças e idades: %v", err)
	}

	racasNomes := make([]string, 0, len(racas))
	for nome := range racas {
		racasNomes = append(racasNomes, nome)
	}

	return &PersonagemRacas{
		Racas:     racasNomes,
		RacasInfo: racas,
	}, nil
}
