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

	switch table[orix][oriy].Piece {
	case 'R':
		if invalidRookMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'N':
		if invalidKnightMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'B':
		if invalidBishopMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'Q':
		if invalidQueenMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'K':
		if invalidKingMovement (orix, oriy, desx, desy, table) {
			return false
		}
	}

	if invalidPassThrough (orix, oriy, desx, desy, table) {
		return false
	}

	return true
}

func invalidKingMovement(orix int, oriy int, desx int, desy int, table [8][8]Square) bool {
	return false
}

func invalidQueenMovement(orix int, oriy int, desx int, desy int) bool {
	return false
}

func invalidBishopMovement(orix int, oriy int, desx int, desy int) bool {
	return false
}

func invalidKnightMovement(orix int, oriy int, desx int, desy int) bool {
	return false
}

func invalidRookMovement(orix int, oriy int, desx int, desy int) bool {
	return false
}

func rookPath(orix int, oriy int, desx int, desy int, table [8][8]Square) []Square {
	var path []Square
	if orix != desx {
		if oriy < desy {
			for i := oriy+1; i < desy; i++ {
				path = append(path, table[orix][i])
			}
		} else {
			for i := oriy-1; i > desy; i-- {
				path = append(path, table[orix][i])
			}
		}
	} else {
		if orix < desx {
			for i := orix+1; i < desx; i++ {
				path = append(path, table[oriy][i])
			}
		} else {
			for i := orix-1; i > desx; i-- {
				path = append(path, table[oriy][i])
			}
		}
	}
	return path
}

func invalidPassThrough(orix int, oriy int, desx int, desy int, table [8][8]Square) bool {

	switch table[orix][oriy].Piece {
	case 'R':
		for _, square := range rookPath(orix, oriy, desx, desy, table) {
			if square.Piece != 0 {
				return true
			}
		}
	case 'B':
		if invalidBishopMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'Q':
		if invalidQueenMovement (orix, oriy, desx, desy) {
			return false
		}
	case 'K':
		if invalidKingMovement (orix, oriy, desx, desy, table) {
			return false
		}
	}

	return false
}