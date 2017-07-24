package validations

import (
	"strings"
	"errors"
	"github.com/jfsanin/challenge/constants"
)

func isDone(board string) bool {
	whites := strings.Count(board, constants.WhitePlayer)
	blacks := strings.Count(board, constants.BlackPlayer)

	return !(whites > 0 && blacks > 0)
}

func isEmpty(board string) bool {
	return strings.Count(board, constants.EmptySquare) == constants.NumberOfSquares
}

func hasCorrectLength(board string) bool{
	return len(board) == constants.NumberOfSquares
}

func hasValidPieces(board string) bool {
	whites := strings.Count(board, constants.WhitePlayer)
	blacks := strings.Count(board, constants.BlackPlayer)
	empties := strings.Count(board, constants.EmptySquare)

	return whites+blacks+empties == constants.NumberOfSquares
}

func hasValidStructure(board string) bool{
	for index := 0; index < len(board); index++ {
		row := int( index / 8)
		if string(board[index]) != constants.EmptySquare{
			if (row % 2 == 1 && index % 2 == 1) || (row % 2 == 0 && index % 2 == 0){
				return false
			}
		}
	}
	return true
}

func isValidPlayer(player string) bool {
	return player == constants.BlackPlayer || player == constants.WhitePlayer
}

func ValidateRequest(board *string, player *string) error{
	if board == nil{
		return errors.New("Se debe enviar el campo board")
	}
	if player == nil {
		return errors.New("Se debe enviar el campo player")
	}
	return nil
}


func Validate(board string, player string) error{
	if !hasCorrectLength(board){
		return errors.New("El string enviado debe tener exactamente 64 caracteres")
	} else if isEmpty(board) {
		return errors.New("El tablero esta vacio")
	}  else if !hasValidPieces(board){
		return errors.New("En el tablero solo pueden estar los siguientes caracteres : b, w, 0, por favor verifique")
	} else if !isValidPlayer(player){
		return errors.New("El jugador debe ser o w o b")
	} else if !hasValidStructure(board){
		return errors.New("Las fichas solo pueden estar en cuadro negros, por favor verifique")
	} else if isDone(board) {
		return errors.New("No hay movimientos por hacer, ya hay un ganador")
	} else{
		return nil
	}

}
