package main

// Equipamentos holds the equipment for a character.
type Equipamentos struct {
	Items                    []string
	ArmaEscolhida            string
	ArtefatoMusicalEscolhido string
	ItensComercio            []string
	Armas1M                  map[string]map[string]string
	Armas2M                  map[string]map[string]string
	ListaArmas               map[string]map[string]string
	ArmasDistancia1M         map[string]map[string]string
	ArmasDistancia2M         map[string]map[string]string
	ListaArmasADistancia     map[string]map[string]string
	ListaArmasFinal          map[string]map[string]string
	ListaEscudos             map[string]map[string]string
	ListaArmaduras           map[string]map[string]string
}

// NewEquipamentos initializes the Equipamentos struct with the appropriate data.
func NewEquipamentos() *Equipamentos {
	equipamentos := &Equipamentos{
		Items: []string{},
		ItensComercio: []string{
			"Flechas, Ponta de Ferro", "Flechas, Ponta de Madeira", "Aljava", "Arpéu",
			"Corda, 10 Metros", "Vela de Sebo", "Lamparina a Óleo", "Lanterna", "Tochas",
			"Saco", "Mochila", "Cantil", "Bandagens", "Óleo de Lâmpada", "Pena e Tinta",
			"Pergaminho", "Cobertor", "Cobertor de Peles", "Pederneira", "Gazuas",
			"Cozinha de Campo", "Caldeirão", "Cálice de Metal", "Caneca", "Prato de Metal",
			"Faca de Cozinha", "Colher", "Armadilha de Urso", "Armadilha de Laço", "Barril",
			"Jarra", "Tenda Pequena", "Tenda Grande", "Anzol e Linha", "Rede de Pesca",
			"Lupa", "Símbolo Sagrado", "Giz", "Mapa", "Luneta", "Bola de Cristal",
			"Ampulheta", "Balança", "Flauta", "Berrante", "Lira", "Harpa", "Tambor",
			"Veneno Letal/Antídoto", "Veneno Paralisante/Antídoto", "Veneno Sonífero/Antídoto", "Veneno Alucinógeno/Antídoto",
		},
		Armas1M: map[string]map[string]string{
			"Faca":               {"Bonus": "+1", "Dano": "1"},
			"Adaga":              {"Bonus": "+1", "Dano": "1"},
			"Facão":              {"Bonus": "+2", "Dano": "2"},
			"Espada Curta":       {"Bonus": "+2", "Dano": "1"},
			"Espada Larga":       {"Bonus": "+2", "Dano": "2"},
			"Espada Longa":       {"Bonus": "+2", "Dano": "2"},
			"Cimitarra":          {"Bonus": "+1", "Dano": "2"},
			"Machadinha":         {"Bonus": "+2", "Dano": "2"},
			"Machado de Batalha": {"Bonus": "+2", "Dano": "2"},
			"Maça":               {"Bonus": "+2", "Dano": "1"},
			"Maça Estrela":       {"Bonus": "+2", "Dano": "2"},
			"Martelo de Guerra":  {"Bonus": "+2", "Dano": "2"},
			"Mangual":            {"Bonus": "+1", "Dano": "2"},
			"Porrete":            {"Bonus": "+1", "Dano": "1"},
			"Lança Curta":        {"Bonus": "+1", "Dano": "1"},
			"Tridente":           {"Bonus": "+1", "Dano": "2"},
		},
		Armas2M: map[string]map[string]string{
			"Montante":             {"Bonus": "+3", "Dano": "3"},
			"Machado de Duas Mãos": {"Bonus": "+2", "Dano": "3"},
			"Clava":                {"Bonus": "+1", "Dano": "2"},
			"Malho":                {"Bonus": "+2", "Dano": "3"},
			"Bastão":               {"Bonus": "+1", "Dano": "1"},
			"Lança Longa":          {"Bonus": "+2", "Dano": "1"},
			"Pique":                {"Bonus": "+2", "Dano": "2"},
			"Alabarda":             {"Bonus": "+2", "Dano": "2"},
		},
		ArmasDistancia1M: map[string]map[string]string{
			"Pedra":                {"Bonus": "+0", "Dano": "1"},
			"Faca de Arremesso":    {"Bonus": "+1", "Dano": "1"},
			"Machado de Arremesso": {"Bonus": "+1", "Dano": "2"},
			"Azagaia":              {"Bonus": "+2", "Dano": "1"},
			"Funda":                {"Bonus": "+1", "Dano": "1"},
		},
		ArmasDistancia2M: map[string]map[string]string{
			"Arco Curto":   {"Bonus": "+2", "Dano": "1"},
			"Arco Longo":   {"Bonus": "+2", "Dano": "1"},
			"Besta Leve":   {"Bonus": "+1", "Dano": "2"},
			"Besta Pesada": {"Bonus": "+1", "Dano": "3"},
		},
		ListaEscudos: map[string]map[string]string{
			"Escudo Pequeno": {"Bonus": "+1"},
			"Escudo Grande":  {"Bonus": "+2"},
		},
		ListaArmaduras: map[string]map[string]string{
			"Couro":                 {"Valor de Armadura": "2", "Parte do Corpo": "Corpo"},
			"Couro Batido":          {"Valor de Armadura": "3", "Parte do Corpo": "Corpo"},
			"Cota de Malha":         {"Valor de Armadura": "6", "Parte do Corpo": "Corpo"},
			"Armadura de Placas":    {"Valor de Armadura": "8", "Parte do Corpo": "Corpo"},
			"Capuz de Couro Batido": {"Valor de Armadura": "1", "Parte do Corpo": "Cabeça"},
			"Elmo Aberto":           {"Valor de Armadura": "2", "Parte do Corpo": "Cabeça"},
			"Elmo Fechado":          {"Valor de Armadura": "3", "Parte do Corpo": "Cabeça"},
			"Grande Elmo":           {"Valor de Armadura": "4", "Parte do Corpo": "Cabeça"},
		},
	}

	equipamentos.ListaArmas = mergeMaps(equipamentos.Armas1M, equipamentos.Armas2M)
	equipamentos.ListaArmasADistancia = mergeMaps(equipamentos.ArmasDistancia1M, equipamentos.ArmasDistancia2M)
	equipamentos.ListaArmasFinal = mergeMaps(equipamentos.ListaArmasADistancia, equipamentos.ListaArmas)

	return equipamentos
}

// mergeMaps merges two maps of string to map[string]string.
func mergeMaps(map1, map2 map[string]map[string]string) map[string]map[string]string {
	merged := make(map[string]map[string]string)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}
