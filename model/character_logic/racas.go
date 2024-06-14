package character_logic

import (
	"math/rand"
	DataChar "psBackKG/model/character_data"
)

// Função para Gerar raça aleatória
func GerarRaca(racas *DataChar.PersonagemRacas) (string, DataChar.Raca) {
	DataChar.NovoGeradorAleatorio()
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para calcular idade
func CalcularIdade(raca string, racas *DataChar.PersonagemRacas) (int, string) {
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
