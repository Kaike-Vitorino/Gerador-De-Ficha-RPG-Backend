package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estruturas e variáveis globais
var (
	raca                     string
	artefatoMusicalEscolhido string
	armaEscolhida            string
	equipamentos             []string
)

// Função para gerar raça aleatória.
func gerarRaca(racas *PersonagemRacas) (string, Raca) {
	rand.Seed(time.Now().UnixNano())
	racaAleatoria := racas.Racas[rand.Intn(len(racas.Racas))]
	racaInfo := racas.RacasInfo[racaAleatoria]
	return racaAleatoria, racaInfo
}

// Função para gerar classe.
func gerarClasse(raca string, racas *PersonagemRacas) string {
	rand.Seed(time.Now().UnixNano())
	profissoes := racas.RacasInfo[raca].ProfissoesTipicas
	return profissoes[rand.Intn(len(profissoes))]
}

// Função para obter os atributos chave.
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

// Função para calcular idade.
func calcularIdade(raca string, racas *PersonagemRacas) (int, string) {
	rand.Seed(time.Now().UnixNano())

	if raca == "Elfo" {
		return rand.Intn(975) + 26, "Adulto"
	}

	intervalos := racas.IdadeRacas[raca]
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

// Função onde os pontos de atributo são distribuídos e o valor/nível dos atributos determinados.
func escolherAtributos(faixaEtaria string, atributosChave []string) map[string]int {
	// Definindo pontos disponíveis com base na faixa etária
	pontosDisponiveis := map[string]int{"Jovem": 15,
		"Adulto": 14,
		"Idoso":  13}[faixaEtaria]

	// Lista de todos os atributos possíveis
	todosAtributos := []string{"Força", "Agilidade", "Empatia", "Inteligência"}

	// Inicializando os atributos com valores mínimos e máximos possíveis
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

func main() {
	// Inicializando dados de classes e raças
	classes := []string{"Caçador", "Druida", "Mago", "Rider", "Guerreiro", "Ladino", "Mascate", "Bardo"}
	racas := NewPersonagemRacas(classes)
	classeInfo := map[string]Classe{
		"Caçador":   {AtributoChave: "Agilidade", Pericias: []string{"Furtividade", "Movimentação", "Pontaria", "Patrulha", "Sobrevivência"}, Equipamentos: Equipamento{Arma: []string{"Arco", "Funda"}, Armadura: nil, Itens: 2, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D8", "Flechas": "D10", "Prata": "D6"}},
		"Druida":    {AtributoChave: "Inteligência", Pericias: []string{"Resiliência", "Sobrevivência", "Discernimento", "Cura", "Adestramento"}, Equipamentos: Equipamento{Arma: []string{"Bastão", "Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D8", "Prata": "D6"}},
		"Mago":      {AtributoChave: "Inteligência", Pericias: []string{"Artesanato", "Artimanha", "Tradição", "Discernimento", "Manipulação"}, Equipamentos: Equipamento{Arma: []string{"Bastão", "Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D6", "Água": "D8", "Prata": "D8"}},
		"Rider":     {AtributoChave: "Agilidade", Pericias: []string{"Resiliência", "Luta", "Pontaria", "Sobrevivência", "Adestramento"}, Equipamentos: Equipamento{Arma: []string{"Lança Curta", "Machadinha", "Arco Curto", "Funda"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 1}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D8", "Flechas": "D10", "Prata": "D6"}},
		"Guerreiro": {AtributoChave: "Força", Pericias: []string{"Potência", "Resiliência", "Luta", "Movimentação"}, Equipamentos: Equipamento{Arma: []string{"Faca", "Adaga", "Facão", "Espada Curta", "Espada Larga", "Espada Longa", "Cimitarra", "Machadinha", "Machado de Batalha", "Maça", "Maça Estrela", "Martelo de Guerra", "Mangual", "Porrete", "Lança Curta", "Tridente"}, Armadura: StringPtr("Couro"), Itens: 1, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D6", "Prata": "D6"}},
		"Ladino":    {AtributoChave: "Agilidade", Pericias: []string{"Luta", "Furtividade", "Artimanha", "Movimentação", "Manipulação"}, Equipamentos: Equipamento{Arma: []string{"Adaga"}, Armadura: nil, Itens: 2, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D6", "Água": "D6", "Prata": "D10"}},
		"Mascate":   {AtributoChave: "Empatia", Pericias: []string{"Artesanato", "Artimanha", "Patrulha", "Discernimento", "Manipulação"}, Equipamentos: Equipamento{Arma: []string{"Faca"}, Armadura: nil, Itens: 3, ArtefatoMusical: nil, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D8", "Prata": "D12"}},
		"Bardo":     {AtributoChave: "Empatia", Pericias: []string{"Tradição", "Discernimento", "Manipulação", "Atuação", "Cura"}, Equipamentos: Equipamento{Arma: []string{"Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: []string{"Lira", "Flauta"}, Cavalo: 0}, DadosRecurso: map[string]string{"Comida": "D8", "Água": "D6", "Prata": "D8"}},
	}

	// Testando as funções
	racaAleatoria, racaInfo := gerarRaca(racas)
	fmt.Printf("Raça: %s\n", racaAleatoria)
	fmt.Printf("Informações da Raça: %+v\n", racaInfo)

	classe := gerarClasse(racaAleatoria, racas)
	fmt.Printf("Classe: %s\n", classe)

	atributosChave := obterAtributosChave(classe, racaInfo, classeInfo)
	fmt.Printf("Atributos Chave: %v\n", atributosChave)

	idade, faixaEtaria := calcularIdade(racaAleatoria, racas)
	fmt.Printf("Idade: %d, Faixa Etária: %s\n", idade, faixaEtaria)

	atributos := escolherAtributos(faixaEtaria, atributosChave)
	fmt.Printf("Atributos: %v\n", atributos)
}
