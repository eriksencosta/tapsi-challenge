package validations

import (
	"testing"
	"github.com/eriksencosta/tapsi-challenge/constants"
)

func TestIsEmpty(t *testing.T) {
	emptyBoard := "00000000" +
	              "00000000" +
				  "00000000" +
				  "00000000" +
				  "00000000" +
				  "00000000" +
				  "00000000" +
				  "00000000"

	if !isEmpty(emptyBoard){
		t.Error("expected true")
	}
}

func TestIsNotEmpty(t *testing.T) {
	notEmptyBoard := "0b000000" +
					 "00000000" +
					 "00000000" +
					 "00000000" +
					 "00000000" +
					 "00000000" +
					 "00000000" +
					 "00000000"

	if isEmpty(notEmptyBoard){
		t.Error("expected false")
	}
}

func TestIsDone(t *testing.T) {
	boardDone := 	 "0b000000" +
					 "00b00000" +
					 "00000000" +
					 "00000000" +
		             "00000000" +
		             "00000000" +
		             "00000000" +
		             "00000000"

	if !isDone(boardDone){
		t.Error("expected true")
	}
}

func TestIsNotDone(t *testing.T) {
	boardNotDone := "0b000000" +
					"00b00000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"000000w0"

	if isDone(boardNotDone){
		t.Error("expected false")
	}
}

func TestHasCorrectLength(t *testing.T) {
	correctLength :="0b000000" +
					"00b00000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"000000w0"

	if !hasCorrectLength(correctLength){
		t.Error("expected true")
	}
}

func TestHasIncorrectLength(t *testing.T) {
	incorrectLength :=	"0b000000" +
						"00b00000" +
						"00000000" +
						"00000000" +
						"00000000" +
						"00000000" +
						"00000000"

	if hasCorrectLength(incorrectLength){
		t.Error("expected false")
	}
}

func TestHasValidPieces(t *testing.T) {
	validPieces := 	"0b0b0b0b" +
					"00b0b0b0" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"0w0w0w0w" +
					"w0w0w0w0"

	if !hasValidPieces(validPieces){
		t.Error("expected true")
	}
}

func TestHasInvalidPieces(t *testing.T) {
	invalidPieces := 	"0b0b0b0b" +
						"00b0t0b0" +
						"00000000" +
						"00000000" +
						"00000000" +
						"00000000" +
						"0w0w0w0w" +
						"w0w0w0w0"

	if hasValidPieces(invalidPieces){
		t.Error("expected false")
	}
}

func TestHasValidStructure(t *testing.T) {
	validStructure := 	"0b0b0b0b" +
						"00b0b0b0" +
						"00000000" +
						"00000000" +
						"00000000" +
						"00000000" +
						"0w0w0w0w" +
						"w0w0w0w0"

	if !hasValidStructure(validStructure){
		t.Error("expected true")
	}
}

func TestHasInvalidStructure(t *testing.T) {
	invalidStructure := 	"0b0b0b0b" +
							"00b0b0b0" +
							"00000000" +
							"00000000" +
							"00000000" +
							"00000000" +
							"0w0w0w0w" +
							"0w0w0w0w"

	if hasValidStructure(invalidStructure){
		t.Error("expected false")
	}
}

func TestIsValidPlayer(t *testing.T) {
	if !(isValidPlayer(constants.BlackPlayer) && isValidPlayer(constants.WhitePlayer)){
		t.Error("expected true")
	}
}

func TestIsNotValidPlayer(t *testing.T) {
	if isValidPlayer(constants.EmptySquare){
		t.Error("expected false")
	}
}

func TestHasMorePiecesThanAllowed(t *testing.T) {
	board := "0b0b0b0b" +
			 "b0b0b0b0" +
			 "0b0b0b0b" +
			 "b0b0b0b0" +
			 "00000000" +
			 "w0w0w0w0" +
			 "0w0w0w0w" +
			 "w0w0w0w0"

	if !hasMorePiecesThanAllowed(board, "b") {
		t.Error("expected true")
	}
}

func TestHasNotMorePiecesThanAllowed(t *testing.T) {
	board := "0b0b0b0b" +
			 "b0b0b0b0" +
			 "0b0b0b0b" +
			 "b0b0b0b0" +
			 "00000000" +
			 "w0w0w0w0" +
			 "0w0w0w0w" +
			 "w0w0w0w0"

	if hasMorePiecesThanAllowed(board, "w") {
		t.Error("expected false")
	}
}
