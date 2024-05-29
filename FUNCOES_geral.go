package main

import (
	"math/rand"
	"time"
)

// Estruturas e variáveis globais
var (
	raca                     string
	artefatoMusicalEscolhido string
	armaEscolhida            string
	equipamentos             []string
)

// Função para gerar raça aleatória.
func gerarRaca(racas *PersonagemRacas) (string, Raca) {
	rand.Seed(time.Now().UnixNano())
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para gerar classe.
func gerarClasse(raca string, racas *PersonagemRacas) string {
	rand.Seed(time.Now().UnixNano())
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}

// Função para obter os atributos chave.
func obterAtributosChave(classe string, racaInfo Raca, classes map[string]Classe) []string {
	atributoChaveClasse := classes[classe].AtributoChave
	atributoChaveRaca := racaInfo.AtributoChave
	atributosChave := map[string]bool{
		atributoChaveClasse: true,
		atributoChaveRaca:   true,
	}
	result := []string{}
	for k := range atributosChave {
		result = append(result, k)
	}
	return result
}

// Função para calcular idade.
func calcularIdade(raca string, racas *PersonagemRacas) (int, string) {
	rand.Seed(time.Now().UnixNano())

	if raca == "Elfo" {
		return rand.Intn(975) + 26, "Adulto"
	}

	intervalos := racas.IdadeRacas[raca]
	faixasEtarias := map[string]*[2]int{
		"Jovem":  intervalos.Jovem,
		"Adulto": &intervalos.Adulto,
		"Idoso":  intervalos.Idoso,
	}

	validFaixas := []string{}
	for faixa, valores := range faixasEtarias {
		if valores != nil {
			validFaixas = append(validFaixas, faixa)
		}
	}

	escolhida := validFaixas[rand.Intn(len(validFaixas))]
	idadeRange := faixasEtarias[escolhida]
	return rand.Intn(idadeRange[1]-idadeRange[0]+1) + idadeRange[0], escolhida
}
