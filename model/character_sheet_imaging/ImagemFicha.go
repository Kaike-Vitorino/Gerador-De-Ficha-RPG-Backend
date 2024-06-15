package character_sheet_imaging

import (
	"fmt"
	"image"
	DataChar "psBackKG/model/character_data"
)

// Adiciona informações básicas (raça, classe, idade e faixa etária)
func AdicionarInformacoesBasicas(imagem *image.RGBA, personagem *DataChar.Personagem, coordenadas *Coordenadas) {
	AdicionarTextoNaFichaY(imagem, personagem.Raca, coordenadas.CoordenadasSimples["raca_cord"], 0, 0)
	AdicionarTextoNaFichaY(imagem, personagem.Classe, coordenadas.CoordenadasSimples["classe_cord"], 0, 0)
	AdicionarTextoNaFichaY(imagem, fmt.Sprintf("%d", personagem.Idade), coordenadas.CoordenadasSimples["idade_cord"], 0, 0)
	AdicionarTextoNaFichaY(imagem, personagem.FaixaEtaria, coordenadas.CoordenadasSimples["faixa_etaria_cord"], 0, 0)
}

// Adiciona atributos
func AdicionarAtributos(imagem *image.RGBA, personagem *DataChar.Personagem, coordenadas *Coordenadas) {
	for atributo, valor := range personagem.Atributos {
		AdicionarTextoNaFichaY(imagem, fmt.Sprintf("%d", valor), coordenadas.CoordenadasSimples[atributo+"_cord"], 0, 0)
	}
}

// Adiciona talentos
func AdicionarTalentos(imagem *image.RGBA, personagem *DataChar.Personagem, coordenadas *Coordenadas) {
	espacamento := 70
	i := 0
	for _, talento := range personagem.Talentos {
		AdicionarTextoNaFichaY(imagem, talento.Nome, coordenadas.CoordenadasSimples["talento_cord"], espacamento*i, 0)
		AdicionarTextoNaFichaY(imagem, talento.Nivel, coordenadas.CoordenadasSimples["nivel_talento_cord"], espacamento*i, 0)
		i++
	}
}

// Adiciona perícias
func AdicionarPericias(imagem *image.RGBA, personagem *DataChar.Personagem, coordenadas *Coordenadas) {
	periciasDistribuidas := personagem.Pericias

	// Adicione todas as perícias que faltarem com valor 0
	periciasFaltando := map[string]int{
		"Potencia": 0, "Resiliência": 0, "Luta": 0, "Artesanato": 0,
		"Furtividade": 0, "Artimanha": 0, "Movimentação": 0, "Pontaria": 0,
		"Patrulha": 0, "Tradição": 0, "Sobrevivência": 0, "Discernimento": 0,
		"Manipulação": 0, "Atuação": 0, "Cura": 0, "Adestramento": 0,
	}

	for pericia, valor := range periciasDistribuidas {
		periciasFaltando[pericia] = valor
	}

	for pericia, valor := range periciasFaltando {
		valorFormatado := fmt.Sprintf("+%d", valor)
		AdicionarTextoNaFichaY(imagem, valorFormatado, coordenadas.CoordenadasSimples[pericia+"_cord"], 0, 0)
	}
}

// Adiciona armas
func AdicionarArmas(imagem *image.RGBA, armasEscolhidas []string, coordenadas *Coordenadas) {
	if len(armasEscolhidas) > 0 {
		AdicionarTextoNaFichaY(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 0, 0)
		AdicionarTextoNaFichaY(imagem, armasEscolhidas[1], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 70, 0)
	} else {
		AdicionarTextoNaFichaComDefault(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"])
	}
}

// Adiciona informações detalhadas das armas
func AdicionarInfoArmas(imagem *image.RGBA, infoArmas DataChar.Arma, coordenadas *Coordenadas, armasEscolhidas []string) {
	// Prints para debug
	fmt.Printf("Equipamentos Escolhidas: %v\n", armasEscolhidas)
	fmt.Printf("Informações das Equipamentos: Bonus - %s, Dano - %s\n", infoArmas.Bonus, infoArmas.Dano)
	fmt.Printf("Coordenadas info_armas_cord: %v\n", coordenadas.CoordenadasSimples["info_armas_cord"])

	if len(armasEscolhidas) > 1 {
		AdicionarTextoNaFichaY(imagem, infoArmas.Bonus, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 0)
		AdicionarTextoNaFichaY(imagem, infoArmas.Dano, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 140)
	} else {
		AdicionarTextoNaFichaY(imagem, infoArmas.Bonus, coordenadas.CoordenadasSimples["info_armas_cord"], 70, 0)
		AdicionarTextoNaFichaY(imagem, infoArmas.Dano, coordenadas.CoordenadasSimples["info_armas_cord"], 70, 140)
	}
}
