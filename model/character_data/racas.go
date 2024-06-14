package character_data

import (
	"encoding/json"
	"fmt"
	"os"
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
	data, err := os.ReadFile(filename)
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
