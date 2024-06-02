package main

import (
	"math/rand"
)

// Talento representa um talento com seu nível.
type Talento struct {
	Nome  string
	Nivel int
}

// TalentosManager gerencia a lógica de talentos.
type TalentosManager struct {
	RacasInfo       map[string]Raca
	TalentosClasses map[string][]string
	TalentosGerais  map[string][]string
}

// NewTalentosManager cria uma nova instância de TalentosManager.
func NewTalentosManager(racasInfo map[string]Raca, talentosClasses, talentosGerais map[string][]string) *TalentosManager {
	return &TalentosManager{
		RacasInfo:       racasInfo,
		TalentosClasses: talentosClasses,
		TalentosGerais:  talentosGerais,
	}
}

// RandomizarTalentoAscendente adiciona o talento ascendente da raça à lista de talentos sem nível.
func (tm *TalentosManager) RandomizarTalentoAscendente(raca string, talentosSemNivel []string) []string {
	talentoAscendente := tm.RacasInfo[raca].TalentoAscendente
	return append(talentosSemNivel, talentoAscendente)
}

// RandomizarTalentoClasse adiciona um talento de classe à lista de talentos sem nível.
func (tm *TalentosManager) RandomizarTalentoClasse(classe string, talentosSemNivel []string) []string {
	if talentosDisponiveis, ok := tm.TalentosClasses[classe]; ok {
		talentoEscolhido := talentosDisponiveis[rand.Intn(len(talentosDisponiveis))]
		return append(talentosSemNivel, talentoEscolhido)
	}
	return talentosSemNivel
}

// DeterminarQuantidadeTalentos determina a quantidade de talentos com base na faixa etária usando um mapa.
func (tm *TalentosManager) DeterminarQuantidadeTalentos(faixaEtaria string) int {
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

// ProcessarTalentoSacrificado processa a lógica de talentos sacrificados.
func (tm *TalentosManager) ProcessarTalentoSacrificado(quantidadeTalentos int, classe string, nivel int, talentosSemNivel []string) []Talento {
	talentosEscolhidos := []Talento{}
	if quantidadeTalentos-1 > 0 {
		talentosAtuaisEscolhidos := randomSample(tm.TalentosGerais[classe], quantidadeTalentos-1)
		talentosSemNivel = append(talentosSemNivel, talentosAtuaisEscolhidos...)

		talentoNivel2 := talentosSemNivel[rand.Intn(len(talentosSemNivel))]
		talentosSemNivel = remove(talentosSemNivel, talentoNivel2)

		talentosEscolhidos = append(talentosEscolhidos, Talento{Nome: talentoNivel2, Nivel: nivel + 1})
		for _, talento := range talentosSemNivel {
			talentosEscolhidos = append(talentosEscolhidos, Talento{Nome: talento, Nivel: nivel})
		}
	}
	return talentosEscolhidos
}

// RandomizarTalentosGerais randomiza os talentos gerais de acordo com a lógica especificada.
func (tm *TalentosManager) RandomizarTalentosGerais(faixaEtaria, classe string, nivel int, talentosSemNivel []string) []Talento {
	quantidadeTalentos := tm.DeterminarQuantidadeTalentos(faixaEtaria)
	talentosEscolhidos := []Talento{}

	if rand.Intn(2) == 0 {
		return tm.ProcessarTalentoSacrificado(quantidadeTalentos, classe, nivel, talentosSemNivel)
	}

	talentosGeraisEscolhidos := randomSample(tm.TalentosGerais[classe], quantidadeTalentos)
	for _, talento := range talentosGeraisEscolhidos {
		talentosEscolhidos = append(talentosEscolhidos, Talento{Nome: talento, Nivel: nivel})
	}
	return talentosEscolhidos
}

// EscolherTalentos centraliza a lógica para escolher os talentos do personagem.
func (tm *TalentosManager) EscolherTalentos(classe, raca, faixaEtaria string) []Talento {
	talentosSemNivel := tm.RandomizarTalentoAscendente(raca, nil)
	talentosSemNivel = tm.RandomizarTalentoClasse(classe, talentosSemNivel)
	talentosEscolhidos := tm.RandomizarTalentosGerais(faixaEtaria, classe, 1, talentosSemNivel)
	return talentosEscolhidos
}
