package main

import (
	"fmt"
)

// Arma representa as caracteristicas de uma arma
type Arma struct {
	Bonus string `json:"Bonus"`
	Dano  string `json:"Dano"`
}

// Escudo representa as caracteristicas de um escudo
type Escudo struct {
	Bonus string `json:"Bonus"`
}

// Armadura representa as caracteristicas de uma armadura
type Armadura struct {
	ValorDeArmadura string `json:"ValorDeArmadura"`
	ParteDoCorpo    string `json:"ParteDoCorpo"`
}

// Equipamentos detem todas as opções de equipamento do personagem
type Equipamentos struct {
	Items                    []string            `json:"Items"`
	ArmaEscolhida            string              `json:"ArmaEscolhida"`
	ArtefatoMusicalEscolhido string              `json:"ArtefatoMusicalEscolhido"`
	ItensComercio            []string            `json:"ItensComercio"`
	Armas1M                  map[string]Arma     `json:"Armas1M"`
	Armas2M                  map[string]Arma     `json:"Armas2M"`
	ArmasDistancia1M         map[string]Arma     `json:"ArmasDistancia1M"`
	ArmasDistancia2M         map[string]Arma     `json:"ArmasDistancia2M"`
	ListaEscudos             map[string]Escudo   `json:"Escudos"`
	ListaArmaduras           map[string]Armadura `json:"Armaduras"`
	ListaArmas               map[string]Arma     `json:"-"`
	ListaArmasADistancia     map[string]Arma     `json:"-"`
	ListaArmasFinal          map[string]Arma     `json:"-"`
}

// NewEquipamentos inicializa a struct Equipamentos com os dados do json
func NewEquipamentos() (*Equipamentos, error) {
	equipamentos := &Equipamentos{}

	err := readJSON("data/itensComercio.json", &equipamentos.ItensComercio)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar itens de comercio: %v", err)
	}

	err = readJSON("data/equipamentos.json", &equipamentos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar equipamentos: %v", err)
	}

	// Merge de armas
	equipamentos.ListaArmas = mergeArmas(equipamentos.Armas1M, equipamentos.Armas2M)
	equipamentos.ListaArmasADistancia = mergeArmas(equipamentos.ArmasDistancia1M, equipamentos.ArmasDistancia2M)
	equipamentos.ListaArmasFinal = mergeArmas(equipamentos.ListaArmasADistancia, equipamentos.ListaArmas)

	return equipamentos, nil
}

// mergeArmas junta dois mapas de arma
func mergeArmas(map1, map2 map[string]Arma) map[string]Arma {
	merged := make(map[string]Arma)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}
