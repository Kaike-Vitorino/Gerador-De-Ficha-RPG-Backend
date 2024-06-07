package main

import (
	"fmt"
)

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
	Talentos                 map[string]Talento
}

// Função para gerar um personagem
func gerarPersonagem(racas *PersonagemRacas, classes *PersonagemClasses, status *PersonagemStatus, equipamentos *Equipamentos, talentos *Talentos) (*Personagem, error) {
	raca, racaInfo := gerarRaca(racas)
	classe := gerarClasse(raca, racas)
	atributosChave := obterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := calcularIdade(raca, racas)
	atributos := escolherAtributos(faixaEtaria, atributosChave)

	// Distribuir perícias
	pericias := DistribuirPericias(faixaEtaria, classe, status.Pericias, classes.ClasseInfo)

	// Distribuir talentos
	talentosDistribuidos := EscolherTalentos(classe, raca, faixaEtaria, racas.RacasInfo, classes.TalentosClasses, talentos.TalentosGerais)

	// Aqui você pode adicionar a lógica para selecionar equipamentos e outros detalhes

	return &Personagem{
		Raca:           raca,
		Classe:         classe,
		AtributosChave: atributosChave,
		Atributos:      atributos,
		Idade:          idade,
		FaixaEtaria:    faixaEtaria,
		Pericias:       pericias,
		Talentos:       talentosDistribuidos,
		// Preencha os demais campos conforme necessário
	}, nil
}

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

	talentos, err := carregarTalentos("data/talentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar talentos:", err)
		return
	}

	personagem, err := gerarPersonagem(racas, classes, status, equipamentos, &talentos)
	if err != nil {
		fmt.Println("Erro ao gerar personagem:", err)
		return
	}

	fmt.Printf("Personagem: %+v\n", personagem)

	// Impressão formatada do personagem
	fmt.Printf("Personagem:\n")
	fmt.Printf("  Raça: %s\n", personagem.Raca)
	fmt.Printf("  Classe: %s\n", personagem.Classe)
	fmt.Printf("  Idade: %d\n", personagem.Idade)
	fmt.Printf("  Faixa Etária: %s\n", personagem.FaixaEtaria)
	fmt.Printf("  Atributos Chave: %v\n", personagem.AtributosChave)
	fmt.Printf("  Atributos:\n")
	for atributo, valor := range personagem.Atributos {
		fmt.Printf("    %s: %d\n", atributo, valor)
	}
	fmt.Printf("  Perícias:\n")
	for pericia, valor := range personagem.Pericias {
		fmt.Printf("    %s: %d\n", pericia, valor)
	}
	fmt.Printf("  Talentos:\n")
	for nome, talento := range personagem.Talentos {
		fmt.Printf("    %s: Nível %d\n", nome, talento.Nivel)
	}
}
