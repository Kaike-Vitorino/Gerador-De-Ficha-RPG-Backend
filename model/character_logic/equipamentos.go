package character_logic

import (
	"fmt"
	DataChar "psBackKG/model/character_data"
)

// Função para gerar uma arma aleatória para uma classe
func GerarArma(classe string, classeInfo map[string]DataChar.Classe, equipamentos *DataChar.Equipamentos) ([]string, error) {
	equipamentosResultantes := []string{}
	armasDisponiveis := classeInfo[classe].Equipamentos.Arma

	if len(armasDisponiveis) == 0 {
		return equipamentosResultantes, fmt.Errorf("Nenhuma arma disponível para a classe %s", classe)
	}

	funcs := map[string]func(){
		"Rider": func() {
			armasEscolhidas := DataChar.RandomSample(armasDisponiveis, 2)
			for ContemCombinacaoInvalida(armasEscolhidas) {
				armasEscolhidas = DataChar.RandomSample(armasDisponiveis, 2)
			}
			equipamentosResultantes = append(equipamentosResultantes, armasEscolhidas...)
		},
		"Bardo": func() {
			armaEscolhida := armasDisponiveis[0]
			artefatoMusicalEscolhido := DataChar.RandomChoiceSlice(classeInfo[classe].Equipamentos.ArtefatoMusical)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida, artefatoMusicalEscolhido)
		},
		"Guerreiro": func() {
			armaEscolhida := DataChar.RandomChoiceFromMap(equipamentos.Armas1M)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida)
		},
		"default": func() {
			armaEscolhida := DataChar.RandomChoiceSlice(armasDisponiveis)
			equipamentosResultantes = append(equipamentosResultantes, armaEscolhida)
		},
	}

	if f, ok := funcs[classe]; ok {
		f()
	} else {
		funcs["default"]()
	}

	// Verificar se precisa alterar a arma escolhida
	if len(equipamentosResultantes) == 1 {
		armaEscolhida := equipamentosResultantes[0]
		if armaEscolhida == "Arco" {
			armaEscolhida = DataChar.RandomChoiceSlice([]string{"Arco Curto", "Arco Longo"})
			equipamentosResultantes[0] = armaEscolhida
		} else if armaEscolhida == "Lança" {
			armaEscolhida = DataChar.RandomChoiceSlice([]string{"Lança Curta", "Lança Longa"})
			equipamentosResultantes[0] = armaEscolhida
		}
	} else if len(equipamentosResultantes) == 2 {
		for i, arma := range equipamentosResultantes {
			if arma == "Arco" {
				equipamentosResultantes[i] = DataChar.RandomChoiceSlice([]string{"Arco Curto", "Arco Longo"})
			} else if arma == "Lança" {
				equipamentosResultantes[i] = DataChar.RandomChoiceSlice([]string{"Lança Curta", "Lança Longa"})
			}
		}
	}

	return equipamentosResultantes, nil
}

// ContemCombinacaoInvalida verifica se a combinação de armas é inválida (Só usada no caso do Rider)
func ContemCombinacaoInvalida(armas []string) bool {
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
