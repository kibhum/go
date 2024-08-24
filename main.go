package main

import (
	"fmt"
	"mygoproject/types"
	"mygoproject/util"
)

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

type WeaponType int

const (
	Axe WeaponType = iota //Increment below
	Sword  
	WoodenStick  
	Knife  
)

func (w WeaponType) String() string {
	switch w {
	case Sword:{
		return "SWORD"
	}
	case Axe:{
		return "AXE"
	}
	default:
		return ""
	}
}
// const (
// 	Axe WeaponType =1
// 	Sword WeaponType =2
// 	WoodenStick WeaponType =3
// 	Knife WeaponType =4
// )

func getDamage(weaponType WeaponType) int{
	switch (weaponType) {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("unknown weapon type")	
	}
}

type NumberStorer interface {
	GetAll()([]int, error)
	Put(int) error
}

type ApiServer struct{
	numberStore NumberStorer
}

type MongoDBNumberStore struct{
	// some value
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
return []int{1,2,3},nil
}

func (m MongoDBNumberStore) Put(number int)  error {
fmt.Println("Store the number to MongoDB")
return nil
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
fmt.Println(getDamage(Axe))

// Control structures
nums := []int{1,2,3,4,5,6,7} 
for i:=0;i<len(nums); i++ {
	fmt.Println(nums[i])
}

names := []string{"a", "b", "c", "d", "e", "f"}
for i,name := range names {
	fmt.Println(i,name)
}

apiServer :=ApiServer{
	numberStore: MongoDBNumberStore{},
}
nms,err:=apiServer.numberStore.GetAll();
if err != nil {
	panic(err)
}
fmt.Println(nms)


user := types.User{
	Username: "James",
	Age: util.GetAge(),
}
 
fmt.Printf("The user is %+v\n: ", user)

e := &types.Entity{
	Name: "my entity",
	Id: "1",
	Version: "1.1",
	Posx: 100,
	Posy: 200,
}
fmt.Printf("\n%+v\n", e)
}