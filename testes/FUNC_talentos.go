package main

import (
	"math/rand"
)

// RandomizarTalentoAscendente adiciona o talento ascendente da raça à lista de talentos sem nível.
func RandomizarTalentoAscendente(raca string, talentosSemNivel []string, racasInfo map[string]Raca) []string {
	talentoAscendente := racasInfo[raca].TalentoAscendente
	return append(talentosSemNivel, talentoAscendente)
}

// RandomizarTalentoClasse adiciona um talento de classe à lista de talentos sem nível.
func RandomizarTalentoClasse(classe string, talentosSemNivel []string, talentosClasses map[string][]string) []string {
	if talentosDisponiveis, ok := talentosClasses[classe]; ok {
		talentoEscolhido := talentosDisponiveis[rand.Intn(len(talentosDisponiveis))]
		return append(talentosSemNivel, talentoEscolhido)
	}
	return talentosSemNivel
}

// DeterminarQuantidadeTalentos determina a quantidade de talentos com base na faixa etária.
func DeterminarQuantidadeTalentos(faixaEtaria string) int {
	switch faixaEtaria {
	case "Jovem":
		return 1
	case "Adulto":
		return 2
	case "Idoso":
		return 3
	default:
		return 0
	}
}

// ProcessarTalentoSacrificado processa a lógica de talentos sacrificados.
func ProcessarTalentoSacrificado(quantidadeTalentos int, classe string, nivel int, talentosSemNivel []string, talentosGerais map[string][]string) map[string]map[string]int {
	talentosEscolhidos := make(map[string]map[string]int)
	if quantidadeTalentos-1 > 0 {
		talentosAtuaisEscolhidos := randomSample(talentosGerais[classe], quantidadeTalentos-1)
		talentosSemNivel = append(talentosSemNivel, talentosAtuaisEscolhidos...)

		talentoNivel2 := talentosSemNivel[rand.Intn(len(talentosSemNivel))]
		talentosSemNivel = remove(talentosSemNivel, talentoNivel2)

		talentosEscolhidos[talentoNivel2] = map[string]int{"Nivel": nivel + 1}
		for _, talento := range talentosSemNivel {
			talentosEscolhidos[talento] = map[string]int{"Nivel": nivel}
		}
	}
	return talentosEscolhidos
}

// RandomizarTalentosGerais randomiza os talentos gerais de acordo com a lógica especificada.
func RandomizarTalentosGerais(faixaEtaria, classe string, nivel int, talentosSemNivel []string, talentosGerais map[string][]string) map[string]map[string]int {
	quantidadeTalentos := DeterminarQuantidadeTalentos(faixaEtaria)
	talentosEscolhidos := make(map[string]map[string]int)

	if rand.Intn(2) == 0 {
		return ProcessarTalentoSacrificado(quantidadeTalentos, classe, nivel, talentosSemNivel, talentosGerais)
	}

	talentosGeraisEscolhidos := randomSample(talentosGerais[classe], quantidadeTalentos)
	for _, talento := range talentosGeraisEscolhidos {
		talentosEscolhidos[talento] = map[string]int{"Nivel": nivel}
	}
	return talentosEscolhidos
}
