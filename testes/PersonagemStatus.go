package main

import (
	"fmt"
)

type PersonagemStatus struct {
	Atributos Atributos
	Pericias  Pericias
	Talentos  Talentos
}

// NewPersonagemStatus inicializa PersonagemStatus com valores carregados dos arquivos JSON
func NewPersonagemStatus() (*PersonagemStatus, error) {
	atributos, err := carregarAtributos("")
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar atributos: %v", err)
	}

	pericias, err := carregarPericias("")
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar perícias: %v", err)
	}

	talentos, err := carregarTalentos("")
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar talentos: %v", err)
	}

	return &PersonagemStatus{
		Atributos: atributos,
		Pericias:  pericias,
		Talentos:  talentos,
	}, nil
}
