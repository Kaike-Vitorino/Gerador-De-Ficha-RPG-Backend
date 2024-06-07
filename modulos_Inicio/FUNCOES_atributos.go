package main

import (
	"math/rand"
	"time"
)

// Função para obter os atributos chave
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

// Função onde os pontos de atributo são distribuídos e o valor/nível dos atributos determinados
func escolherAtributos(faixaEtaria string, atributosChave []string, atributos Atributos) map[string]int {
	// Inicializando os atributos com valores mínimos
	atributosRandomizados := map[string]int{
		"Forca":        atributos.Forca[0],
		"Agilidade":    atributos.Agilidade[0],
		"Inteligencia": atributos.Inteligencia[0],
		"Empatia":      atributos.Empatia[0],
	}

	// Definindo pontos disponíveis com base na faixa etária
	pontosDisponiveis := map[string]int{
		"Jovem":  15,
		"Adulto": 14,
		"Idoso":  13,
	}[faixaEtaria]

	// Distribuindo 2 pontos para cada atributo
	for atributo := range atributosRandomizados {
		atributosRandomizados[atributo] = 2
		pontosDisponiveis -= 2
	}

	// Distribuindo pontos igualmente entre os atributos chave
	pontosPorAtributoChave := pontosDisponiveis / len(atributosChave)
	for _, atributo := range atributosChave {
		pontosParaAdicionar := menorValor(pontosPorAtributoChave, 5-atributosRandomizados[atributo])
		atributosRandomizados[atributo] += pontosParaAdicionar
		pontosDisponiveis -= pontosParaAdicionar
	}

	// Distribuindo pontos restantes de forma aleatória entre os atributos não-chave
	rand.Seed(time.Now().UnixNano())
	atributosNaoChave := []string{}
	for atributo := range atributosRandomizados {
		if !contemItem(atributosChave, atributo) {
			atributosNaoChave = append(atributosNaoChave, atributo)
		}
	}

	for i := 0; i < pontosDisponiveis; i++ {
		atributo := atributosNaoChave[rand.Intn(len(atributosNaoChave))]
		if atributosRandomizados[atributo] < 5 {
			atributosRandomizados[atributo]++
		}
	}
	// Adicionando +1 ponto ao atributo-chave se houver apenas um
	if len(atributosChave) == 1 {
		atributoChave := atributosChave[0]
		atributosRandomizados[atributoChave]++
	}

	return atributosRandomizados
}
