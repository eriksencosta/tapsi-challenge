package movements

import (
	"testing"
	"github.com/jfsanin/challenge/constants"
)

var board = "00000000" +
			"00000000" +
			"00000000" +
			"00000000" +
			"000b0000" +
			"00w00000" +
			"00000000" +
			"00000000"

var whiteRightBoard = 	"00000000" +
						"00000000" +
						"00000000" +
						"00000000" +
						"000b0000" +
						"0000w000" +
						"00000000" +
						"00000000"



func TestIsThereAnEnemy(t *testing.T) {
	if !IsThereAnEnemy(representBoardIn2D(board), constants.WhitePlayer, 4 , 3){
		t.Error("Error expect true ")
	}
}

func TestIsNotThereAnEnemy(t *testing.T) {
	if IsThereAnEnemy(representBoardIn2D(board), constants.WhitePlayer, 4 , 1){
		t.Error("Error expect false")
	}
}

func TestMakeEatingMovementWhiteRight(t *testing.T) {

	expectedMove := Movement{
		initialRow: 5,
		initialColumn : 2,
		finalRow: 3,
		finalColumn: 4,
	}

	finalMove, error := MakeMovement(board, constants.WhitePlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeEatingMovementWhiteLeft(t *testing.T) {

	expectedMove := Movement{
		initialRow: 5,
		initialColumn : 4,
		finalRow: 3,
		finalColumn: 2,
	}

	finalMove, error := MakeMovement(whiteRightBoard, constants.WhitePlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeEatingMovementBlackRight(t *testing.T) {

	expectedMove := Movement{
		initialRow: 4,
		initialColumn : 3,
		finalRow: 6,
		finalColumn: 5,
	}

	finalMove, error := MakeMovement(whiteRightBoard, constants.BlackPlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeEatingMovementBlackLeft(t *testing.T) {

	expectedMove := Movement{
		initialRow: 4,
		initialColumn : 3,
		finalRow: 6,
		finalColumn: 1,
	}

	finalMove, error := MakeMovement(board, constants.BlackPlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}


func TestMakeNormalMovementBlackLeft(t *testing.T) {

	normalBoard := 	"00000000" +
					"w0000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"000b0000" +
					"00000000"

	expectedMove := Movement{
		initialRow: 5,
		initialColumn : 2,
		finalRow: 6,
		finalColumn: 1,
	}

	finalMove, error := MakeMovement(normalBoard, constants.BlackPlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeNormalMovementBlackRight(t *testing.T) {

	normalBoard := 	"00000000" +
					"w0000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"0b000000" +
					"00000000"

	expectedMove := Movement{
		initialRow: 5,
		initialColumn : 2,
		finalRow: 6,
		finalColumn: 3,
	}

	finalMove, error := MakeMovement(normalBoard, constants.BlackPlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeNormalMovementWhiteLeft(t *testing.T) {

	normalBoard := 	"00000000" +
					"00000000" +
					"0000000w" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"000b0000" +
					"00000000"

	expectedMove := Movement{
		initialRow: 2,
		initialColumn : 7,
		finalRow: 1,
		finalColumn: 6,
	}

	finalMove, error := MakeMovement(normalBoard, constants.WhitePlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestMakeNormalMovementWhiteRight(t *testing.T) {

	normalBoard := 	"00000000" +
					"w0000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"0b000000" +
					"00000000"

	expectedMove := Movement{
		initialRow: 1,
		initialColumn : 0,
		finalRow: 0,
		finalColumn: 1,
	}

	finalMove, error := MakeMovement(normalBoard, constants.WhitePlayer)

	if error != nil{
		t.Error("Error unexpected " + error.Error())
	}

	if finalMove != expectedMove{
		t.Error("Error the expect value is different to the got value ")
	}
}

func TestNoMoveCornerWhite(t *testing.T) {

	noMoveBoard := 	"0w000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"0b000000" +
					"00000000"


	_, error := MakeMovement(noMoveBoard, constants.WhitePlayer)

	if error == nil{
		t.Error("Error expected ")
	}

}

func TestNoMoveWhite(t *testing.T) {

	noMoveBoard := 	"0w0w0w00" +
					"00w0w000" +
					"000w0000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"0b000000" +
					"00000000"


	_, error := MakeMovement(noMoveBoard, constants.WhitePlayer)

	if error == nil{
		t.Error("Error expected ")
	}
}

func TestNoMoveBlack(t *testing.T) {
	noMoveBoard := 	"0w000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00b00000" +
					"0b0b0000" +
					"b0b0b000"


	_, error := MakeMovement(noMoveBoard, constants.WhitePlayer)

	if error == nil{
		t.Error("Error expected ")
	}
}

func TestNoMoveCornerBlack(t *testing.T) {
	noMoveBoard := 	"0w000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"00000000" +
					"b0000000"


	_, error := MakeMovement(noMoveBoard, constants.WhitePlayer)

	if error == nil{
		t.Error("Error expected ")
	}
}