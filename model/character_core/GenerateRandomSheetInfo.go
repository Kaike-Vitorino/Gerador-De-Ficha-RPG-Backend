package character_core

import (
	"fmt"
	DataChar "psBackKG/model/character_data"
	LogicChar "psBackKG/model/character_logic"
	Sheet "psBackKG/model/character_sheet_imaging"
)

// Função para Gerar e juntar as informações da character_sheet_imaging
func GerarInfoFichaRandom() (*DataChar.Personagem, *Sheet.Coordenadas, []string, string, string, string, string, error) {
	// Carregar dados de raças
	racas, err := DataChar.NewPersonagemRacas("assets/racas.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar raças: %v", err)
	}

	// Carregar dados de classes
	classesMap, talentosClasses, err := DataChar.CarregarClasses("assets/classes.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar classes: %v", err)
	}
	classes := &DataChar.PersonagemClasses{
		ClasseInfo:      classesMap,
		TalentosClasses: talentosClasses,
	}

	// Carregar dados de status
	status, err := DataChar.NewPersonagemStatus("assets/atributos.json", "assets/pericias.json", "assets/talentos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar status: %v", err)
	}

	// Carregar dados de talentos
	talentos, err := DataChar.CarregarTalentos("assets/talentos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar talentos: %v", err)
	}

	// Carregar dados de atributos
	atributosData, err := DataChar.CarregarAtributos("assets/atributos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar dados de atributos: %v", err)
	}

	// Carregar dados de equipamentos
	equipamentos, err := DataChar.CarregarEquipamentos("assets/equipamentos.json")
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar equipamentos: %v", err)
	}

	// Definir quantidade de XP
	var pontosXP int
	fmt.Print("Quantidade de XP: ")
	fmt.Scan(&pontosXP)

	// Gerar personagem aleatório
	personagem, err := GerarPersonagemAleatorio(racas, classes, status, &talentos, atributosData, pontosXP, &equipamentos)
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao gerar personagem: %v", err)
	}

	// Gerar armas para o personagem
	armasEscolhidas, err := LogicChar.GerarArma(personagem.Classe, classes.ClasseInfo, &equipamentos)
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao gerar armas: %v", err)
	}

	// Pegar informações das armas
	var bonusArma1, danoArma1, bonusArma2, danoArma2 string
	if personagem.Classe == "Rider" && len(armasEscolhidas) == 2 {
		infoArma1 := obterInfoArma(armasEscolhidas[0], &equipamentos)
		infoArma2 := obterInfoArma(armasEscolhidas[1], &equipamentos)
		bonusArma1 = infoArma1.Bonus
		danoArma1 = infoArma1.Dano
		bonusArma2 = infoArma2.Bonus
		danoArma2 = infoArma2.Dano
	} else if len(armasEscolhidas) > 0 {
		infoArma := obterInfoArma(armasEscolhidas[0], &equipamentos)
		bonusArma1 = infoArma.Bonus
		danoArma1 = infoArma.Dano
	}

	// Carregar coordenadas
	coordenadas, err := Sheet.NewCoordinates()
	if err != nil {
		return nil, nil, nil, "", "", "", "", fmt.Errorf("Erro ao carregar coordenadas: %v", err)
	}

	return personagem, coordenadas, armasEscolhidas, bonusArma1, danoArma1, bonusArma2, danoArma2, nil
}
