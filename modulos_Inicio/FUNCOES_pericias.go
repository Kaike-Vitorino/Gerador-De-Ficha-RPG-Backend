package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função para distribuir os pontos de perícia com base na faixa etária e na classe.
func DistribuirPericias(faixaEtaria, classe string, pericias map[string]string, classesInfo map[string]Classe) map[string]int {
	pontosPorFaixaEtaria := map[string]int{
		"Jovem":  8,
		"Adulto": 10,
		"Idoso":  12,
	}

	pontosDisponiveis, ok := pontosPorFaixaEtaria[faixaEtaria]
	if !ok {
		fmt.Printf("Faixa etária %s inválida\n", faixaEtaria)
		return nil
	}

	// Obter as perícias permitidas para a classe
	classeInfo, ok := classesInfo[classe]
	if !ok {
		fmt.Printf("Classe %s não encontrada ou sem perícias permitidas\n", classe)
		return nil
	}

	periciasPermitidas := classeInfo.Pericias

	periciasDistribuidas := make(map[string]int)
	for _, pericia := range periciasPermitidas {
		periciasDistribuidas[pericia] = 0
	}

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
