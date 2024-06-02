package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

// Classe representa as informacoes de uma classe
type Classe struct {
	AtributoChave string            `json:"atributo_chave"`
	Pericias      []string          `json:"pericias"`
	Equipamentos  EquipamentoClasse `json:"equipamentos"`
	DadosRecurso  map[string]string `json:"dados_recurso"`
}

// EquipamentoClasse representa os equipamentos de uma classe
type EquipamentoClasse struct {
	Arma            []string `json:"arma"`
	Armadura        *string  `json:"armadura"`
	Itens           int      `json:"itens"`
	ArtefatoMusical []string `json:"artefato_musical"`
	Cavalo          int      `json:"cavalo"`
}

// PersonagemClasses armazena todas as classes e as infos dela
type PersonagemClasses struct {
	Classes         []string
	ClasseInfo      map[string]Classe
	TalentosClasses map[string][]string
}

// Função para carregar dados de classes de um arquivo JSON
func carregarClasses(filename string) (map[string]Classe, map[string][]string, error) {
	var data struct {
		Classes         map[string]Classe   `json:"Classes"`
		TalentosClasses map[string][]string `json:"TalentosClasses"`
	}
	err := readJSON(filename, &data)
	return data.Classes, data.TalentosClasses, err
}

// Função para ler dados JSON de um arquivo
func readJSON(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// NewPersonagemClasses inicializa PersonagemClasses com valores carregados do arquivo JSON
func NewPersonagemClasses(filename string) (*PersonagemClasses, error) {
	if filename == "" {
		filename = "data/classes.json"
	}

	classes, talentosClasses, err := carregarClasses(filename)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar classes: %v", err)
	}

	classNames := make([]string, 0, len(classes))
	for name := range classes {
		classNames = append(classNames, name)
	}

	return &PersonagemClasses{
		Classes:         classNames,
		ClasseInfo:      classes,
		TalentosClasses: talentosClasses,
	}, nil
}

// Função para gerar classe
func gerarClasse(raca string, racas *PersonagemRacas) string {
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}