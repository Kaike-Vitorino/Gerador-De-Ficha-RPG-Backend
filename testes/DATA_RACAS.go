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
func carregarRacasEIdades(filename string) (map[string]Raca, error) {
	var racas map[string]Raca
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return racas, err
	}
	err = json.Unmarshal(data, &racas)
	return racas, err
}

// NewPersonagemRacas inicializa PersonagemRacas com valores carregados do arquivo JSON unificado.
func NewPersonagemRacas(classes []string, racasFile string) (*PersonagemRacas, error) {
	racas, err := carregarRacasEIdades(racasFile)
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

// Função auxiliar para imprimir as informações de uma raça de forma legível
func imprimirRaca(raca Raca) {
	fmt.Printf("Atributo Chave: %s\n", raca.AtributoChave)
	fmt.Printf("Talento Ascendente: %s\n", raca.TalentoAscendente)
	fmt.Printf("Profissões Típicas: %v\n", raca.ProfissoesTipicas)
	fmt.Print("Idades: ")
	if raca.Idades.Jovem != nil {
		fmt.Printf("Jovem: [%d, %d] ", raca.Idades.Jovem[0], raca.Idades.Jovem[1])
	}
	fmt.Printf("Adulto: [%d, %d] ", raca.Idades.Adulto[0], raca.Idades.Adulto[1])
	if raca.Idades.Idoso != nil {
		fmt.Printf("Idoso: [%d, %d] ", raca.Idades.Idoso[0], raca.Idades.Idoso[1])
	}
	fmt.Println()
}

// Funcao main para testar oq foi feito nesse modulo
/*
func main() {
	// Codigo para testar
	classes := []string{"Caçador", "Druida", "Mago", "Rider", "Guerreiro", "Ladino", "Mascate", "Bardo"}
	personagemRacas, err := NewPersonagemRacas(classes, "racas_e_idades.json")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Racas:")
	for _, raca := range personagemRacas.Racas {
		fmt.Println(raca)
	}

	fmt.Println("\nRacas Info:")
	for nome, info := range personagemRacas.RacasInfo {
		fmt.Printf("%s: \n", nome)
		imprimirRaca(info)
	}
}
*/
