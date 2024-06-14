package personagem

import (
	"math/rand"
	"psBackKG/model/charbuilder"
)

// Função para Gerar raça aleatória
func GerarRaca(racas *charbuilder.PersonagemRacas) (string, charbuilder.Raca) {
	charbuilder.NovoGeradorAleatorio()
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para calcular idade
func CalcularIdade(raca string, racas *charbuilder.PersonagemRacas) (int, string) {
	if raca == "Elfo" {
		return rand.Intn(975) + 26, "Adulto"
	}

	intervalos := racas.RacasInfo[raca].Idades
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
