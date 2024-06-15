package character_sheet_imaging

import (
	"fmt"
	"image"
	DataChar "psBackKG/model/character_data"
)

type Coord [2]int

// Coordenadas tem as coordenadas do elementos da character_sheet_imaging do sistema RPG
type Coordenadas struct {
	CoordenadasSimples     map[string]Coord  `json:"COORDENADAS"`
	ArmorCoordenadas       []Coord           `json:"COORDENADAS_ARMADURA"`
	PlayerCoordenadas      []Coord           `json:"COORDENADAS_USER"`
	MestreCoordenadas      []Coord           `json:"COORDENADAS_MESTRE"`
	CoordenadasSimplesPag2 map[string]Coord  `json:"COORDENADAS_PAG2"`
	KeyMap                 map[string]string `json:"MAPA_CHAVES"`
}

// NewCoordinates inicializa a estrutura Coordenadas com os valores do arquivo JSON
func NewCoordinates() (*Coordenadas, error) {
	var coords Coordenadas
	err := DataChar.ReadJSON("assets/coordenadas.json", &coords)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar coordenadas: %v", err)
	}
	return &coords, nil
}

func ConvertCoordToPoint(c Coord) image.Point {
	return image.Point{X: c[0], Y: c[1] + 10}
}
