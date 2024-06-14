package character_data

import (
	"fmt"
)

// PersonagemStatus representa os stats geral do character_logic, ou seja, atributos, habilidades e talentos
type PersonagemStatus struct {
	Atributos Atributos
	Pericias  Pericias
	Talentos  Talentos
}

// Função para carregar dados de atributos, Pericias e talentos de arquivos JSON
func NewPersonagemStatus(atributosFile, periciasFile, talentosFile string) (*PersonagemStatus, error) {
	status := &PersonagemStatus{}

	err := ReadJSON(atributosFile, &status.Atributos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar atributos: %v", err)
	}

	err = ReadJSON(periciasFile, &status.Pericias)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar Pericias: %v", err)
	}

	err = ReadJSON(talentosFile, &status.Talentos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar talentos: %v", err)
	}

	return status, nil
}
