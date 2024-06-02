package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// Raca representa as informacoes de uma raca no sistema
type Raca struct {
	AtributoChave     string   `json:"atributo_chave"`
	TalentoAscendente string   `json:"talento_ascendente"`
	ProfissoesTipicas []string `json:"profissoes_tipicas"`
	Idades            struct {
		Jovem  *[2]int `json:"jovem"`
		Adulto [2]int  `json:"adulto"`
		Idoso  *[2]int `json:"idoso"`
	} `json:"idades"`
}

// PersonagemRacas armazena todas as racas e suas infos
type PersonagemRacas struct {
	Racas     []string
	RacasInfo map[string]Raca
}

// Função para carregar dados das raças de um arquivo JSON
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
func NewPersonagemRacas(filename string) (*PersonagemRacas, error) {
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

// Função para gerar raça aleatória
func gerarRaca(racas *PersonagemRacas) (string, Raca) {
	rand.Seed(time.Now().UnixNano())
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para calcular idade
func calcularIdade(raca string, racas *PersonagemRacas) (int, string) {
	if raca == "Elfo" {
		return rand.Intn(975) + 26, "Adulto"
	}

	intervalos := racas.RacasInfo[raca].Idades
	faixasEtarias := map[string]*[2]int{
		"Jovem":  intervalos.Jovem,
		"Adulto": &intervalos.Adulto,
		"Idoso":  intervalos.Idoso,
	}

	validFaixas := []string{}
	for faixa, valores := range faixasEtarias {
		if valores != nil {
			validFaixas = append(validFaixas, faixa)
		}
	}

	escolhida := validFaixas[rand.Intn(len(validFaixas))]
	idadeRange := faixasEtarias[escolhida]
	return rand.Intn(idadeRange[1]-idadeRange[0]+1) + idadeRange[0], escolhida
}
