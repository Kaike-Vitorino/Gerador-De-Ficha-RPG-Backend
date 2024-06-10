package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

func escreverTextoEmVariasCoordenadasJOGADOR(imagem *image.RGBA, coordenadas []Coord) {
	const tamanhoDaFonte = 50
	const caminhoDaFonte = "data/MedievalSharp.ttf"

	texto := "Escolha do JOGADOR"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := convertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

func escreverTextoEmVariasCoordenadasMESTRE(imagem *image.RGBA, coordenadas []Coord) {
	const tamanhoDaFonte = 40
	const caminhoDaFonte = "data/MedievalSharp.ttf"

	texto := "Escolha do MESTRE"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := convertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

// Função para escrever as informações da armadura
func escreverArmadura(imagem *image.RGBA, texto string, coordenadas []Coord) {
	const tamanhoDaFonte = 55
	const caminhoDaFonte = "data/MedievalSharp.ttf"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := convertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

// Função para adicionar texto na ficha
func adicionarTextoNaFicha(imagem *image.RGBA, texto interface{}, coord Coord, espacamento int) {
	const tamanhoDaFonte = 50
	const caminhoDaFonte = "data/MedievalSharp.ttf"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	point := convertCoordToPoint(coord)
	x, y := point.X, point.Y+espacamento
	textoStr := fmt.Sprintf("%v", texto)
	fmt.Printf("Adicionando texto '%s' em coordenadas (%d, %d) com espaçamento %d\n", textoStr, x, y, espacamento) // Adicionado para ver as coordenadas e o texto
	dc.DrawStringAnchored(textoStr, float64(x), float64(y), 0, 1)
}

func adicionarTextoNaFichaComDefault(imagem *image.RGBA, texto interface{}, coord Coord) {
	adicionarTextoNaFicha(imagem, texto, coord, 0)
}
