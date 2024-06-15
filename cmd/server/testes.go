package main

import (
	"fmt"
	"log"
	DataChar "psBackKG/model/character_data"
	LogicChar "psBackKG/model/character_logic"
)

func main2() {
	// Carregar dados de classes e equipamentos
	classes, talentosClasses, err := DataChar.CarregarClasses("data/classes.json")
	if err != nil {
		log.Fatal("Erro ao carregar classes:", err, talentosClasses)
	}

	equipamentos, err := DataChar.CarregarEquipamentos("data/equipamentos.json")
	if err != nil {
		log.Fatal("Erro ao carregar equipamentos:", err)
	}

	// Testar função GerarArma com a classe "Bardo"
	classe := "Bardo"
	equipamentosGerados, err := LogicChar.GerarArma(classe, classes, &equipamentos)
	if err != nil {
		log.Fatal("Erro ao gerar arma para a classe Bardo:", err)
	}

	fmt.Printf("Equipamentos gerados para a classe %s: %v\n", classe, equipamentosGerados)

	// Obter informações da arma a partir dos equipamentos
	if len(equipamentosGerados) > 0 {
		armaEscolhida := equipamentosGerados[0]
		infoArma, existe := equipamentos.Armas1M[armaEscolhida]
		if !existe {
			infoArma, existe = equipamentos.Armas2M[armaEscolhida]
		}
		if !existe {
			infoArma, existe = equipamentos.ArmasDistancia1M[armaEscolhida]
		}
		if !existe {
			infoArma, existe = equipamentos.ArmasDistancia2M[armaEscolhida]
		}
		if !existe {
			log.Fatal("Arma não encontrada nos dados de equipamentos:", armaEscolhida)
		}

		// Exibir informações da arma
		fmt.Printf("Arma escolhida: %s\n", armaEscolhida)
		fmt.Printf("Bonus: %s, Dano: %s\n", infoArma.Bonus, infoArma.Dano)
	}
}
