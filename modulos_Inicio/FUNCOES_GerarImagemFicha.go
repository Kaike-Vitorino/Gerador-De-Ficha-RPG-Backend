package main

import (
	"image"
)

// Adiciona informações básicas (raça, classe, idade e faixa etária)
func adicionarInformacoesBasicas(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	adicionarTextoNaFichaComDefault(imagem, personagem.Raca, coordenadas.CoordenadasSimples["raca_cord"])
	adicionarTextoNaFichaComDefault(imagem, personagem.Classe, coordenadas.CoordenadasSimples["classe_cord"])
	adicionarTextoNaFichaComDefault(imagem, personagem.Idade, coordenadas.CoordenadasSimples["idade_cord"])
	adicionarTextoNaFichaComDefault(imagem, personagem.FaixaEtaria, coordenadas.CoordenadasSimples["faixa_etaria_cord"])
}

// Adiciona atributos
func adicionarAtributos(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	for atributo, valor := range personagem.Atributos {
		adicionarTextoNaFichaComDefault(imagem, valor, coordenadas.CoordenadasSimples[atributo+"_cord"])
	}
}

// Adiciona talentos
func adicionarTalentos(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	espacamento := 70
	for i, talentoInfo := range personagem.Talentos {
		talento, info := talentoInfo.Nome, talentoInfo.Nivel
		adicionarTextoNaFicha(imagem, talento, coordenadas.CoordenadasSimples["talento_cord"], espacamento)
		adicionarTextoNaFicha(imagem, info, coordenadas.CoordenadasSimples["nivel_talento_cord"], espacamento)
		espacamento = +70
	}
}

// Adiciona perícias
func adicionarPericias(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas) {
	for pericia, valor := range personagem.Pericias {
		adicionarTextoNaFichaComDefault(imagem, valor, coordenadas.CoordenadasSimples[pericia+"_cord"])
	}
}

// Adiciona armas
func adicionarArmas(imagem *image.RGBA, armasEscolhidas []string, coordenadas *Coordenadas, classe string) {
	if len(armasEscolhidas) > 1 {
		adicionarTextoNaFicha(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 0)
		adicionarTextoNaFicha(imagem, armasEscolhidas[1], coordenadas.CoordenadasSimples["arma_escolhida_cord"], 70)
	} else {
		adicionarTextoNaFichaComDefault(imagem, armasEscolhidas[0], coordenadas.CoordenadasSimples["arma_escolhida_cord"])
	}
}

// Função principal para adicionar todas as informações na ficha
func adicionarInformacoesFicha(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas, armasEscolhidas []string, classeInfo map[string]Classe, equipamentos *Equipamentos) {
	adicionarInformacoesBasicas(imagem, personagem, coordenadas)
	adicionarAtributos(imagem, personagem, coordenadas)
	adicionarTalentos(imagem, personagem, coordenadas)
	adicionarPericias(imagem, personagem, coordenadas)
	adicionarArmas(imagem, armasEscolhidas, coordenadas, personagem.Classe)
}
