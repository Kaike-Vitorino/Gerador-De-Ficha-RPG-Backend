package main

import "fmt"

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

	atributosData, err := carregarAtributos("data/atributos.json")
	if err != nil {
		fmt.Println("Erro ao carregar atributos:", err)
		return
	}

	// Defina o número de pontos de XP adicionais que deseja distribuir
	pontosXP := 50

	personagem, err := gerarPersonagemAleatorio(racas, classes, status, &talentos, atributosData, pontosXP)
	if err != nil {
		fmt.Println("Erro ao gerar personagem:", err)
		return
	}

	equipamentosGerados, err := GerarArma(personagem.Classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		fmt.Println("Erro ao gerar equipamentos:", err)
		return
	}
	personagem.Equipamentos = equipamentosGerados

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
	fmt.Printf("  Pericias:\n")
	for pericia, valor := range personagem.Pericias {
		fmt.Printf("    %s: %d\n", pericia, valor)
	}
	fmt.Printf("  Talentos:\n")
	for nome, talento := range personagem.Talentos {
		fmt.Printf("    %s: Nível %d\n", nome, talento.Nivel)
	}
	fmt.Printf("  Equipamentos: %v\n", personagem.Equipamentos)
}
