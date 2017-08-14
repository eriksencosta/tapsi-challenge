package movements

import (
	"errors"
	"github.com/eriksencosta/tapsi-challenge/constants"
)

func findSimpleMoveWhites(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string, row int, column int) (int, int) {
	if IsValidPosition(row - constants.SingleMovement, column - constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row - constants.SingleMovement, column - constants.SingleMovement) {
		return row - constants.SingleMovement, column - constants.SingleMovement
	}

	if IsValidPosition(row - constants.SingleMovement, column + constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row - constants.SingleMovement, column + constants.SingleMovement) {
		return row - constants.SingleMovement, column + constants.SingleMovement
	}

	return row, column
}

func findSimpleMoveBlacks(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string, row int, column int) (int, int) {
	if IsValidPosition(row + constants.SingleMovement, column - constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row + constants.SingleMovement, column - constants.SingleMovement) {
		return row + constants.SingleMovement, column - constants.SingleMovement
	}

	if IsValidPosition(row + constants.SingleMovement, column + constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row + constants.SingleMovement, column + constants.SingleMovement) {
		return row + constants.SingleMovement, column + constants.SingleMovement
	}

	return row, column
}

func getSimpleMoveFunction(player string) func([constants.NumberOfRows][constants.NumberOfColumns]string, string, int, int) (int, int){
	if player == constants.WhitePlayer {
		return  findSimpleMoveWhites
	} else {
		return findSimpleMoveBlacks
	}
}

func FindSimpleMove(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string) (Movement, error) {
	simpleMoveFunction := getSimpleMoveFunction(player)

	for row := 0; row < constants.NumberOfRows; row++ {
		for column := 0; column < constants.NumberOfColumns; column++ {
			if player == board2D[row][column] {
				newRow, newColumn := simpleMoveFunction(board2D, player, row, column)
				if newRow != row && newColumn != column {
					return Movement{row, column,newRow, newColumn}, nil
				}
			}

		}
	}
	return Movement{}, errors.New("No hay posiciones para realizar un movimiento valido")

}