package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ladygogo/BattleGopher/player"
)

type Session struct {
	playerArray []player.Player
	gameId int
}

type GameInput struct {
	BoardDimension int `json:"board_dimension"`
	PlayerNames []string `json:"player_names"`
}

type SessionData struct {
	SessionId int `json:"session_id"`
}

var SessionCount = 0
var sessions []Session



func main() {
  // Initialize the sessions
	initialLength := 1
	arrayCapacity := 5
	sessions = make([]Session, initialLength, arrayCapacity)

	      // Add router
	router := mux.NewRouter()

	      // Add Handler Functions
	router.HandleFunc("/guess", GuessHandler)
	router.HandleFunc("/newgame", NewGameHandler).Methods("POST")
	router.HandleFunc("/turn", TurnHandler)
	// Separating Turns from Guesses because remote users don't know
	// when it's the opponent's turn...so it's for polling
	router.HandleFunc("/", APIDocHandler)

	// Start the webserver
	fmt.Println("Listening")
	http.ListenAndServe(":8080", router)
}

// Add a Handler function with required arguments
//func Example(w http.ResponseWriter, r *http.Request)
func GuessHandler(w http.ResponseWriter, r *http.Request) {
}

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	var gameInput GameInput

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  fmt.Println("errr is: ", err)
	}
	err = json.Unmarshal(body, &gameInput)
	if err != nil {
	  fmt.Println("unmarshal errr is: ", err)
	}

	session := Session{ gameId: SessionCount, playerArray: []player.Player{} }
	SessionCount++
	for _, playerName := range gameInput.PlayerNames {
		p, err := player.InitializePlayer(playerName)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			errorResponse, _ := json.Marshal(string(err.Error()))
			w.Write(errorResponse)
			return
		}
		p.NewBoard(gameInput.BoardDimension)
		session.playerArray = append(session.playerArray, p)
	}
	sessions = append(sessions, session)
	sessionData := SessionData{SessionId: session.gameId}
	response, err := json.Marshal(sessionData)
	if err != nil {
		fmt.Println("error marshaling", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func TurnHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Battle Gopher Turn Handler!")
	fmt.Println("---------------------")


}

func APIDocHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Battle Gopher New Game Handler!")
	fmt.Println("---------------------")

}
