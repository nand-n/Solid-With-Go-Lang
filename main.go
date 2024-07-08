// package main

// import "fmt"

// type Shape interface {
// 	Area() float64
// }

// type Triangle struct {
// 	Base, Height float64
// }

// func (t Triangle) Area() float64 {
// 	return 0.5 * t.Base * t.Height
// }

// type Square struct {
// 	Side float64
// }

// func (s Square) Area() float64 {
// 	return s.Side * s.Side
// }

// type Rectangle struct {
// 	Length, Breadth float64
// }

// func printArea(s Shape) {
// 	fmt.Println("Area of shape is : ", s.Area())
// }

// type blogPost struct {
// 	author string
// 	title  string
// 	postId int
// }

// func NewBlogPost() *blogPost {
// 	return &blogPost{
// 		author: "Nahom",
// 		title:  "The go delima",
// 		postId: 22,
// 	}
// }
// func main() {
// 	t := Triangle{Base: 0.002, Height: 22}
// 	printArea(t)

// 	var b blogPost
// 	newBlogPost := *NewBlogPost()
// 	fmt.Println(newBlogPost, "new Blog ")

//		b = blogPost{
//			author: "Nah",
//			title:  "The go lang delima 2 ",
//			postId: 23,
//		}
//		fmt.Println(b, "b")
//	}

package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Player represents the player's character.
type Player struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	Inventory []string
}

// Room represents a room in the dungeon.
type Room struct {
	Name        string
	Description string
	Enemies     []Enemy
	Puzzles     []Puzzle
	NextRooms   map[string]*Room
}

// Enemy represents an enemy character in the game.
type Enemy struct {
	Name   string
	Health int
	Attack int
}

// Puzzle represents a puzzle or challenge in a room.
type Puzzle struct {
	Description string
	Solved      bool
}

// Game manages the game state and logic.
type Game struct {
	Player      Player
	CurrentRoom *Room
	Rooms       map[string]*Room
}

// NewGame initializes a new game with default settings.
func NewGame(playerName string) *Game {
	player := Player{
		Name:      playerName,
		Health:    100,
		Attack:    15,
		Defense:   5,
		Inventory: []string{"Sword"},
	}

	// Define rooms
	rooms := make(map[string]*Room)
	forestRoom := &Room{
		Name:        "Forest Room",
		Description: "You find yourself in a dark forest. A path leads deeper into the dungeon.",
		Enemies:     []Enemy{{Name: "Goblin", Health: 20, Attack: 8}},
		Puzzles: []Puzzle{
			{Description: "A mysterious altar stands in the center of the room."},
		},
		NextRooms: map[string]*Room{
			"east": nil, // Placeholder for connecting rooms
		},
	}
	caveRoom := &Room{
		Name:        "Cave Room",
		Description: "You enter a dimly lit cave. Shadows flicker on the walls.",
		Enemies:     []Enemy{{Name: "Troll", Health: 30, Attack: 12}},
		Puzzles: []Puzzle{
			{Description: "A boulder blocks the passage ahead."},
		},
		NextRooms: map[string]*Room{
			"west": forestRoom,
		},
	}
	// Connect rooms
	forestRoom.NextRooms["east"] = caveRoom

	rooms["forest"] = forestRoom
	rooms["cave"] = caveRoom

	return &Game{Player: player, CurrentRoom: forestRoom, Rooms: rooms}
}

// HandleMove handles player movement between rooms.
func (game *Game) HandleMove(direction string) *Room {
	if nextRoom, ok := game.CurrentRoom.NextRooms[direction]; ok {
		game.CurrentRoom = nextRoom
		return game.CurrentRoom
	}
	return nil
}

// SolvePuzzle simulates solving a puzzle in a room.
func (game *Game) SolvePuzzle() string {
	// Placeholder logic for solving puzzles
	for _, puzzle := range game.CurrentRoom.Puzzles {
		if !puzzle.Solved {
			// Example: Check if player has required item to solve puzzle
			if containsItem(game.Player.Inventory, "key") {
				puzzle.Solved = true
				return "You unlocked the puzzle!"
			} else {
				return "You need a key to unlock this puzzle."
			}
		}
	}
	return "No puzzles to solve."
}

// containsItem checks if an item exists in the player's inventory.
func containsItem(inventory []string, item string) bool {
	for _, i := range inventory {
		if i == item {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize game
	game := NewGame("Player1")

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define HTTP handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, game)
	})

	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		direction := r.Form.Get("direction")
		nextRoom := game.HandleMove(direction)
		if nextRoom != nil {
			fmt.Fprintf(w, "Moved to %s", nextRoom.Name)
		} else {
			fmt.Fprintf(w, "Cannot move in that direction.")
		}
	})

	http.HandleFunc("/solve-puzzle", func(w http.ResponseWriter, r *http.Request) {
		message := game.SolvePuzzle()
		fmt.Fprintf(w, message)
	})

	// Start server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
