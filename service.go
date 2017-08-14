/*
The code of the service is based on https://thenewstack.io/make-a-restful-json-api-go/
 */
package main


import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/eriksencosta/tapsi-challenge/movements"
	"github.com/eriksencosta/tapsi-challenge/validations"
)

type InitialGameState struct {
	Board  *string `json:"board"`
	Player *string `json:"player"`
}

type FinalGameState struct {
	Play  string `json:"play"`
	Play2d string `json:"play2d"`
	Error string `json:"error"`
}

const ContentType = "Content-Type"
const JsonType = "application/json; charset=UTF-8"


func getMovement(w http.ResponseWriter, req *http.Request) {
	var initialGameState InitialGameState
	var finalGameState FinalGameState

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := req.Body.Close(); err != nil {
		panic(err)
	}


	if err := json.Unmarshal(body, &initialGameState); err != nil {
		w.Header().Set(ContentType, JsonType)
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else if err := validations.ValidateRequest(initialGameState.Board, initialGameState.Player); err != nil{
		http.Error(w, err.Error(), 400)
	} else if err := validations.Validate(*initialGameState.Board, *initialGameState.Player); err != nil{
		http.Error(w, err.Error(), 400)
	} else {
		move, gameError := movements.MakeMovement(*initialGameState.Board, *initialGameState.Player)

		w.Header().Set(ContentType, JsonType)
		w.WriteHeader(http.StatusOK)

		if gameError != nil {
			finalGameState.Error = gameError.Error()
			if err := json.NewEncoder(w).Encode(finalGameState); err != nil {
				panic(err)
			}
		}else{
			finalGameState.Play = movements.MovementToString(move)
			finalGameState.Play2d = movements.Movement2DToString(move)

			if err := json.NewEncoder(w).Encode(finalGameState); err != nil {
				panic(err)
			}
		}

	}
}
