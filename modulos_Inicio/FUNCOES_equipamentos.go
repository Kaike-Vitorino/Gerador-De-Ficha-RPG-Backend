package main

import (
	"fmt"
)

// GerarArma gera a arma do personagem com base na classe
func GerarArma(classe string, classeInfo map[string]Classe, equipamentos *Equipamentos) ([]string, error) {
	equipamentosResultantes := []string{}
	armasDisponiveis := classeInfo[classe].Equipamentos.Arma

	if len(armasDisponiveis) == 0 {
		return equipamentosResultantes, fmt.Errorf("Nenhuma arma disponível para a classe %s", classe)
	}

	funcs := map[string]func(){
		"Rider": func() {
			armasEscolhidas := randomSample(armasDisponiveis, 2)
			for contemCombinacaoInvalida(armasEscolhidas) {
				armasEscolhidas = randomSample(armasDisponiveis, 2)
			}
			equipamentosResultantes = append(equipamentosResultantes, armasEscolhidas...)
		},
		"Bardo": func() {
			armaEscolhida := armasDisponiveis[0]
			artefatoMusicalEscolhido := randomChoiceSlice(classeInfo[classe].Equipamentos.ArtefatoMusical)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida, artefatoMusicalEscolhido)
		},
		"Guerreiro": func() {
			armaEscolhida := randomChoiceFromMap(equipamentos.Armas1M)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida)
		},
		"default": func() {
			armaEscolhida := randomChoiceSlice(armasDisponiveis)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida)
		},
	}

	if f, ok := funcs[classe]; ok {
		f()
	} else {
		funcs["default"]()
	}

	return equipamentosResultantes, nil
}

// contemCombinacaoInvalida verifica se a combinação de armas é inválida (So usada no caso do Rider)
func contemCombinacaoInvalida(armas []string) bool {
	invalidas := [][]string{
		{"Lança", "Machadinha"},
		{"Arco Curto", "Funda"},
	}

	for _, invalida := range invalidas {
		if (armas[0] == invalida[0] && armas[1] == invalida[1]) || (armas[0] == invalida[1] && armas[1] == invalida[0]) {
			return true
		}
	}

	return false
}
