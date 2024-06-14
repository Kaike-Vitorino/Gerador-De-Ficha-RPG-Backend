package handler

import (
	"image"
	"psBackKG/model/charbuilder"
	"psBackKG/model/ficha"
)

// Função principal para adicionar todas as informações na ficha
func AdicionarTodasInformacoesFicha(imagem *image.RGBA, personagem *charbuilder.Personagem, coordenadas *ficha.Coordenadas, armasEscolhidas []string, equipamentos *charbuilder.Equipamentos, bonusArma1, danoArma1, bonusArma2, danoArma2 string) {
	ficha.AdicionarInformacoesBasicas(imagem, personagem, coordenadas)
	ficha.AdicionarAtributos(imagem, personagem, coordenadas)
	ficha.AdicionarTalentos(imagem, personagem, coordenadas)
	ficha.AdicionarPericias(imagem, personagem, coordenadas)
	ficha.AdicionarArmas(imagem, armasEscolhidas, coordenadas)

	if personagem.Classe == "Rider" {
		ficha.AdicionarInfoArmas(imagem, charbuilder.Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)
		ficha.AdicionarInfoArmas(imagem, charbuilder.Arma{Bonus: bonusArma2, Dano: danoArma2}, coordenadas, armasEscolhidas)
	} else {
		ficha.AdicionarInfoArmas(imagem, charbuilder.Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)
	}

	if personagem.Classe == "Guerreiro" {
		ficha.EscreverArmadura(imagem, "Couro", coordenadas.ArmorCoordenadas)
	}

	ficha.EscreverTextoEmVariasCoordenadasJOGADOR(imagem, coordenadas.PlayerCoordenadas)
	ficha.EscreverTextoEmVariasCoordenadasMESTRE(imagem, coordenadas.MestreCoordenadas)
}
