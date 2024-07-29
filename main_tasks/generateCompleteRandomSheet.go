package main_tasks

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	LogicChar "psBackKG/model/character_core"
	DataChar "psBackKG/model/character_data"
)

func GerarFichaRandomCompleta() {
	personagem, coordenadas, armasEscolhidas, bonusArma1, danoArma1, bonusArma2, danoArma2, err := LogicChar.GerarInfoFichaRandom()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Carregar dados de character_sheet_imaging
	equipamentos, err := DataChar.CarregarEquipamentos("assets/equipamentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar character_sheet_imaging:", err)
		return
	}

	// Carregar imagem da character_sheet_imaging
	imagem, err := imaging.Open("assets/FichaM.png")
	if err != nil {
		fmt.Println("Erro ao carregar imagem da character_sheet_imaging:", err)
		return
	}

	// Converter a imagem para RGBA
	bounds := imagem.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, imagem, bounds.Min, draw.Src)

	// Adicionar todas as informações do character_logic na character_sheet_imaging
	LogicChar.AdicionarTodasInformacoesFicha(imgRGBA, personagem, coordenadas, armasEscolhidas, &equipamentos, bonusArma1, danoArma1, bonusArma2, danoArma2)

	// Salvar a imagem da character_sheet_imaging preenchida
	err = imaging.Save(imgRGBA, "FichaPreenchida.jpg")
	if err != nil {
		fmt.Println("Erro ao salvar a imagem da character_sheet_imaging preenchida:", err)
		return
	}

	fmt.Println("Ficha do character_logic gerada com sucesso!")
}
