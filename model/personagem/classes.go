package personagem

import (
	"math/rand"
	"psBackKG/model/charbuilder"
)

// Função para Gerar classe
func GerarClasse(raca string, racas *charbuilder.PersonagemRacas) string {
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}
