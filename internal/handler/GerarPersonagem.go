package handler

// Personagem representa todas as informações de um personagem
type Personagem struct {
	Raca                     string
	Classe                   string
	Idade                    int
	FaixaEtaria              string
	AtributosChave           []string
	Atributos                map[string]int
	ArtefatoMusicalEscolhido string
	ArmaEscolhida            string
	Equipamentos             []string
	Pericias                 map[string]int
	Talentos                 map[string]Talento
}

// Função para gerar um personagem aleatório
func gerarPersonagemAleatorio(racas *PersonagemRacas, classes *PersonagemClasses, status *PersonagemStatus, talentos *Talentos, atributos Atributos, pontosXP int, equipamentos *Equipamentos) (*Personagem, error) {
	raca, racaInfo := gerarRaca(racas)
	classe := gerarClasse(raca, racas)
	classe = "Rider"
	atributosChave := obterAtributosChave(classe, racaInfo, classes.ClasseInfo)
	idade, faixaEtaria := calcularIdade(raca, racas)
	atributosDistribuidos := escolherAtributos(faixaEtaria, atributosChave, atributos)
	periciasDistribuidas := DistribuirPericias(faixaEtaria, classe, status.Pericias, classes.ClasseInfo)
	talentosDistribuidos := EscolherTalentos(classe, raca, faixaEtaria, racas.RacasInfo, classes.TalentosClasses, talentos.TalentosGerais)
	talentosDistribuidos, periciasDistribuidas = DividirXP(talentosDistribuidos, periciasDistribuidas, classe, status.Pericias, talentos.TalentosGerais, pontosXP)

	artefatoMusicalEscolhido := ""
	if classe == "Bardo" {
		artefatoMusicalEscolhido = escolherArtefatoMusical(classes, classe)
	}

	armasEscolhidas, err := GerarArma(classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		return nil, err
	}

	return &Personagem{
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
