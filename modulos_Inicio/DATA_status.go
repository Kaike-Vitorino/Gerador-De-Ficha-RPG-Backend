package main

import (
	"fmt"
)

// PersonagemStatus representa os stats geral do personagem, ou seja, atributos, habilidades e talentos
type PersonagemStatus struct {
	Atributos Atributos
	Pericias  Pericias
	Talentos  Talentos
}

// Função para carregar dados de atributos, Pericias e talentos de arquivos JSON
func NewPersonagemStatus(atributosFile, periciasFile, talentosFile string) (*PersonagemStatus, error) {
	status := &PersonagemStatus{}

	err := readJSON(atributosFile, &status.Atributos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar atributos: %v", err)
	}

	err = readJSON(periciasFile, &status.Pericias)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar Pericias: %v", err)
	}

	err = readJSON(talentosFile, &status.Talentos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar talentos: %v", err)
	}

	return status, nil
}
