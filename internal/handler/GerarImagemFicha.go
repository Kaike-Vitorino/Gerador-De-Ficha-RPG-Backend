package handler

import (
	"image"
)

// Função principal para adicionar todas as informações na ficha
func adicionarTodasInformacoesFicha(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas, armasEscolhidas []string, equipamentos *Equipamentos, bonusArma1, danoArma1, bonusArma2, danoArma2 string) {
	adicionarInformacoesBasicas(imagem, personagem, coordenadas)
	adicionarAtributos(imagem, personagem, coordenadas)
	adicionarTalentos(imagem, personagem, coordenadas)
	adicionarPericias(imagem, personagem, coordenadas)
	adicionarArmas(imagem, armasEscolhidas, coordenadas)

	if personagem.Classe == "Rider" {
		adicionarInfoArmas(imagem, Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)
		adicionarInfoArmas(imagem, Arma{Bonus: bonusArma2, Dano: danoArma2}, coordenadas, armasEscolhidas)
	} else {
		adicionarInfoArmas(imagem, Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)
	}

	if personagem.Classe == "Guerreiro" {
		escreverArmadura(imagem, "Couro", coordenadas.ArmorCoordenadas)
	}

	escreverTextoEmVariasCoordenadasJOGADOR(imagem, coordenadas.PlayerCoordenadas)
	escreverTextoEmVariasCoordenadasMESTRE(imagem, coordenadas.MestreCoordenadas)
}
