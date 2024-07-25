package main

import "fmt"

var name = "Kibet"
const firstName="Humphrey"

var (
	occupation = "developer"
	age        int
	language   = "golang"
)

const (
	mentor = "HHH"
	show = "WWE"
)
// Built-in types and custom types

var (
	floatVar32 float32 = 1.4
	floatVar64 float32 = 3.4
	int64Var int64 = 365
	int32Var int32 = 56
	uint64Var uint64 =434
	uint32Var uint32 = 32
	intVar int =33
	byteVar byte = 0x2
	runeVar rune = '5'
)

type Player struct {
	name string
	health int
	attackPower float64
}
// Attaching method to a struct
func (player Player) getHealth() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}


func main() {
	country := "Kenya"

	fmt.Println(country)
	fmt.Println(firstName)
	fmt.Println(mentor)

	player :=Player{
		name:firstName,
		health: 100,
		attackPower: 45.1,
	}
fmt.Printf("This is the player: \n%v\n",player)
// Printing with properties
fmt.Printf("Player: \n%+v\n",player)
fmt.Printf("Player's health: %d\n",getHealth((player)))
// Using the methdd
fmt.Printf("Player's health: %d\n",player.getHealth())

}