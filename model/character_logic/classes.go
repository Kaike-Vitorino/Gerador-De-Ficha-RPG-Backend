package character_logic

import (
	"math/rand"
	DataChar "psBackKG/model/character_data"
)

// Função para Gerar classe
func GerarClasse(raca string, racas *DataChar.PersonagemRacas) string {
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}
