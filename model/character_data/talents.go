package character_data

import (
	"encoding/json"
	"os"
)

// Talentos todos os talentos do sistema
type Talentos struct {
	TalentosGerais map[string][]string `json:"TalentosGerais"`
}

// Talento representa um talento do character_logic
type Talento struct {
	Nome  string
	Nivel int
}

// Função para carregar dados de talentos de um arquivo JSON
func CarregarTalentos(filename string) (Talentos, error) {
	var talentos Talentos
	data, err := os.ReadFile(filename)
	if err != nil {
		return talentos, err
	}
	err = json.Unmarshal(data, &talentos)
	if err != nil {
		return talentos, err
	}
	return talentos, nil
}
