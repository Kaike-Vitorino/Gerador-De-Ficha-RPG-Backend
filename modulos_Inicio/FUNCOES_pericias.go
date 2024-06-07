package main

import (
	"fmt"
	"math/rand"
)

// Função para distribuir os pontos de Pericia com base na faixa etária e na classe.
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

	// Obter as Pericias permitidas para a classe
	classeInfo, ok := classesInfo[classe]
	if !ok {
		fmt.Printf("Classe %s não encontrada ou sem Pericias permitidas\n", classe)
		return nil
	}

	periciasPermitidas := classeInfo.Pericias

	periciasDistribuidas := make(map[string]int)
	for _, pericia := range periciasPermitidas {
		periciasDistribuidas[pericia] = 0
	}

	novoGeradorAleatorio()
	for pontosDisponiveis > 0 {
		pericia := periciasPermitidas[rand.Intn(len(periciasPermitidas))]
		if periciasDistribuidas[pericia] < 3 {
			periciasDistribuidas[pericia]++
			pontosDisponiveis--
		}
	}

	return periciasDistribuidas
}
