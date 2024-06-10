package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
)

func main() {
	// Carregar dados de raças
	racas, err := NewPersonagemRacas("data/racas.json")
	if err != nil {
		fmt.Println("Erro ao carregar raças:", err)
		return
	}

	// Carregar dados de classes
	classes, err := NewPersonagemClasses("data/classes.json")

	if err != nil {
		fmt.Println("Erro ao carregar classes:", err)
		return
	}

	// Carregar dados de status
	status, err := NewPersonagemStatus("data/atributos.json", "data/pericias.json", "data/talentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar status:", err)
		return
	}

	// Carregar dados de equipamentos
	equipamentos, err := CarregarEquipamentos()
	if err != nil {
		fmt.Println("Erro ao carregar equipamentos:", err)
		return
	}

	// Carregar dados de talentos
	talentos, err := carregarTalentos("data/talentos.json")
	if err != nil {
		fmt.Println("Erro ao carregar talentos:", err)
		return
	}

	// Carregar dados de atributos
	atributosData, err := carregarAtributos("data/atributos.json")
	if err != nil {
		fmt.Println("Erro ao carregar dados de atributos:", err)
		return
	}

	// Carregar coordenadas
	coordenadas, err := NewCoordinates()
	if err != nil {
		fmt.Println("Erro ao carregar coordenadas:", err)
		return
	}

	// Definir quantidade de XP
	var pontosXP int
	//fmt.Print("Quantidade de XP: ")
	//fmt.Scan(&pontosXP)
	pontosXP = 150

	// Gerar personagem aleatório
	personagem, err := gerarPersonagemAleatorio(racas, classes, status, &talentos, atributosData, pontosXP)
	if err != nil {
		fmt.Println("Erro ao gerar personagem:", err)
		return
	}

	// Gerar armas para o personagem
	armasEscolhidas, err := GerarArma(personagem.Classe, classes.ClasseInfo, equipamentos)
	if err != nil {
		fmt.Println("Erro ao gerar armas:", err)
		return
	}

	// Carregar imagem da ficha
	imagem, err := imaging.Open("data/Ficha.jpg")
	if err != nil {
		fmt.Println("Erro ao carregar imagem da ficha:", err)
		return
	}

	// Converter a imagem para RGBA
	bounds := imagem.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, imagem, bounds.Min, draw.Src)

	// Adicionar todas as informações do personagem na ficha
	adicionarTodasInformacoesFicha(imgRGBA, personagem, coordenadas, armasEscolhidas, classes.ClasseInfo, equipamentos)

	// Salvar a imagem da ficha preenchida
	err = imaging.Save(imgRGBA, "FichaPreenchida.jpg")
	if err != nil {
		fmt.Println("Erro ao salvar a imagem da ficha preenchida:", err)
		return
	}

	fmt.Println("Ficha do personagem gerada com sucesso!")
}
