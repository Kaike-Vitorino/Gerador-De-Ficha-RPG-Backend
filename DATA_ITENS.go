package main

// Arma representa as caracteristicas de uma arma
type Arma struct {
	Bonus string
	Dano  string
}

// Escudo representa as caracteristicas de um escudo
type Escudo struct {
	Bonus string
}

// Armadura representa as caracteristicas de uma armadura
type Armadura struct {
	ValorDeArmadura string
	ParteDoCorpo    string
}

// Equipamentos contem todas as opcoes de equipamento do personagem
type Equipamentos struct {
	Items                    []string
	ArmaEscolhida            string
	ArtefatoMusicalEscolhido string
	ItensComercio            []string
	Armas1M                  map[string]Arma
	Armas2M                  map[string]Arma
	ListaArmas               map[string]Arma
	ArmasDistancia1M         map[string]Arma
	ArmasDistancia2M         map[string]Arma
	ListaArmasADistancia     map[string]Arma
	ListaArmasFinal          map[string]Arma
	ListaEscudos             map[string]Escudo
	ListaArmaduras           map[string]Armadura
}

// NewEquipamentos inicializando a struct Equipamentos com os dados possiveis
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
		Armas1M: map[string]Arma{
			"Faca":               {Bonus: "+1", Dano: "1"},
			"Adaga":              {Bonus: "+1", Dano: "1"},
			"Facão":              {Bonus: "+2", Dano: "2"},
			"Espada Curta":       {Bonus: "+2", Dano: "1"},
			"Espada Larga":       {Bonus: "+2", Dano: "2"},
			"Espada Longa":       {Bonus: "+2", Dano: "2"},
			"Cimitarra":          {Bonus: "+1", Dano: "2"},
			"Machadinha":         {Bonus: "+2", Dano: "2"},
			"Machado de Batalha": {Bonus: "+2", Dano: "2"},
			"Maça":               {Bonus: "+2", Dano: "1"},
			"Maça Estrela":       {Bonus: "+2", Dano: "2"},
			"Martelo de Guerra":  {Bonus: "+2", Dano: "2"},
			"Mangual":            {Bonus: "+1", Dano: "2"},
			"Porrete":            {Bonus: "+1", Dano: "1"},
			"Lança Curta":        {Bonus: "+1", Dano: "1"},
			"Tridente":           {Bonus: "+1", Dano: "2"},
		},
		Armas2M: map[string]Arma{
			"Montante":             {Bonus: "+3", Dano: "3"},
			"Machado de Duas Mãos": {Bonus: "+2", Dano: "3"},
			"Clava":                {Bonus: "+1", Dano: "2"},
			"Malho":                {Bonus: "+2", Dano: "3"},
			"Bastão":               {Bonus: "+1", Dano: "1"},
			"Lança Longa":          {Bonus: "+2", Dano: "1"},
			"Pique":                {Bonus: "+2", Dano: "2"},
			"Alabarda":             {Bonus: "+2", Dano: "2"},
		},
		ArmasDistancia1M: map[string]Arma{
			"Pedra":                {Bonus: "+0", Dano: "1"},
			"Faca de Arremesso":    {Bonus: "+1", Dano: "1"},
			"Machado de Arremesso": {Bonus: "+1", Dano: "2"},
			"Azagaia":              {Bonus: "+2", Dano: "1"},
			"Funda":                {Bonus: "+1", Dano: "1"},
		},
		ArmasDistancia2M: map[string]Arma{
			"Arco Curto":   {Bonus: "+2", Dano: "1"},
			"Arco Longo":   {Bonus: "+2", Dano: "1"},
			"Besta Leve":   {Bonus: "+1", Dano: "2"},
			"Besta Pesada": {Bonus: "+1", Dano: "3"},
		},
		ListaEscudos: map[string]Escudo{
			"Escudo Pequeno": {Bonus: "+1"},
			"Escudo Grande":  {Bonus: "+2"},
		},
		ListaArmaduras: map[string]Armadura{
			"Couro":                 {ValorDeArmadura: "2", ParteDoCorpo: "Corpo"},
			"Couro Batido":          {ValorDeArmadura: "3", ParteDoCorpo: "Corpo"},
			"Cota de Malha":         {ValorDeArmadura: "6", ParteDoCorpo: "Corpo"},
			"Armadura de Placas":    {ValorDeArmadura: "8", ParteDoCorpo: "Corpo"},
			"Capuz de Couro Batido": {ValorDeArmadura: "1", ParteDoCorpo: "Cabeça"},
			"Elmo Aberto":           {ValorDeArmadura: "2", ParteDoCorpo: "Cabeça"},
			"Elmo Fechado":          {ValorDeArmadura: "3", ParteDoCorpo: "Cabeça"},
			"Grande Elmo":           {ValorDeArmadura: "4", ParteDoCorpo: "Cabeça"},
		},
	}

	equipamentos.ListaArmas = mergeArmas(equipamentos.Armas1M, equipamentos.Armas2M)
	equipamentos.ListaArmasADistancia = mergeArmas(equipamentos.ArmasDistancia1M, equipamentos.ArmasDistancia2M)
	equipamentos.ListaArmasFinal = mergeArmas(equipamentos.ListaArmasADistancia, equipamentos.ListaArmas)

	return equipamentos
}

// mergeArmas junta dois mapas de arma
func mergeArmas(map1, map2 map[string]Arma) map[string]Arma {
	merged := make(map[string]Arma)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}
