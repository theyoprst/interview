package solution

const (
	Empty byte = '.'
	Wall byte = 'X'
	GuardUp byte = '^'
	GuardDown byte = 'v'
	GuardLeft byte = '<'
	GuardRight byte = '>'
	Assassin byte = 'A'
	Visited byte = '#'
	SeenByGuard byte = 's'
)

func buildWall(board []string) [][]byte {
	var wallRow []byte
	cols := len(board[0])
	for i := 0; i < cols+2; i++ {
		wallRow = append(wallRow, Wall)
	}
	var mutableBoard [][]byte
	mutableBoard = append(mutableBoard, wallRow)
	for _, rowStr := range board {
		var row []byte
		row = append(row, Wall)
		for _, ch := range rowStr {
			row = append(row, byte(ch))
		}
		row = append(row, Wall)
		mutableBoard = append(mutableBoard, row)
	}
	mutableBoard = append(mutableBoard, wallRow)
	return mutableBoard
}

func fillLineTillObstacle(board [][]byte, i, j, di, dj int) {
	for {
		i += di
		j += dj
		switch board[i][j] {
		case Wall, GuardLeft, GuardRight, GuardUp, GuardDown:
			return
		case Empty, Assassin, SeenByGuard:
			board[i][j] = SeenByGuard
		default:
			panic("unreached")
		}
	}
}

var dirs = []struct{
	di int
	dj int
} {
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func dfs(board [][]byte, i, j int) {
	for _, dir := range dirs {
		ni := i + dir.di
		nj := j + dir.dj
		if board[ni][nj] == Empty {
			board[ni][nj] = Visited
			dfs(board, ni, nj)
		}
	}
}

func CanAssassinEscape(B []string) bool {
	// 1. Build a wall in order to get rid of out-of-board checks
	// Also convert it into the mutable structure
	// O(N*M)
	board := buildWall(B)

	// Convert guard lines to walls
	// O(N*M)
	for i, row := range board {
		for j, cell := range row {
			switch cell {
			case GuardLeft:
				fillLineTillObstacle(board, i, j, 0, -1)
			case GuardRight:
				fillLineTillObstacle(board, i, j, 0, 1)
			case GuardUp:
				fillLineTillObstacle(board, i, j, -1, 0)
			case GuardDown:
				fillLineTillObstacle(board, i, j, 1, 0)
			}
		}
	}

	assassinRow := -1
	assassinCol := -1
	for i, row := range board {
		for j, cell := range row {
			if cell == Assassin {
				assassinRow = i
				assassinCol = j
			}
		}
	}
	if assassinCol == -1 || assassinRow == -1 {
		return false
	}
	board[assassinRow][assassinCol] = Visited
	dfs(board, assassinRow, assassinCol)
	return board[len(board)-2][len(board[0])-2] == Visited
}
