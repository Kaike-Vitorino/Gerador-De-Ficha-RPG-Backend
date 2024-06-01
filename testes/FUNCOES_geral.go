package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Personagem representa todas as informações de um personagem
type Personagem struct {
	Raca                     string
	Classe                   string
	Idade                    int
	FaixaEtaria              string
	Atributos                map[string]int
	ArtefatoMusicalEscolhido string
	ArmaEscolhida            string
	Equipamentos             []string
	Pericias                 map[string]string
	Talentos                 map[string][]string
}

// Função para gerar um personagem
func gerarPersonagem(racas *PersonagemRacas, classes *PersonagemClasses, status *PersonagemStatus, equipamentos *Equipamentos) (*Personagem, error) {
	raca, racaInfo := gerarRaca(racas)
	classe := gerarClasse(raca, racas)
	atributosChave := obterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := calcularIdade(raca, racas)
	atributos := escolherAtributos(faixaEtaria, atributosChave)

	// Aqui você pode adicionar a lógica para selecionar equipamentos e outros detalhes

	return &Personagem{
		Raca:        raca,
		Classe:      classe,
		Atributos:   atributos,
		Idade:       idade,
		FaixaEtaria: faixaEtaria,
		// Preencha os demais campos conforme necessário
	}, nil
}

// Função para gerar raça aleatória
func gerarRaca(racas *PersonagemRacas) (string, Raca) {
	rand.Seed(time.Now().UnixNano())
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para gerar classe
func gerarClasse(raca string, racas *PersonagemRacas) string {
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}

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

// Função para calcular idade
func calcularIdade(raca string, racas *PersonagemRacas) (int, string) {
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
		if !contains(atributosChave, atributo) {
			atributosRandomizados[atributo] = 2
			pontosDisponiveis -= 2
		}
	}

	// Distribuindo pontos igualmente entre os atributos chave
	pontosPorAtributoChave := pontosDisponiveis / len(atributosChave)
	for _, atributo := range atributosChave {
		atributosRandomizados[atributo] = min(pontosPorAtributoChave, 5)
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

// Função auxiliar para verificar se um slice contém um item específico
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Função auxiliar para obter o menor valor entre dois inteiros
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Função main para testar a criação do personagem
func main() {
	racas, err := NewPersonagemRacas("")
	if err != nil {
		fmt.Println("Erro ao carregar raças:", err)
		return
	}

	classes, err := NewPersonagemClasses("")
	if err != nil {
		fmt.Println("Erro ao carregar classes:", err)
		return
	}

	status, err := NewPersonagemStatus("data/atributos.json", "data/pericias.json", "data/talentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar status:", err)
		return
	}

	equipamentos, err := CarregarEquipamentos()
	if err != nil {
		fmt.Println("Erro ao carregar equipamentos:", err)
		return
	}

	personagem, err := gerarPersonagem(racas, classes, status, equipamentos)
	if err != nil {
		fmt.Println("Erro ao gerar personagem:", err)
		return
	}

	fmt.Printf("Personagem: %+v\n", personagem)
}
