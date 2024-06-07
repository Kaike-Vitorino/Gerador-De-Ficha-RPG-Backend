package main

import (
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().UnixNano())
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
