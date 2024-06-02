package main

import (
	"math/rand"
	"time"
)

// DistribuiPontosPericia distribui os pontos de perícia com base na faixa etária e classe do personagem
func DistribuiPontosPericia(faixaEtaria, classe string, classes *PersonagemClasses) map[string]int {
	// Definindo pontos disponíveis com base na faixa etária
	pontosDisponiveis := map[string]int{
		"Jovem":  8,
		"Adulto": 10,
		"Idoso":  12,
	}[faixaEtaria]

	// Obter as perícias permitidas para a classe
	periciasPermitidas := classes.ClasseInfo[classe].Pericias
	periciasDistribuidas := make(map[string]int)
	for _, pericia := range periciasPermitidas {
		periciasDistribuidas[pericia] = 0
	}

	// Distribuir os pontos aleatoriamente
	rand.Seed(time.Now().UnixNano())
	for pontosDisponiveis > 0 {
		pericia := periciasPermitidas[rand.Intn(len(periciasPermitidas))]
		if periciasDistribuidas[pericia] < 3 {
			periciasDistribuidas[pericia]++
			pontosDisponiveis--
		}
	}

	return periciasDistribuidas
}
