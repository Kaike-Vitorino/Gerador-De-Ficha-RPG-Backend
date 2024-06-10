package main

import (
	"image"
)

// Função principal para adicionar todas as informações na ficha
func adicionarTodasInformacoesFicha(imagem *image.RGBA, personagem *Personagem, coordenadas *Coordenadas, armasEscolhidas []string, classeInfo map[string]Classe, equipamentos *Equipamentos) {
	adicionarInformacoesBasicas(imagem, personagem, coordenadas)
	adicionarAtributos(imagem, personagem, coordenadas)
	adicionarTalentos(imagem, personagem, coordenadas)
	adicionarPericias(imagem, personagem, coordenadas)
	adicionarArmas(imagem, armasEscolhidas, coordenadas)

	infoArmas := equipamentos.ListaArmas[equipamentos.ArmaEscolhida]

	adicionarInfoArmas(imagem, infoArmas, coordenadas, armasEscolhidas)

	if personagem.Classe == "Guerreiro" {
		escreverArmadura(imagem, "Couro", coordenadas.ArmorCoordenadas)
	}

	escreverTextoEmVariasCoordenadasJOGADOR(imagem, coordenadas.PlayerCoordenadas)
	escreverTextoEmVariasCoordenadasMESTRE(imagem, coordenadas.MestreCoordenadas)
}
