package character_core

import (
	DataChar "psBackKG/model/character_data"
	LogicChar "psBackKG/model/character_logic"
)

// Função para Gerar um personagem baseado nas escolhas do usuário
func GerarPersonagemCustomizado(racaEscolhida, classeEscolhida string, idadeEscolhida int, racas *DataChar.PersonagemRacas, classes *DataChar.PersonagemClasses, status *DataChar.PersonagemStatus, talentos *DataChar.Talentos, atributos DataChar.Atributos, pontosXP int, equipamentos *DataChar.Equipamentos) (*DataChar.Personagem, error) {

	var raca string
	var racaInfo DataChar.Raca

	if racaEscolhida == "random" {
		raca, racaInfo = LogicChar.GerarRaca(racas)
	} else {
		raca = racaEscolhida
		racaInfo = racas.RacasInfo[raca]
	}

	var classe string
	if classeEscolhida == "random" {
		classe = LogicChar.GerarClasse(raca, racas)
	} else {
		classe = classeEscolhida
	}

	atributosChave := LogicChar.ObterAtributosChave(classe, racaInfo, classes.ClasseInfo)

	var idade int
	var faixaEtaria string
	if idadeEscolhida == -1 { // Usuário escolheu idade aleatória
		idade, faixaEtaria = LogicChar.CalcularIdade(raca, racas)
	} else {
		idade = idadeEscolhida
		faixaEtaria = LogicChar.ObterFaixaEtaria(raca, idade, racas)
	}

	atributosDistribuidos := LogicChar.EscolherAtributos(faixaEtaria, atributosChave, atributos)
	periciasDistribuidas := LogicChar.DistribuirPericias(faixaEtaria, classe, status.Pericias, classes.ClasseInfo)
	talentosDistribuidos := LogicChar.EscolherTalentos(classe, raca, faixaEtaria, racas.RacasInfo, classes.TalentosClasses, talentos.TalentosGerais)
	talentosDistribuidos, periciasDistribuidas = LogicChar.DividirXP(talentosDistribuidos, periciasDistribuidas, classe, status.Pericias, talentos.TalentosGerais, pontosXP)

	artefatoMusicalEscolhido := ""
	if classe == "Bardo" {
		artefatoMusicalEscolhido = DataChar.EscolherArtefatoMusical(classes, classe)
	}

	armasEscolhidas, err := LogicChar.GerarArma(classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		return nil, err
	}

	return &DataChar.Personagem{
		Raca:                     raca,
		Classe:                   classe,
		AtributosChave:           atributosChave,
		Atributos:                atributosDistribuidos,
		Idade:                    idade,
		FaixaEtaria:              faixaEtaria,
		Pericias:                 periciasDistribuidas,
		Talentos:                 talentosDistribuidos,
		ArtefatoMusicalEscolhido: artefatoMusicalEscolhido,
		Equipamentos:             armasEscolhidas,
	}, nil
}
