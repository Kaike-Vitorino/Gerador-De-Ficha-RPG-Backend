package character_core

import (
	"image"
	DataChar "psBackKG/model/character_data"
	Sheet "psBackKG/model/character_sheet_imaging"
)

// Função principal para adicionar todas as informações na character_sheet_imaging
func AdicionarTodasInformacoesFicha(imagem *image.RGBA, personagem *DataChar.Personagem, coordenadas *Sheet.Coordenadas, armasEscolhidas []string, equipamentos *DataChar.Equipamentos, bonusArma1, danoArma1, bonusArma2, danoArma2 string) {
	Sheet.AdicionarInformacoesBasicas(imagem, personagem, coordenadas)
	Sheet.AdicionarAtributos(imagem, personagem, coordenadas)
	Sheet.AdicionarTalentos(imagem, personagem, coordenadas)
	Sheet.AdicionarPericias(imagem, personagem, coordenadas)
	Sheet.AdicionarArmas(imagem, armasEscolhidas, coordenadas)

	if personagem.Classe == "Rider" {
		// Adiciona a primeira arma
		Sheet.AdicionarInfoArmas(imagem, DataChar.Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)

		// Cria uma nova coordenada para a segunda arma
		novaCoordenada := Sheet.Coordenadas{
			CoordenadasSimples: map[string]Sheet.Coord{
				"info_armas_cord": {
					coordenadas.CoordenadasSimples["info_armas_cord"][0],
					coordenadas.CoordenadasSimples["info_armas_cord"][1] + 70,
				},
			},
		}

		// Adiciona a segunda arma com a nova coordenada
		Sheet.AdicionarInfoArmas(imagem, DataChar.Arma{Bonus: bonusArma2, Dano: danoArma2}, &novaCoordenada, armasEscolhidas)
	} else {
		// Adiciona a arma normalmente se não for "Rider"
		Sheet.AdicionarInfoArmas(imagem, DataChar.Arma{Bonus: bonusArma1, Dano: danoArma1}, coordenadas, armasEscolhidas)
	}

	if personagem.Classe == "Guerreiro" {
		Sheet.EscreverArmadura(imagem, "Couro", coordenadas.ArmorCoordenadas)
	}

	Sheet.EscreverTextoEmVariasCoordenadasJOGADOR(imagem, coordenadas.PlayerCoordenadas)
	Sheet.EscreverTextoEmVariasCoordenadasMESTRE(imagem, coordenadas.MestreCoordenadas)
}
