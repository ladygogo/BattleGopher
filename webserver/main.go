package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ladygogo/BattleGopher/player"
)

type Session struct {
	playerArray []player.Player
	gameId      int
	turnId      int
}

type GameInput struct {
	BoardDimension int      `json:"board_dimension"`
	PlayerNames    []string `json:"player_names"`
}

type SessionData struct {
	SessionId int `json:"session_id"`
}

type GuessInput struct {
	SessionId int `json:"session_id"`
	PlayerId  int `json:"player_id"`
	Row       int `json:"row"`
	Col       int `json:"col"`
}

var SessionCount = 0
var sessions []Session

func main() {
	// Initialize the sessions
	sessions = make([]Session, 0, 5)

	// Add router
	router := mux.NewRouter()

	// Add Handler Functions
	router.HandleFunc("/guess", GuessHandler).Methods("POST")
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
	// Create struct to hold Guess inputs and unmarshal into struct
	var guessInput GuessInput

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("errr is: ", err)
	}
	err = json.Unmarshal(body, &guessInput)
	if err != nil {
		fmt.Println("unmarshal errr is: ", err)
	}
	// Need session_id, player_id (who is guessing), row, col

	// Retrieve session from sessions array
	if guessInput.SessionId >= len(sessions) {
		fmt.Println("Invalid session id")
		return
	}
	session := &sessions[guessInput.SessionId]
	// Ensure game is not over AND player_id matches eligible player's id (DetermineEligiblePlayer) and return if correct player did not guess
	playerId := DetermineEligiblePlayer(*session)
	if playerId != guessInput.PlayerId {
		fmt.Println("Not this player's turn!!!")
		return
	}
	// Take row and column, and use *other* player's CheckForHit function; check for error
	gameboard := &session.playerArray[(playerId+1)%2].Gameboard
	hit, err := gameboard.CheckForHit(guessInput.Row, guessInput.Col)
	if err != nil {
		fmt.Println(err)
		return
	}
	// If CheckForHit returns true, check to see if all gophers are sunk on *other* players board
	// Else If False, increment the turn counter (in the session) (if hit, get another turn)
	gameOver := false
	if hit {
		gameOver = gameboard.AllGophersSunk()
	} else {
		session.turnId++
	}
	// Return the API response (hit: true/false, game_over: true/false ?)
	response := map[string]bool{
		"hit": hit, "game_over": gameOver,
	}
	output, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(output)
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

	session := Session{gameId: SessionCount, playerArray: []player.Player{}}
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
	// Unmarshal into Session struct

	// Lookup Session in array of Sessions

	// Ensure game is not over

	// Use DetermineEligiblePlayer to determine which player is eligible to take a turn

	// Build API response (map of string to string and marshal?)
	fmt.Println("Battle Gopher Turn Handler!")
	fmt.Println("---------------------")

}

func APIDocHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Battle Gopher New Game Handler!")
	fmt.Println("---------------------")

}

// Create function that computes the index of the player who is eligible for a turn
func DetermineEligiblePlayer(session Session) int {
	return session.turnId % 2
}
