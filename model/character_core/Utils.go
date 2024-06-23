package character_core

import DataChar "psBackKG/model/character_data"

// Função auxiliar para obter as informações da arma
func obterInfoArma(arma string, equipamentos *DataChar.Equipamentos) DataChar.Arma {
	if infoArma, existe := equipamentos.Armas1M[arma]; existe {
		return infoArma
	}
	if infoArma, existe := equipamentos.Armas2M[arma]; existe {
		return infoArma
	}
	if infoArma, existe := equipamentos.ArmasDistancia1M[arma]; existe {
		return infoArma
	}
	if infoArma, existe := equipamentos.ArmasDistancia2M[arma]; existe {
		return infoArma
	}
	return DataChar.Arma{}
}
