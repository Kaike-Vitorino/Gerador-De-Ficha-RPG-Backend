package main

import (
	"fmt"
)

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
		fmt.Println("Erro ao carregar dados de atributos:", err)
		return
	}

	var pontosXP int
	fmt.Print("Quantidade de XP: ")
	fmt.Scan(&pontosXP)

	personagem, err := gerarPersonagemAleatorio(racas, classes, status, &talentos, atributosData, pontosXP)
	if err != nil {
		fmt.Println("Erro ao gerar personagem:", err)
		return
	}

	armasEscolhidas, err := GerarArma(personagem.Classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		fmt.Println("Erro ao gerar armas:", err)
		return
	}

	gerarInfoFicha(
		personagem.Classe,
		personagem.Raca,
		personagem.AtributosChave,
		personagem.Idade,
		personagem.FaixaEtaria,
		personagem.Atributos,
		personagem.Talentos,
		personagem.Pericias,
		armasEscolhidas,
		classes.ClasseInfo,
		equipamentos,
	)
}
