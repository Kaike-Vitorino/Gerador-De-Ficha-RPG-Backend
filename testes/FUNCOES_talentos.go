package main

import (
	"math/rand"
)

// Talento representa um talento do personagem
type Talento struct {
	Nome  string
	Nivel int
}

// DistribuidorDeTalentos gerencia a distribuição de talentos
type DistribuidorDeTalentos struct {
	talentosClasses map[string][]string
	talentosGerais  map[string][]string
	racasInfo       map[string]Raca
}

// NovaDistribuidorDeTalentos cria uma nova instância de DistribuidorDeTalentos
func NovaDistribuidorDeTalentos(talentosClasses map[string][]string, talentosGerais map[string][]string, racasInfo map[string]Raca) *DistribuidorDeTalentos {
	return &DistribuidorDeTalentos{
		TalentosClasses: talentosClasses,
		TalentosGerais:  talentosGerais,
		RacasInfo:       racasInfo,
	}
}

// EscolherTalentos centraliza a lógica para escolher os talentos do personagem.
func (d *DistribuidorDeTalentos) EscolherTalentos(classe, raca, faixaEtaria string) map[string]Talento {
	talentosSemNivel := d.randomizarTalentoAscendente(raca, nil)
	talentosSemNivel = d.randomizarTalentoClasse(classe, talentosSemNivel)
	talentosEscolhidos := d.randomizarTalentosGerais(faixaEtaria, classe, 1, talentosSemNivel)

	// Converta o map[string]Talento para map[string]Talento
	resultado := make(map[string]Talento)
	for nome, talento := range talentosEscolhidos {
		resultado[nome] = talento
	}
	return resultado
}

func (d *DistribuidorDeTalentos) randomizarTalentoAscendente(raca string, talentosSemNivel []string) []string {
	talentoAscendente := d.racasInfo[raca].TalentoAscendente
	return append(talentosSemNivel, talentoAscendente)
}

func (d *DistribuidorDeTalentos) randomizarTalentoClasse(classe string, talentosSemNivel []string) []string {
	if talentosDisponiveis, ok := d.talentosClasses[classe]; ok {
		talentoEscolhido := talentosDisponiveis[rand.Intn(len(talentosDisponiveis))]
		return append(talentosSemNivel, talentoEscolhido)
	}
	return talentosSemNivel
}

// RandomizarTalentosGerais randomiza os talentos gerais de acordo com a lógica especificada.
func (d *DistribuidorDeTalentos) randomizarTalentosGerais(faixaEtaria, classe string, nivel int, talentosSemNivel []string) map[string]Talento {
	quantidadeTalentos := determinarQuantidadeTalentos(faixaEtaria)
	talentosEscolhidos := make(map[string]Talento)

	if rand.Intn(2) == 0 {
		return d.ProcessarTalentoSacrificado(quantidadeTalentos, classe, nivel, talentosSemNivel)
	}

	talentosGeraisEscolhidos := randomSample(d.talentosGerais[classe], quantidadeTalentos)
	for _, talento := range talentosGeraisEscolhidos {
		talentosEscolhidos[talento] = Talento{Nome: talento, Nivel: nivel}
	}
	return talentosEscolhidos
}

// ProcessarTalentoSacrificado processa a lógica de talentos sacrificados.
func (d *DistribuidorDeTalentos) ProcessarTalentoSacrificado(quantidadeTalentos int, classe string, nivel int, talentosSemNivel []string) map[string]Talento {
	talentosEscolhidos := make(map[string]Talento)
	if quantidadeTalentos-1 > 0 {
		talentosAtuaisEscolhidos := randomSample(d.talentosGerais[classe], quantidadeTalentos-1)
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

func determinarQuantidadeTalentos(faixaEtaria string) int {
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
