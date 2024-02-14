package utils

func CheckWin(cells []string) (string, bool) {
	var winner string
	var win bool
	if winner, win = checkRowWin(cells); win {
		return winner, win
	} else if winner, win = checkColumnWin(cells); win {
		return winner, win
	} else if winner, win = checkDiagonalWin(cells); win {
		return winner, win
	}
	return winner, win
}

func checkRowWin(cells []string) (string, bool) {
	for i := 0; i < 9; i += 3 {
		if cells[i] == cells[i+1] && cells[i] == cells[i+2] && cells[i] != "-" {
			winner := cells[i]
			return winner, true
		}
	}
	return "-", false
}

func checkColumnWin(cells []string) (string, bool) {
	for i := 0; i < 3; i++ {
		if cells[i] == cells[i+3] && cells[i] == cells[i+6] && cells[i] != "-" {
			winner := cells[i]
			return winner, true
		}
	}
	return "-", false
}

func checkDiagonalWin(cells []string) (string, bool) {
	if cells[0] == cells[4] && cells[0] == cells[8] && cells[0] != "-" {
		winner := cells[0]
		return winner, true
	} else if cells[2] == cells[4] && cells[2] == cells[6] && cells[2] != "-" {
		winner := cells[2]
		return winner, true
	}
	return "-", false
}
