package handler

import (
	"fmt"
)

// Função para gerar e juntar as informações da ficha
func gerarInfoFicha() (*Personagem, *Coordenadas, []string, string, string, string, string, error) {
	// Carregar dados de raças
	racas, err := NewPersonagemRacas("data/racas.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar raças: %v", err)
	}

	// Carregar dados de classes
	classes, err := NewPersonagemClasses("data/classes.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar classes: %v", err)
	}

	// Carregar dados de status
	status, err := NewPersonagemStatus("data/atributos.json", "data/pericias.json", "data/talentos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar status: %v", err)
	}

	// Carregar dados de talentos
	talentos, err := carregarTalentos("data/talentos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar talentos: %v", err)
	}

	// Carregar dados de atributos
	atributosData, err := carregarAtributos("data/atributos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar dados de atributos: %v", err)
	}

	// Carregar dados de ficha
	equipamentos, err := CarregarEquipamentos()
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar ficha: %v", err)
	}

	// Definir quantidade de XP
	var pontosXP int
	fmt.Print("Quantidade de XP: ")
	fmt.Scan(&pontosXP)

	// Gerar personagem aleatório
	personagem, err := gerarPersonagemAleatorio(racas, classes, status, &talentos, atributosData, pontosXP, equipamentos)
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao gerar personagem: %v", err)
	}

	// Gerar armas para o personagem
	armasEscolhidas, err := GerarArma(personagem.Classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao gerar armas: %v", err)
	}

	// Pegar informações das armas
	var bonusArma1, danoArma1, bonusArma2, danoArma2 string
	if personagem.Classe == "Rider" && len(armasEscolhidas) == 2 {
		infoArma1 := equipamentos.ListaArmas[armasEscolhidas[0]]
		infoArma2 := equipamentos.ListaArmas[armasEscolhidas[1]]
		bonusArma1 = infoArma1.Bonus
		danoArma1 = infoArma1.Dano
		bonusArma2 = infoArma2.Bonus
		danoArma2 = infoArma2.Dano
	} else if len(armasEscolhidas) > 0 {
		infoArma := equipamentos.ListaArmas[armasEscolhidas[0]]
		bonusArma1 = infoArma.Bonus
		danoArma1 = infoArma.Dano
	}

	// Carregar coordenadas
	coordenadas, err := NewCoordinates()
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar coordenadas: %v", err)
	}

	return personagem, coordenadas, armasEscolhidas, bonusArma1, danoArma1, bonusArma2, danoArma2, nil
}

/*
	// Exibir informações do personagem gerado
	fmt.Println("\n--- Ficha do Personagem ---")
	fmt.Printf("Raça: %s\n", personagem.Raca)
	fmt.Printf("Classe: %s\n", personagem.Classe)
	fmt.Printf("=============================\n")
	fmt.Printf("Atributo(s) Chave: %v\n", personagem.AtributosChave)
	fmt.Printf("Atributos: %v\n", personagem.Atributos)
	fmt.Printf("=============================\n")
	fmt.Printf("Idade: %d\n", personagem.Idade)
	fmt.Printf("Faixa Etária: %s\n", personagem.FaixaEtaria)
	fmt.Printf("=============================\n")
	fmt.Printf("Talentos:\n")
	for talento, info := range personagem.Talentos {
		fmt.Printf("%s - Nível %d\n", talento, info.Nivel)
	}
	fmt.Printf("=============================\n")
	fmt.Printf("Perícias: %v\n", personagem.Pericias)
	fmt.Printf("=============================\n")
	fmt.Printf("Equipamentos: %v\n", personagem.Equipamentos)
	if personagem.ArtefatoMusicalEscolhido != "" {
		fmt.Printf("Artefato Musical: %s\n", personagem.ArtefatoMusicalEscolhido)
	}
	fmt.Printf("=============================\n")
	fmt.Printf("Equipamentos Escolhidas:\n")
	for _, arma := range personagem.Equipamentos {
		if infoArma, ok := ficha.ListaArmas[arma]; ok {
			fmt.Printf("Arma: %s - Bônus: %s, Dano: %s\n", arma, infoArma.Bonus, infoArma.Dano)
		}
	}

*/
