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

// Função para obter a faixa etária com base na idade
func ObterFaixaEtaria(raca string, idade int, racas *DataChar.PersonagemRacas) string {
	idades := racas.RacasInfo[raca].Idades
	faixas := map[string][2]int{
		"Jovem":  *idades.Jovem,
		"Adulto": idades.Adulto,
		"Idoso":  *idades.Idoso,
	}
	for faixa, intervalo := range faixas {
		if idade >= intervalo[0] && idade <= intervalo[1] {
			return faixa
		}
	}
	return "Indefinido" // Valor de fallback caso não encontre uma faixa etária correspondente
}
