package character_sheet_imaging

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

func EscreverTextoEmVariasCoordenadasJOGADOR(imagem *image.RGBA, coordenadas []Coord) {
	const tamanhoDaFonte = 50
	const caminhoDaFonte = "assets/MedievalSharp.ttf"

	texto := "Escolha do JOGADOR"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := ConvertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

func EscreverTextoEmVariasCoordenadasMESTRE(imagem *image.RGBA, coordenadas []Coord) {
	const tamanhoDaFonte = 40
	const caminhoDaFonte = "assets/MedievalSharp.ttf"

	texto := "Escolha do MESTRE"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := ConvertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

// Função para escrever as informações da armadura
func EscreverArmadura(imagem *image.RGBA, texto string, coordenadas []Coord) {
	const tamanhoDaFonte = 55
	const caminhoDaFonte = "assets/MedievalSharp.ttf"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	for _, coord := range coordenadas {
		point := ConvertCoordToPoint(coord)
		fmt.Printf("Escrevendo texto '%s' em coordenadas (%d, %d)\n", texto, point.X, point.Y) // Adicionado para ver as coordenadas e o texto
		dc.DrawStringAnchored(texto, float64(point.X), float64(point.Y), 0, 1)
	}
}

// Função para adicionar texto na character_sheet_imaging com espacamento no Y
func AdicionarTextoNaFichaY(imagem *image.RGBA, texto interface{}, coord Coord, espacamento int, espacamentoX int) {
	const tamanhoDaFonte = 50
	const caminhoDaFonte = "assets/MedievalSharp.ttf"

	dc := gg.NewContextForRGBA(imagem)
	if err := dc.LoadFontFace(caminhoDaFonte, float64(tamanhoDaFonte)); err != nil {
		panic(err)
	}

	dc.SetColor(color.Black)

	point := ConvertCoordToPoint(coord)
	x, y := point.X+espacamentoX, point.Y+espacamento
	textoStr := fmt.Sprintf("%v", texto)
	fmt.Printf("Adicionando texto '%s' em coordenadas (%d, %d) com espaçamento %d\n", textoStr, x, y, espacamento) // Adicionado para ver as coordenadas e o texto
	dc.DrawStringAnchored(textoStr, float64(x), float64(y), 0, 1)
}

func AdicionarTextoNaFichaComDefault(imagem *image.RGBA, texto interface{}, coord Coord) {
	AdicionarTextoNaFichaY(imagem, texto, coord, 0, 0)
}
