package movements

import (
	"errors"
	"github.com/eriksencosta/tapsi-challenge/constants"
)

func findEatingMoveWhites(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string, row int, column int) (int, int) {
	// It is a little bit hard for the eye to track these operations in the if statements. The minus and plus operations could be
	// carried by the functions and/or extracted to helpers functions.
	if IsValidPosition(row - constants.DoubleMovement, column-constants.DoubleMovement) &&
		IsThereAnEnemy(board2D, player, row-constants.SingleMovement, column-constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row-constants.DoubleMovement, column-constants.DoubleMovement) {
		return row - constants.DoubleMovement, column - constants.DoubleMovement
	}

	if IsValidPosition(row-constants.DoubleMovement, column+constants.DoubleMovement) &&
		IsThereAnEnemy(board2D, player, row-constants.SingleMovement, column+constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row-constants.DoubleMovement, column+constants.DoubleMovement) {

		return row - constants.DoubleMovement, column + constants.DoubleMovement
	}

	return row, column
}

func findEatingMoveBlacks(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string, row int, column int) (int, int) {
	if IsValidPosition(row+constants.DoubleMovement, column-constants.DoubleMovement) &&
		IsThereAnEnemy(board2D, player, row+constants.SingleMovement, column-constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row+constants.DoubleMovement, column-constants.DoubleMovement) {
		return row + constants.DoubleMovement, column - constants.DoubleMovement
	}

	if IsValidPosition(row+constants.DoubleMovement, column+constants.DoubleMovement) &&
		IsThereAnEnemy(board2D, player, row+constants.SingleMovement, column+constants.SingleMovement) &&
		IsAnEmptyPlace(board2D, row+constants.DoubleMovement, column+constants.DoubleMovement) {
		return row + constants.DoubleMovement, column + constants.DoubleMovement
	}

	return row, column
}

func getEatingFunction(player string) func([constants.NumberOfRows][constants.NumberOfColumns]string, string, int, int) (int, int){
	if player == constants.WhitePlayer {
		return  findEatingMoveWhites
	} else {
		return findEatingMoveBlacks
	}
}

func FindEatingMove(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string) (Movement, error) {
	eatingFunction := getEatingFunction(player)

	for row := 0; row < constants.NumberOfRows; row++ {
		for column := 0; column < constants.NumberOfColumns; column++ {
			if player == board2D[row][column] {
				newRow, newColumn := eatingFunction(board2D, player, row, column)
				// nil would be a better return than the same row/column. It would be more explicit that no solution was found.
				if newRow != row && newColumn != column {
					return Movement{row, column,newRow, newColumn}, nil
				}
			}

		}
	}
	return Movement{}, errors.New("No hay posiciones donde se pueda comer una ficha")
}