package character_data

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

// Função para ler dados JSON de um arquivo
func ReadJSON(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func NovoGeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Função auxiliar para verificar se um slice contém um item específico
func ContemItem(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Função auxiliar para obter o menor valor entre dois inteiros
func MenorValor(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// RandomSample retorna uma amostra aleatória dos elementos fornecidos.
func RandomSample(elements []string, n int) []string {
	if n > len(elements) {
		n = len(elements)
	}
	if n <= 0 {
		return []string{}
	}
	NovoGeradorAleatorio()
	perm := rand.Perm(len(elements))
	sample := make([]string, n)
	for i := 0; i < n; i++ {
		sample[i] = elements[perm[i]]
	}
	return sample
}

// Remove removes an element from a slice.
func Remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Função para verificar a igualdade de dois slices
func EqualSlices(a, b []string) bool {
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
func RandomChoice(options []string) string {
	NovoGeradorAleatorio()
	return options[rand.Intn(len(options))]
}

// RandomChoiceSlice retorna um item aleatório de um slice de strings
func RandomChoiceSlice(options []string) string {
	NovoGeradorAleatorio()
	return options[rand.Intn(len(options))]
}

// RandomChoiceFromMap retorna uma chave aleatória de um mapa
func RandomChoiceFromMap(m map[string]Arma) string {
	NovoGeradorAleatorio()
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

// Função utilitária para verificar se um slice contém algum dos itens fornecidos
func ContemItemVariadico(slice []string, itens ...string) bool {
	for _, item := range itens {
		if ContemItem(slice, item) {
			return true
		}
	}
	return false
}

// ContemItem2 verifica se um item está presente em uma lista
func ContemItem2(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
