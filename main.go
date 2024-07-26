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
// MAPS
users:=map[string]int{
	"Humphrey":28,
} 
users["Brian"]=29
// Second way
clubs :=make(map[string]string)
clubs["Bayern"]="Germany"
clubs["Chelsea"]="England"
fmt.Printf("Users: \n%+v\n",users)
fmt.Printf("Clubs: \n%+v\n",clubs)

// Checking if value exists

age, ok:=users["Brian"]
// deleting from the map (MAPTODELETEFROM, KEY)
delete(users, "Humphrey")
if !ok {
	fmt.Println("Doesn't exist")
	}else{
	fmt.Println("Exist: ",age)
}
// lOOPING THROUGH MAPS
for k,v := range clubs {
	fmt.Printf("The key is %s and the value is %s\n",k,v)
}
// Slices
numbers :=[]int{}
otherNumbers:=make([]int,0)
fmt.Println(numbers)
fmt.Println(otherNumbers)

// Arrays
scores :=[3]int{}
scores[0]= 3
scores[1]= 6
scores[2]= 5
numbers = append(numbers, scores[2])
numbers = append(numbers, scores[0])
fmt.Println(scores)
fmt.Println(numbers)
}