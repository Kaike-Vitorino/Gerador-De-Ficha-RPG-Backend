package main

import (
	"math/rand"
)

// Talento representa um talento do personagem
type Talento struct {
	Nome  string
	Nivel int
}

// EscolherTalentos centraliza a lógica para escolher os talentos do personagem.
func EscolherTalentos(classe, raca, faixaEtaria string, racasInfo map[string]Raca, talentosClasses map[string][]string, talentosGerais map[string][]string) map[string]Talento {
	// Inicializar talentos sem nível com talento ascendente e de classe
	talentosSemNivel := RandomizarTalentoAscendente(raca, nil, racasInfo)
	talentosSemNivel = RandomizarTalentoClasse(classe, talentosSemNivel, talentosClasses)

	// Limitar a quantidade máxima de talentos
	quantidadeTalentosRestantes := 12 - len(talentosSemNivel)
	talentosEscolhidos := RandomizarTalentosGerais(faixaEtaria, classe, 1, talentosSemNivel, talentosGerais, quantidadeTalentosRestantes)

	// Converta o map[string]Talento para map[string]Talento
	resultado := make(map[string]Talento)
	for nome, talento := range talentosEscolhidos {
		resultado[nome] = talento
	}
	return resultado
}

func RandomizarTalentoAscendente(raca string, talentosSemNivel []string, racasInfo map[string]Raca) []string {
	talentoAscendente := racasInfo[raca].TalentoAscendente
	return append(talentosSemNivel, talentoAscendente)
}

func RandomizarTalentoClasse(classe string, talentosSemNivel []string, talentosClasses map[string][]string) []string {
	if talentosDisponiveis, ok := talentosClasses[classe]; ok {
		talentoEscolhido := talentosDisponiveis[rand.Intn(len(talentosDisponiveis))]
		return append(talentosSemNivel, talentoEscolhido)
	}
	return talentosSemNivel
}

// RandomizarTalentosGerais randomiza os talentos gerais de acordo com a lógica especificada e o limite de talentos.
func RandomizarTalentosGerais(faixaEtaria, classe string, nivel int, talentosSemNivel []string, talentosGerais map[string][]string, quantidadeTalentosRestantes int) map[string]Talento {
	quantidadeTalentos := DeterminarQuantidadeTalentos(faixaEtaria)
	if quantidadeTalentos > quantidadeTalentosRestantes {
		quantidadeTalentos = quantidadeTalentosRestantes
	}

	talentosEscolhidos := make(map[string]Talento)
	if rand.Intn(2) == 0 {
		return ProcessarTalentoSacrificado(quantidadeTalentos, classe, nivel, talentosSemNivel, talentosGerais)
	}

	talentosGeraisEscolhidos := randomSample(talentosGerais[classe], quantidadeTalentos)
	for _, talento := range talentosGeraisEscolhidos {
		talentosEscolhidos[talento] = Talento{Nome: talento, Nivel: nivel}
	}
	return talentosEscolhidos
}

// ProcessarTalentoSacrificado processa a lógica de talentos sacrificados.
func ProcessarTalentoSacrificado(quantidadeTalentos int, classe string, nivel int, talentosSemNivel []string, talentosGerais map[string][]string) map[string]Talento {
	talentosEscolhidos := make(map[string]Talento)
	if quantidadeTalentos-1 > 0 {
		talentosAtuaisEscolhidos := randomSample(talentosGerais[classe], quantidadeTalentos-1)
		talentosSemNivel = append(talentosSemNivel, talentosAtuaisEscolhidos...)

		talentoNivel2 := talentosSemNivel[rand.Intn(len(talentosSemNivel))]
		talentosSemNivel = remove(talentosSemNivel, talentoNivel2)

		talentosEscolhidos[talentoNivel2] = Talento{Nome: talentoNivel2, Nivel: nivel + 1}
		for _, talento := range talentosSemNivel {
			talentosEscolhidos[talento] = Talento{Nome: talento, Nivel: nivel}
		}
	}
	return talentosEscolhidos
}

func DeterminarQuantidadeTalentos(faixaEtaria string) int {
	talentosPorFaixaEtaria := map[string]int{
		"Jovem":  1,
		"Adulto": 2,
		"Idoso":  3,
	}

	if quantidade, ok := talentosPorFaixaEtaria[faixaEtaria]; ok {
		return quantidade
	}
	return 0
}
