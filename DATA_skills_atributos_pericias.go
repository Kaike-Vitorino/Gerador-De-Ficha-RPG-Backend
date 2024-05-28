package main

import "fmt"

// Atributos represents the attributes of a character.
type Atributos struct {
	Forca        []int
	Agilidade    []int
	Inteligencia []int
	Empatia      []int
}

// Pericias represents the skills and their corresponding attributes.
type Pericias map[string]string

// Talentos represents the talents and the classes that can have them.
type Talentos map[string][]string

// PersonagemStatus represents the character's status, including attributes, skills, and talents.
type PersonagemStatus struct {
	Atributos Atributos
	Pericias  Pericias
	Talentos  Talentos
}

// NewPersonagemStatus initializes the PersonagemStatus with default values.
func NewPersonagemStatus() *PersonagemStatus {
	return &PersonagemStatus{
		Atributos: Atributos{
			Forca:        []int{0, 5},
			Agilidade:    []int{0, 5},
			Inteligencia: []int{0, 5},
			Empatia:      []int{0, 5},
		},
		Pericias: Pericias{
			"Potência": "Força", "Resiliência": "Força", "Luta": "Força", "Artesanato": "Força",
			"Furtividade": "Agilidade", "Artimanha": "Agilidade", "Movimentação": "Agilidade", "Pontaria": "Agilidade",
			"Patrulha": "Inteligência", "Tradição": "Inteligência", "Sobrevivência": "Inteligência", "Discernimento": "Inteligência",
			"Manipulação": "Empatia", "Atuação": "Empatia", "Cura": "Empatia", "Adestramento": "Empatia",
		},
		Talentos: Talentos{
			"Caçador": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Aquartelador", "Arremessador", "Cozinheiro", "Defensor",
				"Atirador preciso", "Atirador veloz", "Desbravador", "Herbalista", "Marinheiro", "Mestre da Caçada",
				"Pés ligeiros", "Pescador", "Sexto sentido",
			},
			"Druida": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Andarilho", "Aquartelador", "Cozinheiro", "Defensor",
				"Mestre em facas", "Desbravador", "Destemido", "Herbalista", "Incorruptível", "Marinheiro",
				"Mestre da Caçada", "Pescador",
			},
			"Mago": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Arremessador", "Arrombador", "Artífice de arcos", "Cozinheiro",
				"Curtidor", "Destemido", "Envenenador", "Ferreiro", "Incorruptível", "Língua afiada", "Mestre em facas",
			},
			"Rider": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Mestre em lanças", "Sangue frio", "Saque rapido", "Sortudo", "Andarilho", "Aquartelador", "Desbravador",
				"Herbalista", "Atirador preciso", "Atirador veloz", "Investida", "Lutador", "Marinheiro", "Mestre da Caçada",
				"Mestre da montaria", "Pés ligeiros", "Pescador", "Mestre em machados",
			},
			"Guerreiro": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Ameaçador", "Andarilho", "Aquartelador", "Defensor", "Investida",
				"Lutador", "Atirador preciso", "Atirador veloz", "Mestre em lanças", "Mestre em machados", "Mestre em espadas",
				"Mestre em facas", "Mestre em martelos",
			},
			"Ladino": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Arrombador", "Defensor", "Investida", "Língua afiada",
				"Lutador", "Pés ligeiros", "Mestre em facas",
			},
			"Mascate": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Arrombador", "Artífice de arcos", "Curtidor", "Destemido",
				"Envenenador", "Ferreiro", "Incorruptível", "Língua afiada", "Sexto sentido", "Mestre em facas",
			},
			"Bardo": {
				"Ambidestria", "Berseker", "Carrasco", "Empunhadura Firme", "Rapido como um raio", "Ruína dos Dragões",
				"Sangue frio", "Saque rapido", "Sortudo", "Artífice de arcos", "Cozinheiro", "Curtidor", "Defensor",
				"Desbravador", "Destemido", "Herbalista", "Incorruptível", "Língua afiada", "Mestre em facas",
			},
		},
	}
}

func main() {
	status := NewPersonagemStatus()
	fmt.Println(status)
}
