package utils

import (
	"math/rand"
	"time"
)

func novoGeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Função auxiliar para verificar se um slice contém um item específico
func contemItem(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Função auxiliar para obter o menor valor entre dois inteiros
func menorValor(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// randomSample retorna uma amostra aleatória dos elementos fornecidos.
func randomSample(elements []string, n int) []string {
	if n > len(elements) {
		n = len(elements)
	}
	if n <= 0 {
		return []string{}
	}
	novoGeradorAleatorio()
	perm := rand.Perm(len(elements))
	sample := make([]string, n)
	for i := 0; i < n; i++ {
		sample[i] = elements[perm[i]]
	}
	return sample
}

// remove removes an element from a slice.
func remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Função para verificar a igualdade de dois slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Função utilitária para selecionar um item aleatório de uma lista
func randomChoice(options []string) string {
	novoGeradorAleatorio()
	return options[rand.Intn(len(options))]
}

// randomChoiceSlice retorna um item aleatório de um slice de strings
func randomChoiceSlice(options []string) string {
	novoGeradorAleatorio()
	return options[rand.Intn(len(options))]
}

// randomChoiceFromMap retorna uma chave aleatória de um mapa
func randomChoiceFromMap(m map[string]Arma) string {
	novoGeradorAleatorio()
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

// Função utilitária para verificar se um slice contém algum dos itens fornecidos
func contemItemVariadico(slice []string, itens ...string) bool {
	for _, item := range itens {
		if contemItem(slice, item) {
			return true
		}
	}
	return false
}

// contemItem verifica se um item está presente em uma lista
func contemItem2(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
