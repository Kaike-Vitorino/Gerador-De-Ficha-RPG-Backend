package main

// Classe representa as informacoes de uma classe
type Classe struct {
	AtributoChave string
	Pericias      []string
	Equipamentos  Equipamento
	DadosRecurso  map[string]string
}

// Equipamento representa os equipamentos de uma classe
type Equipamento struct {
	Arma            []string
	Armadura        *string
	Itens           int
	ArtefatoMusical []string
	Cavalo          int
}

// PersonagemClasses armazena todas as classes e as infos dela
type PersonagemClasses struct {
	Classes         []string
	ClasseInfo      map[string]Classe
	TalentosClasses map[string][]string
}

// NewPersonagemClasses inicializa PersonagemClasses com valores padroes
func NewPersonagemClasses(armas1M map[string]map[string]string) *PersonagemClasses {
	armaduras1MLista := make([]string, 0, len(armas1M))
	for k := range armas1M {
		armaduras1MLista = append(armaduras1MLista, k)
	}

	return &PersonagemClasses{
		Classes: []string{"Caçador", "Druida", "Mago", "Rider", "Guerreiro", "Ladino", "Mascate", "Bardo"},
		ClasseInfo: map[string]Classe{
			"Caçador": {
				AtributoChave: "Agilidade",
				Pericias:      []string{"Furtividade", "Movimentação", "Pontaria", "Patrulha", "Sobrevivência"},
				Equipamentos:  Equipamento{Arma: []string{"Arco", "Funda"}, Armadura: nil, Itens: 2, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D8", "Flechas": "D10", "Prata": "D6"},
			},
			"Druida": {
				AtributoChave: "Inteligência",
				Pericias:      []string{"Resiliência", "Sobrevivência", "Discernimento", "Cura", "Adestramento"},
				Equipamentos:  Equipamento{Arma: []string{"Bastão", "Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D8", "Prata": "D6"},
			},
			"Mago": {
				AtributoChave: "Inteligência",
				Pericias:      []string{"Artesanato", "Artimanha", "Tradição", "Discernimento", "Manipulação"},
				Equipamentos:  Equipamento{Arma: []string{"Bastão", "Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D6", "Água": "D8", "Prata": "D8"},
			},
			"Rider": {
				AtributoChave: "Agilidade",
				Pericias:      []string{"Resiliência", "Luta", "Pontaria", "Sobrevivência", "Adestramento"},
				Equipamentos:  Equipamento{Arma: []string{"Lança Curta", "Machadinha", "Arco Curto", "Funda"}, Armadura: nil, Itens: 1, ArtefatoMusical: nil, Cavalo: 1},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D8", "Flechas": "D10", "Prata": "D6"},
			},
			"Guerreiro": {
				AtributoChave: "Força",
				Pericias:      []string{"Potência", "Resiliência", "Luta", "Movimentação"},
				Equipamentos:  Equipamento{Arma: armaduras1MLista, Armadura: StringPtr("Couro"), Itens: 1, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D6", "Prata": "D6"},
			},
			"Ladino": {
				AtributoChave: "Agilidade",
				Pericias:      []string{"Luta", "Furtividade", "Artimanha", "Movimentação", "Manipulação"},
				Equipamentos:  Equipamento{Arma: []string{"Adaga"}, Armadura: nil, Itens: 2, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D6", "Água": "D6", "Prata": "D10"},
			},
			"Mascate": {
				AtributoChave: "Empatia",
				Pericias:      []string{"Artesanato", "Artimanha", "Patrulha", "Discernimento", "Manipulação"},
				Equipamentos:  Equipamento{Arma: []string{"Faca"}, Armadura: nil, Itens: 3, ArtefatoMusical: nil, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D8", "Prata": "D12"},
			},
			"Bardo": {
				AtributoChave: "Empatia",
				Pericias:      []string{"Tradição", "Discernimento", "Manipulação", "Atuação", "Cura"},
				Equipamentos:  Equipamento{Arma: []string{"Faca"}, Armadura: nil, Itens: 1, ArtefatoMusical: []string{"Lira", "Flauta"}, Cavalo: 0},
				DadosRecurso:  map[string]string{"Comida": "D8", "Água": "D6", "Prata": "D8"},
			},
		},
		TalentosClasses: map[string][]string{
			"Caçador":   {"Caminho da Fera", "Caminho da Flecha", "Caminho da Floresta"},
			"Druida":    {"Caminho da Cura", "Caminho da Visão", "Caminho do Metamorfo"},
			"Mago":      {"Caminho da Morte", "Caminho das Rochas", "Caminho do Sangue", "Caminho dos Símbolos"},
			"Rider":     {"Caminho das Planícies", "Caminho do Cavaleiro", "Caminho do Companheiro"},
			"Guerreiro": {"Caminho da Lâmina", "Caminho do Escudo", "Caminho do Inimigo"},
			"Ladino":    {"Caminho do Matador", "Caminho do Rosto", "Caminho do Veneno"},
			"Mascate":   {"Caminho da Mentira", "Caminho das Muitas Coisas", "Caminho do Ouro"},
			"Bardo":     {"Caminho da Canção", "Caminho do Grito de Guerra", "Caminho do Hino"},
		},
	}
}

// StringPtr retorna um ponteiro para o valor da string que a gente precise.(usado para armadura)
func StringPtr(s string) *string {
	return &s
}
