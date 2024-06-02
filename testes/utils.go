package main

import "math/rand"

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

// randomSample returns a random sample of n elements from a slice.
func randomSample(slice []string, n int) []string {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
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
