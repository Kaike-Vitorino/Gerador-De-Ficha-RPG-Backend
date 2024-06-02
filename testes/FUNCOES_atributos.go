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
func escolherAtributos(faixaEtaria string, atributosChave []string) map[string]int {
	// Definindo pontos disponíveis com base na faixa etária
	pontosDisponiveis := map[string]int{
		"Jovem":  15,
		"Adulto": 14,
		"Idoso":  13,
	}[faixaEtaria]

	// Lista de todos os atributos possíveis
	todosAtributos := []string{"Força", "Agilidade", "Empatia", "Inteligência"}

	// Inicializando os atributos com valores mínimos
	atributosRandomizados := make(map[string]int)
	for _, atributo := range todosAtributos {
		atributosRandomizados[atributo] = 0
	}

	// Distribuindo 2 pontos para cada atributo não-chave
	for _, atributo := range todosAtributos {
		if !contemItem(atributosChave, atributo) {
			atributosRandomizados[atributo] = 2
			pontosDisponiveis -= 2
		}
	}

	// Distribuindo pontos igualmente entre os atributos chave
	pontosPorAtributoChave := pontosDisponiveis / len(atributosChave)
	for _, atributo := range atributosChave {
		atributosRandomizados[atributo] = menorValor(pontosPorAtributoChave, 5)
		pontosDisponiveis -= atributosRandomizados[atributo]
	}

	// Distribuindo pontos restantes de forma aleatória
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < pontosDisponiveis; i++ {
		atributo := todosAtributos[rand.Intn(len(todosAtributos))]
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
