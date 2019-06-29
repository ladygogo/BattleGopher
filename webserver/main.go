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
  PlayerName string `json:"player_name"`
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

	fmt.Println("Battle Gopher New Game Handler!")
	fmt.Println("---------------------")

  // Receive the
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Println("errr is: ", err)
  }
  err = json.Unmarshal(body, &gameInput)
  if err != nil {
    fmt.Println("unmarshal errr is: ", err)
  }
  fmt.Println(gameInput)


  aPlayer, err := player.InitializePlayer(gameInput.PlayerName)
  if err != nil {
      fmt.Println("initializeplayer errr is: ", err)
  }
  aPlayer.NewBoard(gameInput.BoardDimension)
  fmt.Println(aPlayer)
}

func TurnHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Battle Gopher Turn Handler!")
	fmt.Println("---------------------")


}

func APIDocHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Battle Gopher New Game Handler!")
	fmt.Println("---------------------")

}
