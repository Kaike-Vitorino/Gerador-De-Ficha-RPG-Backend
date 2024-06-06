package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DistribuidorDePericias gerencia a distribuição de pontos de perícia.
type DistribuidorDePericias struct {
	pericias map[string][]string
}

// NovaDistribuidorDePericias cria uma nova instância de DistribuidorDePericias.
func NovaDistribuidorDePericias(pericias map[string][]string) *DistribuidorDePericias {
	return &DistribuidorDePericias{pericias: pericias}
}

// Distribuir distribui os pontos de perícia com base na faixa etária e na classe.
func (d *DistribuidorDePericias) Distribuir(faixaEtaria, classe string) map[string]int {
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

	periciasPermitidas, ok := d.pericias[classe]
	if !ok {
		fmt.Printf("Classe %s não encontrada\n", classe)
		return nil
	}

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
