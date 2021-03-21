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

	if InvalidPassThrough(orix, oriy, desx, desy, table) {
		return false
	}

	return true
}

func intAbs(num int) int {
	if num < 0 {
		return num*(-1)
	}
	return num
}

func invalidKingMovement(orix int, oriy int, desx int, desy int, table [8][8]Square) bool {
	return intAbs(orix-desx) > 1 || intAbs(oriy-desy) > 1
}

func invalidKnightMovement(orix int, oriy int, desx int, desy int) bool {
	if intAbs(orix - desx) == 2 && intAbs(oriy - desy) == 1 {
		return false
	}
	if intAbs(orix - desx) == 1 && intAbs(oriy - desy) == 2 {
		return false
	}
	return true
}

func invalidBishopMovement(orix int, oriy int, desx int, desy int) bool {
	if intAbs(orix - desx) == intAbs(oriy - desy) {
		return false
	}
	return true
}

func invalidRookMovement(orix int, oriy int, desx int, desy int) bool {
	if orix == desx || oriy == desy {
		return false
	}
	return true
}

func invalidQueenMovement(orix int, oriy int, desx int, desy int) bool {
	if queenIsBishop(orix, oriy, desx, desy) {
		return invalidBishopMovement(orix, oriy, desx, desy)
	}
	return invalidRookMovement(orix, oriy, desx, desy)
}

func rookPath(orix int, oriy int, desx int, desy int, table [8][8]Square) []Square {
	var path []Square
	if orix != desx {
		if orix < desx {
			for i := orix+1; i < desx; i++ {
				path = append(path, table[i][oriy])
			}
		} else {
			for i := orix-1; i > desx; i-- {
				path = append(path, table[i][oriy])
			}
		}
	} else {
		if oriy < desy {
			for i := oriy+1; i < desy; i++ {
				path = append(path, table[orix][i])
			}
		} else {
			for i := oriy-1; i > desy; i-- {
				path = append(path, table[orix][i])
			}
		}
	}
	return path
}

func bishopPath(orix int, oriy int, desx int, desy int, table [8][8]Square) []Square {
	var path []Square
	if orix < desx {
		if oriy < desy {
			for i := 1; i < (desx-orix); i++ {
				path = append(path, table[orix+i][oriy+i])
			}
		} else {
			for i := 1; i < (desx-orix); i++ {
				path = append(path, table[orix+i][oriy-i])
			}
		}
	} else {
		if oriy < desy {
			for i := 1; i < (orix-desx); i++ {
				path = append(path, table[orix-i][oriy+i])
			}
		} else {
			for i := 1; i < (orix-desx); i++ {
				path = append(path, table[orix-i][oriy-i])
			}
		}
	}
	return path
}

func queenPath(orix int, oriy int, desx int, desy int, table [8][8]Square) []Square {
	if queenIsBishop(orix, oriy, desx, desy) {
		return bishopPath(orix, oriy, desx, desy, table)
	}
	return rookPath(orix, oriy, desx, desy, table)
}

func queenIsBishop(orix int, oriy int, desx int, desy int) bool {
	if orix == desx || oriy == desy {
		return false
	}
	return true
}

func InvalidPassThrough(orix int, oriy int, desx int, desy int, table [8][8]Square) bool {

	switch table[orix][oriy].Piece {
	case 'R':
		for _, square := range rookPath(orix, oriy, desx, desy, table) {
			if square.Piece != 0 {
				return true
			}
		}
	case 'B':
		for _, square := range bishopPath(orix, oriy, desx, desy, table) {
			if square.Piece != 0 {
				return true
			}
		}
	case 'Q':
		for _, square := range queenPath(orix, oriy, desx, desy, table) {
			if square.Piece != 0 {
				return true
			}
		}
	}

	return false
}