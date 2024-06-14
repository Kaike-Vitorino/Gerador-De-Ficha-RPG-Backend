package main

import (
	"fmt"
	"log"
	"psBackKG/model/character_logic"
)

func main() {
	classes, err := character_logic.CarregarClasses("data/classes.json")
	if err != nil {
		log.Fatalf("Erro ao carregar classes: %v", err)
	}

	equipamentos, err := charbuilder.CarregarEquipamentos("data/equipamentos.json")
	if err != nil {
		log.Fatalf("Erro ao carregar equipamentos: %v", err)
	}

	// Exemplo de geração de arma para a classe Bardo
	arma, err := charbuilder.GerarArma("Bardo", classes, equipamentos)
	if err != nil {
		log.Fatalf("Erro ao gerar arma: %v", err)
	}

	fmt.Printf("Arma escolhida para Bardo: %s\n", arma)
}
