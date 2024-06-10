package main

import (
	"fmt"
)

// Função para gerar e juntar as informações da ficha
func gerarInfoFicha(classe, raca string, atributosChave []string, idade int, faixaEtaria string, atributosRandomizados map[string]int, talentosEscolhidos map[string]Talento, periciasDistribuidas map[string]int, armasEscolhidas []string, classeInfo map[string]Classe, equipamentos *Equipamentos) {
	quantidadeItens := classeInfo[classe].Equipamentos.Itens
	var itensRandomizados []string

	if contemItem2([]string{"Arco Curto", "Arco Longo"}, equipamentos.ArmaEscolhida) || contemItem2(armasEscolhidas, "Arco Curto") || contemItem2(armasEscolhidas, "Arco Longo") {
		itensRandomizados = randomSample(equipamentos.ItensComercio, quantidadeItens)
	} else {
		itensRandomizados = randomSample(equipamentos.ItensComercio[2:], quantidadeItens)
	}
	equipamentos.Items = append(equipamentos.Items, itensRandomizados...)

	fmt.Printf("Itens Randomizados: %v\n", itensRandomizados)

	cavalo := classeInfo[classe].Equipamentos.Cavalo
	armaduraDisponivel := classeInfo[classe].Equipamentos.Armadura
	var infoArmaduras Armadura
	if armaduraDisponivel != nil {
		infoArmaduras = equipamentos.ListaArmaduras[*armaduraDisponivel]
	}
	dxComida := classeInfo[classe].DadosRecurso["Comida"]
	dxAgua := classeInfo[classe].DadosRecurso["Água"]
	var infoArma1, infoArma2 Arma

	fmt.Println("\n--- Ficha do Personagem ---")
	fmt.Printf("Raça: %s\n", raca)
	fmt.Printf("Classe: %s\n", classe)
	fmt.Printf("=============================\n")
	fmt.Printf("Atributo(s) Chave: %v\n", atributosChave)
	fmt.Printf("Atributos: %v\n", atributosRandomizados)
	fmt.Printf("=============================\n")
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Faixa Etária: %s\n", faixaEtaria)
	fmt.Printf("=============================\n")
	fmt.Printf("Talentos:\n")
	for talento, info := range talentosEscolhidos {
		fmt.Printf("%s - Nível %d\n", talento, info.Nivel)
	}
	fmt.Printf("=============================\n")
	fmt.Printf("Perícias: %v\n", periciasDistribuidas)
	fmt.Printf("=============================\n")
	fmt.Printf("Equipamentos: %v\n", equipamentos.Items)

	if classe == "Rider" {
		armasEscolhidas1, armasEscolhidas2 := armasEscolhidas[0], armasEscolhidas[1]
		infoArma1 = equipamentos.ListaArmas[armasEscolhidas1]
		infoArma2 = equipamentos.ListaArmas[armasEscolhidas2]
		fmt.Printf("Sua 1ª arma: %s\n", armasEscolhidas1)
		fmt.Printf("Infos da sua 1ª arma: %v\n", infoArma1)
		fmt.Printf("Sua 2ª arma: %s\n", armasEscolhidas2)
		fmt.Printf("Infos da sua 2ª arma: %v\n", infoArma2)
	} else {
		infoArmas := equipamentos.ListaArmas[equipamentos.ArmaEscolhida]
		fmt.Printf("Sua arma: %s\n", equipamentos.ArmaEscolhida)
		fmt.Printf("Infos da sua arma: %v\n", infoArmas)
	}

	if armaduraDisponivel != nil {
		fmt.Printf("Sua armadura: %s\n", *armaduraDisponivel)
		fmt.Printf("Infos da armadura: %v\n", infoArmaduras)
	}

	if cavalo != 0 {
		fmt.Println("Você tem um cavalo")
		fmt.Println("Essas são as infos da sua montaria:")
		// Exemplo: Força: 5, Agilidade: 4, Perícias: Movimentação: +2, Patrulha: +3, Dano: 1
	}

	if equipamentos.ArtefatoMusicalEscolhido != "" {
		fmt.Printf("Seu instrumento: %s\n", equipamentos.ArtefatoMusicalEscolhido)
	}
	fmt.Printf("Comida: %s\n", dxComida)
	fmt.Printf("Água: %s\n", dxAgua)

	if contemItem2(equipamentos.Items, "Arco Curto") || contemItem2(equipamentos.Items, "Arco Longo") {
		fmt.Println("Flecha: D10")
	}

	fmt.Println("--------------------------\n")
}
