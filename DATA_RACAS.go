package main

// Raca representa as informacoes de uma raca no sistema
type Raca struct {
	AtributoChave     string
	TalentoAscendente string
	ProfissoesTipicas []string
}

// IdadeRaca representa as faixas etarias para uma raca
type IdadeRaca struct {
	Jovem  *[2]int
	Adulto [2]int
	Idoso  *[2]int
}

// PersonagemRacas armazena todas as racas e suas infos
type PersonagemRacas struct {
	Racas      []string
	RacasInfo  map[string]Raca
	IdadeRacas map[string]IdadeRaca
}

// NewPersonagemRacas inicializa PersonagemRacas com valores padrao.
func NewPersonagemRacas(classes []string) *PersonagemRacas {
	return &PersonagemRacas{
		Racas: []string{"Humano", "Elfo", "Anão", "Goblin", "Orc", "Meio-Elfo", "Halfling", "Lupino"},
		RacasInfo: map[string]Raca{
			"Humano":    {AtributoChave: "Empatia", TalentoAscendente: "Adaptive", ProfissoesTipicas: classes},
			"Meio-Elfo": {AtributoChave: "Inteligência", TalentoAscendente: "Poder mental", ProfissoesTipicas: []string{"Druida", "Mago", "Ladino"}},
			"Anão":      {AtributoChave: "Força", TalentoAscendente: "Bravura Indômita", ProfissoesTipicas: []string{"Guerreiro", "Mascate", "Bardo"}},
			"Halfling":  {AtributoChave: "Empatia", TalentoAscendente: "Difícil de Acertar", ProfissoesTipicas: []string{"Bardo", "Mascate", "Ladino"}},
			"Lupino":    {AtributoChave: "Agilidade", TalentoAscendente: "Instinto de Caça", ProfissoesTipicas: []string{"Caçador", "Druida", "Guerreiro"}},
			"Orc":       {AtributoChave: "Força", TalentoAscendente: "Imbatível", ProfissoesTipicas: []string{"Caçador", "Guerreiro", "Ladino"}},
			"Goblin":    {AtributoChave: "Agilidade", TalentoAscendente: "Noturno", ProfissoesTipicas: []string{"Caçador", "Rider", "Ladino"}},
			"Elfo":      {AtributoChave: "Agilidade", TalentoAscendente: "Paz Interior", ProfissoesTipicas: []string{"Caçador", "Druida", "Bardo"}},
		},
		IdadeRacas: map[string]IdadeRaca{
			"Humano":    {Jovem: &[2]int{16, 25}, Adulto: [2]int{26, 50}, Idoso: &[2]int{51, 80}},
			"Meio-Elfo": {Jovem: &[2]int{16, 30}, Adulto: [2]int{31, 100}, Idoso: &[2]int{101, 180}},
			"Anão":      {Jovem: &[2]int{20, 40}, Adulto: [2]int{41, 80}, Idoso: &[2]int{81, 121}},
			"Halfling":  {Jovem: &[2]int{16, 25}, Adulto: [2]int{26, 60}, Idoso: &[2]int{61, 98}},
			"Lupino":    {Jovem: &[2]int{13, 20}, Adulto: [2]int{21, 40}, Idoso: &[2]int{41, 65}},
			"Orc":       {Jovem: &[2]int{13, 20}, Adulto: [2]int{21, 45}, Idoso: &[2]int{46, 70}},
			"Goblin":    {Jovem: &[2]int{16, 25}, Adulto: [2]int{26, 60}, Idoso: &[2]int{61, 95}},
			"Elfo":      {Adulto: [2]int{26, 1000}}, // Elfos sempre sao adultos
		},
	}
}
