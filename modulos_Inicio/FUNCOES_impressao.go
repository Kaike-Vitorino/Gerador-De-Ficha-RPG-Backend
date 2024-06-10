package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"os"

	"github.com/fogleman/gg"
)

// Função para escrever o mesmo texto nas várias coordenadas (Jogador)
func escreverTextoEmVariasCoordenadasJOGADOR(imagem *image.RGBA, texto string, coordenadas []Coord) {
	const tamanhoDaFonte = 50
	const caminhoDaFonte = "data/MedievalSharp.ttf" // Caminho para a fonte TrueType

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		x, y := coord[0], coord[1]
		dc.DrawStringAnchored(texto, float64(x), float64(y), 0, 1)
	}
}

// Função para escrever o mesmo texto nas várias coordenadas (Mestre)
func escreverTextoEmVariasCoordenadasMESTRE(imagem *image.RGBA, texto string, coordenadas []Coord) {
	const tamanhoDaFonte = 30
	const caminhoDaFonte = "data/MedievalSharp.ttf" // Caminho para a fonte TrueType

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		x, y := coord[0], coord[1]
		dc.DrawStringAnchored(texto, float64(x), float64(y), 0, 1)
	}
}

// Função para escrever as informações da armadura
func escreverArmadura(imagem *image.RGBA, texto string, coordenadas []Coord) {
	const tamanhoDaFonte = 55
	const caminhoDaFonte = "data/MedievalSharp.ttf" // Caminho para a fonte TrueType

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		x, y := coord[0], coord[1]
		dc.DrawStringAnchored(texto, float64(x), float64(y), 0, 1)
	}
}

// Função para adicionar texto na ficha
func adicionarTextoNaFicha(imagem *image.RGBA, texto interface{}, coord Coord, espacamento int) {
	tamanhoDaFonte := 50
	fonte := loadFont("data/MedievalSharp.ttf", float64(tamanhoDaFonte))
	x, y := coord[0], coord[1]+espacamento
	desenho := image.NewRGBA(imagem.Bounds())
	drawer := &font.Drawer{
		Dst:  desenho,
		Src:  image.NewUniform(color.Black),
		Face: fonte,
		Dot:  fixed.P(x, y),
	}
	textoStr := fmt.Sprintf("%v", texto)
	drawer.DrawString(textoStr)
	imaging.Overlay(imagem, desenho, image.Point{}, 1.0)
}

// Função para carregar a fonte
func loadFont(path string, size float64) font.Face {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Erro ao carregar a fonte:", err)
		os.Exit(1)
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		fmt.Println("Erro ao parsear a fonte:", err)
		os.Exit(1)
	}
	return truetype.NewFace(f, &truetype.Options{Size: size})
}

func adicionarTextoNaFichaComDefault(imagem *image.RGBA, texto interface{}, coord Coord) {
	adicionarTextoNaFicha(imagem, texto, coord, 0)
}
