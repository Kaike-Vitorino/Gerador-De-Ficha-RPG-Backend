package main_tasks

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	LogicChar "psBackKG/model/character_core"
	DataChar "psBackKG/model/character_data"
)

func GerarFichaSpecificCompleta(racaEscolhida, classeEscolhida string, idadeEscolhida int) {
	personagem, coordenadas, armasEscolhidas, bonusArma1, danoArma1, bonusArma2, danoArma2, err := LogicChar.GerarInfoFichaSpecific(racaEscolhida, classeEscolhida, idadeEscolhida)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Carregar dados de equipamentos
	equipamentos, err := DataChar.CarregarEquipamentos("assets/equipamentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar equipamentos:", err)
		return
	}

	// Carregar imagem da ficha
	imagem, err := imaging.Open("assets/Ficha.jpg")
	if err != nil {
		fmt.Println("Erro ao carregar imagem da ficha:", err)
		return
	}

	// Converter a imagem para RGBA
	bounds := imagem.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, imagem, bounds.Min, draw.Src)

	// Adicionar todas as informações do personagem na ficha
	LogicChar.AdicionarTodasInformacoesFicha(imgRGBA, personagem, coordenadas, armasEscolhidas, &equipamentos, bonusArma1, danoArma1, bonusArma2, danoArma2)

	// Salvar a imagem da ficha preenchida
	err = imaging.Save(imgRGBA, "FichaPreenchida.jpg")
	if err != nil {
		fmt.Println("Erro ao salvar a imagem da ficha preenchida:", err)
		return
	}

	fmt.Println("Ficha do personagem gerada com sucesso!")
}
