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

	adicionarArmas(imagem, armasEscolhidas, coordenadas, personagem.Classe)

	// Adiciona informações da armadura se for guerreiro
	if personagem.Classe == "Guerreiro" {
		escreverArmadura(imagem, "Couro", coordenadas.ArmorCoordenadas)
	}

	// Adiciona escolha do jogador e do mestre
	escolhaDoJogador := "Escolha do JOGADOR"
	escolhaDoMestre := "Escolha do MESTRE"
	escreverTextoEmVariasCoordenadasJOGADOR(imagem, escolhaDoJogador, coordenadas.PlayerCoordenadas)
	escreverTextoEmVariasCoordenadasMESTRE(imagem, escolhaDoMestre, coordenadas.MestreCoordenadas)
}
