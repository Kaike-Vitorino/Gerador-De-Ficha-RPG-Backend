package charbuilder

import (
	"fmt"
	"math/rand"
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

// CarregarEquipamentos carrega os ficha do JSON
func CarregarEquipamentos() (*Equipamentos, error) {
	equipamentos := &Equipamentos{}

	err := ReadJSON("data/itensComercio.json", &equipamentos.ItensComercio)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar itens de comercio: %v", err)
	}

	err = ReadJSON("data/equipamentos.json", &equipamentos)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar ficha: %v", err)
	}

	// Merge de armas
	equipamentos.ListaArmas = MergeArmas(equipamentos.Armas1M, equipamentos.Armas2M)
	equipamentos.ListaArmasADistancia = MergeArmas(equipamentos.ArmasDistancia1M, equipamentos.ArmasDistancia2M)
	equipamentos.ListaArmasFinal = MergeArmas(equipamentos.ListaArmasADistancia, equipamentos.ListaArmas)

	return equipamentos, nil
}

// Função para escolher um artefato musical para a classe Bardo
func EscolherArtefatoMusical(classes *PersonagemClasses, classe string) string {
	artefatosMusicaisDisponiveis := classes.ClasseInfo[classe].Equipamentos.ArtefatoMusical
	if len(artefatosMusicaisDisponiveis) == 0 {
		return ""
	}
	indiceAleatorio := rand.Intn(len(artefatosMusicaisDisponiveis))
	return artefatosMusicaisDisponiveis[indiceAleatorio]
}

// MergeArmas junta dois mapas de arma
func MergeArmas(map1, map2 map[string]Arma) map[string]Arma {
	merged := make(map[string]Arma)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}
