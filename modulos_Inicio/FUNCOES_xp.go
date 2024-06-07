package main

import (
	"fmt"
)

// Função principal para dividir XP entre talentos e Pericias
func DividirXP(talentosEscolhidos map[string]Talento, periciasDistribuidas map[string]int, classe string, pericias map[string]string, talentosGerais map[string][]string, pontosXP int) (map[string]Talento, map[string]int) {
	periciasDisponiveis := make([]string, 0, len(pericias))
	for pericia := range pericias {
		periciasDisponiveis = append(periciasDisponiveis, pericia)
	}
	talentosDisponiveis := make([]string, len(talentosGerais[classe]))
	copy(talentosDisponiveis, talentosGerais[classe])

	for pontosXP > 0 && len(talentosEscolhidos) < 12 {
		opcoes := []string{"Pericia nova", "Talento novo", "Aumentar nivel de Pericia", "Aumentar nivel de talento"}
		escolha := randomChoice(opcoes)

		switch escolha {
		case "Pericia nova":
			if len(periciasDisponiveis) == 0 {
				fmt.Println("Todas as Pericias estão distribuídas. Escolhendo novamente.")
				continue
			}
			pontosXP = AdicionarNovaPericia(&periciasDisponiveis, periciasDistribuidas, pontosXP)

		case "Talento novo":
			if len(talentosDisponiveis) == 0 {
				fmt.Println("Todos os talentos disponíveis já foram escolhidos. Escolhendo novamente.")
				continue
			}
			if len(talentosEscolhidos) < 12 {
				pontosXP = AdicionarNovoTalento(&talentosDisponiveis, talentosEscolhidos, pontosXP)
			}

		case "Aumentar nivel de Pericia":
			pontosXP = AumentarNivelPericia(periciasDistribuidas, pontosXP)

		case "Aumentar nivel de talento":
			pontosXP = AumentarNivelTalento(talentosEscolhidos, pontosXP)

		default:
			fmt.Println("Escolha inválida.")
		}
	}

	return talentosEscolhidos, periciasDistribuidas
}

// Função auxiliar para adicionar uma nova Pericia
func AdicionarNovaPericia(periciasDisponiveis *[]string, periciasDistribuidas map[string]int, pontosXP int) int {
	fmt.Println("Pericia Nova!")
	custo := 5
	novaPericia := randomChoice(*periciasDisponiveis)
	*periciasDisponiveis = remove(*periciasDisponiveis, novaPericia)
	periciasDistribuidas[novaPericia] = 1
	return pontosXP - custo
}

// Função auxiliar para adicionar um novo talento
func AdicionarNovoTalento(talentosDisponiveis *[]string, talentosEscolhidos map[string]Talento, pontosXP int) int {
	fmt.Println("Talento Novo!")

	custo := 3
	novoTalento := randomChoice(*talentosDisponiveis)
	*talentosDisponiveis = remove(*talentosDisponiveis, novoTalento)
	talentosEscolhidos[novoTalento] = Talento{Nome: novoTalento, Nivel: 1}
	return pontosXP - custo
}

// Função auxiliar para aumentar o nível de uma Pericia
func AumentarNivelPericia(periciasDistribuidas map[string]int, pontosXP int) int {
	fmt.Println("Nivel de pericia aumentado!")

	periciasDisponiveisParaAumentar := make([]string, 0)
	for pericia, nivel := range periciasDistribuidas {
		if nivel < 3 {
			periciasDisponiveisParaAumentar = append(periciasDisponiveisParaAumentar, pericia)
		}
	}

	if len(periciasDisponiveisParaAumentar) == 0 {
		fmt.Println("Todas as Pericias estão no nível máximo. Escolhendo novamente.")
		return pontosXP
	}

	periciaAumentar := randomChoice(periciasDisponiveisParaAumentar)
	nivelPericia := periciasDistribuidas[periciaAumentar]
	custo := 5 * nivelPericia
	periciasDistribuidas[periciaAumentar]++
	return pontosXP - custo
}

// Função auxiliar para aumentar o nível de um talento
func AumentarNivelTalento(talentosEscolhidos map[string]Talento, pontosXP int) int {
	fmt.Println("Nivel de talento aumentado!")

	talentosDisponiveisParaAumentar := make([]string, 0)
	for talento, info := range talentosEscolhidos {
		if info.Nivel < 3 {
			talentosDisponiveisParaAumentar = append(talentosDisponiveisParaAumentar, talento)
		}
	}

	if len(talentosDisponiveisParaAumentar) == 0 {
		fmt.Println("Todos os talentos estão no nível máximo. Escolhendo novamente.")
		return pontosXP
	}

	talentoAumentar := randomChoice(talentosDisponiveisParaAumentar)
	nivelTalento := talentosEscolhidos[talentoAumentar].Nivel
	custo := 3 * nivelTalento
	talentosEscolhidos[talentoAumentar] = Talento{Nome: talentoAumentar, Nivel: nivelTalento + 1}
	return pontosXP - custo
}
