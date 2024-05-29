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
}

// IdadeRaca representa as faixas etarias para uma raca
type IdadeRaca struct {
	Jovem  *[2]int `json:"Jovem"`
	Adulto [2]int  `json:"Adulto"`
	Idoso  *[2]int `json:"Idoso"`
}

// PersonagemRacas armazena todas as racas e suas infos
type PersonagemRacas struct {
	Racas      []string
	RacasInfo  map[string]Raca
	IdadeRacas map[string]IdadeRaca
}

// Func para carregar dados das racas do arquivo JSON
func carregarRacas(filename string) (map[string]Raca, error) {
	var racas map[string]Raca
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return racas, err
	}
	err = json.Unmarshal(data, &racas)
	return racas, err
}

// Func para carregar dados das idades do arquivo JSON
func carregarIdades(filename string) (map[string]IdadeRaca, error) {
	var idades map[string]IdadeRaca
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return idades, err
	}
	err = json.Unmarshal(data, &idades)
	return idades, err
}

// NewPersonagemRacas inicializa PersonagemRacas com valores carregados dos arquivos JSON.
func NewPersonagemRacas(classes []string, racasFile, idadesFile string) (*PersonagemRacas, error) {
	racas, err := carregarRacas(racasFile)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar raças: %v", err)
	}

	idades, err := carregarIdades(idadesFile)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar idades: %v", err)
	}

	racasNomes := make([]string, 0, len(racas))
	for nome := range racas {
		racasNomes = append(racasNomes, nome)
	}

	return &PersonagemRacas{
		Racas:      racasNomes,
		RacasInfo:  racas,
		IdadeRacas: idades,
	}, nil
}
