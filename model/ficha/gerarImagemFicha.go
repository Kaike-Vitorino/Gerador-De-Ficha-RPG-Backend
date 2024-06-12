package ficha

import (
	"fmt"
	"image"
)

// Adiciona informações básicas (raça, classe, idade e faixa etária)
func adicionarInformacoesBasicas(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	adicionarTextoNaFichaY(imagem, personagem.Raca, coordenadas.CoordenadasSimples["raca_cord"], 0, 0)
	adicionarTextoNaFichaY(imagem, personagem.Classe, coordenadas.CoordenadasSimples["classe_cord"], 0, 0)
	adicionarTextoNaFichaY(imagem, personagem.Idade, coordenadas.CoordenadasSimples["idade_cord"], 0, 0)
	adicionarTextoNaFichaY(imagem, personagem.FaixaEtaria, coordenadas.CoordenadasSimples["faixa_etaria_cord"], 0, 0)
}

// Adiciona atributos
func adicionarAtributos(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	for atributo, valor := range personagem.Atributos {
		adicionarTextoNaFichaY(imagem, valor, coordenadas.CoordenadasSimples[atributo+"_cord"], 0, 0)
	}
}

// Adiciona talentos
func adicionarTalentos(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	espacamento := 70
	i := 0
	for talento, nivel := range personagem.Talentos {
		adicionarTextoNaFichaY(imagem, talento, coordenadas.CoordenadasSimples["talento_cord"], espacamento*i, 0)
		adicionarTextoNaFichaY(imagem, nivel.Nivel, coordenadas.CoordenadasSimples["nivel_talento_cord"], espacamento*i, 0)
		i++
	}
}

// Adiciona perícias
func adicionarPericias(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
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
		adicionarTextoNaFichaY(imagem, valorFormatado, coordenadas.CoordenadasSimples[pericia+"_cord"], 0, 0)
	}
}

// Adiciona armas
func adicionarArmas(imagem *image.RGBA, armasEscolhidas []string, coordenadas *Coordenadas) {
	if len(armasEscolhidas) > 1 {
		adicionarTextoNaFichaY(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 0, 0)
		adicionarTextoNaFichaY(imagem, armasEscolhidas[1], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 70, 0)
	} else {
		adicionarTextoNaFichaComDefault(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"])
	}
}

// Adiciona informações detalhadas das armas
func adicionarInfoArmas(imagem *image.RGBA, infoArmas Arma, coordenadas *Coordenadas, armasEscolhidas []string) {
	// Prints para debug
	fmt.Printf("Equipamentos Escolhidas: %v\n", armasEscolhidas)
	fmt.Printf("Informações das Equipamentos: Bonus - %s, Dano - %s\n", infoArmas.Bonus, infoArmas.Dano)
	fmt.Printf("Coordenadas info_armas_cord: %v\n", coordenadas.CoordenadasSimples["info_armas_cord"])

	if len(armasEscolhidas) > 1 {
		adicionarTextoNaFichaY(imagem, infoArmas.Bonus, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 0)
		adicionarTextoNaFichaY(imagem, infoArmas.Dano, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 140)
	} else {
		adicionarTextoNaFichaY(imagem, infoArmas.Bonus, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 0)
		adicionarTextoNaFichaY(imagem, infoArmas.Dano, coordenadas.CoordenadasSimples["info_armas_cord"], 0, 140)
	}
}
