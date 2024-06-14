package character_core

import (
	DataChar "psBackKG/model/character_data"
	LogicChar "psBackKG/model/character_logic"
)

// Função para Gerar um character_logic aleatório
func GerarPersonagemAleatorio(racas *DataChar.PersonagemRacas, classes *DataChar.PersonagemClasses, status *DataChar.PersonagemStatus, talentos *DataChar.Talentos, atributos DataChar.Atributos, pontosXP int, equipamentos *DataChar.Equipamentos) (*DataChar.Personagem, error) {

	raca, racaInfo := LogicChar.GerarRaca(racas)
	classe := LogicChar.GerarClasse(raca, racas)
	//classe = "Rider"
	atributosChave := LogicChar.ObterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := LogicChar.CalcularIdade(raca, racas)
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

		Equipamentos: armasEscolhidas, // Aqui você pode preencher com os character_sheet_imaging apropriados
	}, nil
}
