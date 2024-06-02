package main

import "fmt"

// Personagem representa todas as informações de um personagem
type Personagem struct {
	Raca                     string
	Classe                   string
	Idade                    int
	FaixaEtaria              string
	AtributosChave           []string
	Atributos                map[string]int
	ArtefatoMusicalEscolhido string
	ArmaEscolhida            string
	Equipamentos             []string
	Pericias                 map[string]int
	Talentos                 map[string][]string
}

// Função para gerar um personagem
func gerarPersonagem(racas *PersonagemRacas, classes *PersonagemClasses, status *PersonagemStatus, equipamentos *Equipamentos) (*Personagem, error) {
	raca, racaInfo := gerarRaca(racas)
	classe := gerarClasse(raca, racas)
	atributosChave := obterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := calcularIdade(raca, racas)
	atributos := escolherAtributos(faixaEtaria, atributosChave)
	pericias := DistribuiPontosPericia(faixaEtaria, classe, classes)

	// Aqui você pode adicionar a lógica para selecionar equipamentos e outros detalhes

	return &Personagem{
		Raca:           raca,
		Classe:         classe,
		AtributosChave: atributosChave,
		Atributos:      atributos,
		Idade:          idade,
		FaixaEtaria:    faixaEtaria,
		Pericias:       pericias,
		// Preencha os demais campos conforme necessário
	}, nil
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
