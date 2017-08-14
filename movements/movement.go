package movements

import (
	"github.com/eriksencosta/tapsi-challenge/constants"
	"fmt"
)

type Movement struct{
	initialRow int
	initialColumn int
	finalRow int
	finalColumn int
}

func Movement2DToString(movement Movement) string {
	return fmt.Sprintf("( %d,%d ) -> ( %d,%d )", movement.initialRow, movement.initialColumn, movement.finalRow, movement.finalColumn)
}

func MovementToString(movement Movement) string {
	initialPosition := movement.initialRow * constants.NumberOfRows + movement.initialColumn
	finalPosition := movement.finalRow * constants.NumberOfRows + movement.finalColumn
	return fmt.Sprintf("%d-%d", initialPosition, finalPosition)
}

func representBoardIn2D(board string) [constants.NumberOfRows][constants.NumberOfColumns]string{
	board2D := [constants.NumberOfRows][constants.NumberOfColumns]string{}
	for row := 0; row < constants.NumberOfRows; row++ {
		for column := 0; column < constants.NumberOfColumns; column++ {
			board2D[row][column] = string(board[ row*constants.NumberOfRows+column]);
		}
	}
	return board2D
}

func IsValidPosition(row int, column int) bool {
	return row < constants.NumberOfRows && row >= 0 && column < constants.NumberOfColumns && column >= 0
}

func IsThereAnEnemy(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, player string, row int, column int) bool {
	if player == constants.BlackPlayer {
		return board2D[row][column] == constants.WhitePlayer
	} else {
		return board2D[row][column] == constants.BlackPlayer
	}
}

func IsAnEmptyPlace(board2D [constants.NumberOfRows][constants.NumberOfColumns]string, row int, column int) bool {
	return board2D[row][column] == constants.EmptySquare
}



func MakeMovement(board string, player string) (Movement, error) {
	board2D := representBoardIn2D(board)
	newMove, error := FindEatingMove(board2D, player)

	if error == nil{
		return newMove, nil
	} else{
		return FindSimpleMove(board2D,player)
	}
}


