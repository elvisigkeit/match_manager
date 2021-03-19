package data

func CheckValidMovement(orix int, oriy int, desx int, desy int, table [8][8]Square) bool {
	if orix < 0 || oriy < 0 || desx < 0 || desy < 0 ||
		orix > 7 || oriy > 7 || desx > 7 || desy > 7 {
		return false
	}

	if table[orix][oriy].Piece == 0 {
		return false
	}

	if table[orix][oriy].Color == table[desx][desy].Color {
		return false
	}

	return true
}