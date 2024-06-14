package handler

import (
	"psBackKG/model/charbuilder"
	personagemPKG "psBackKG/model/personagem"
)

// Função para Gerar um personagem aleatório
func GerarPersonagemAleatorio(racas *charbuilder.PersonagemRacas, classes *charbuilder.PersonagemClasses, status *charbuilder.PersonagemStatus, talentos *charbuilder.Talentos, atributos charbuilder.Atributos, pontosXP int, equipamentos *charbuilder.Equipamentos) (*charbuilder.Personagem, error) {

	raca, racaInfo := personagemPKG.GerarRaca(racas)
	classe := personagemPKG.GerarClasse(raca, racas)
	//classe = "Rider"
	atributosChave := personagemPKG.ObterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := personagemPKG.CalcularIdade(raca, racas)
	atributosDistribuidos := personagemPKG.EscolherAtributos(faixaEtaria, atributosChave, atributos)
	periciasDistribuidas := personagemPKG.DistribuirPericias(faixaEtaria, classe, status.Pericias, classes.ClasseInfo)
	talentosDistribuidos := personagemPKG.EscolherTalentos(classe, raca, faixaEtaria, racas.RacasInfo, classes.TalentosClasses, talentos.TalentosGerais)
	talentosDistribuidos, periciasDistribuidas = personagemPKG.DividirXP(talentosDistribuidos, periciasDistribuidas, classe, status.Pericias, talentos.TalentosGerais, pontosXP)

	artefatoMusicalEscolhido := ""
	if classe == "Bardo" {
		artefatoMusicalEscolhido = charbuilder.EscolherArtefatoMusical(classes, classe)
	}

	armasEscolhidas, err := personagemPKG.GerarArma(classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		return nil, err
	}

	return &charbuilder.Personagem{
		Raca:                     raca,
		Classe:                   classe,
		AtributosChave:           atributosChave,
		Atributos:                atributosDistribuidos,
		Idade:                    idade,
		FaixaEtaria:              faixaEtaria,
		Pericias:                 periciasDistribuidas,
		Talentos:                 talentosDistribuidos,
		ArtefatoMusicalEscolhido: artefatoMusicalEscolhido,

		Equipamentos: armasEscolhidas, // Aqui você pode preencher com os ficha apropriados
	}, nil
}
