package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (_ Cow) Eat() {
	fmt.Println("grass")
}

func (_ Cow) Move() {
	fmt.Println("walk")
}

func (_ Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (_ Bird) Eat() {
	fmt.Println("worms")
}

func (_ Bird) Move() {
	fmt.Println("fly")
}

func (_ Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (_ Snake) Eat() {
	fmt.Println("mice")
}

func (_ Snake) Move() {
	fmt.Println("slither")
}

func (_ Snake) Speak() {
	fmt.Println("hsss")
}

var animalStorage map[string]Animal

func CreateNewAnimal(animalName string, animalType string) {
	if animalStorage == nil {
		animalStorage = make(map[string]Animal)
	}

	switch animalType {
	case "cow":
		animalStorage[animalName] = Cow{}
	case "bird":
		animalStorage[animalName] = Bird{}
	case "snake":
		animalStorage[animalName] = Snake{}
	default:
		fmt.Println("Third command (animalType) for `createanimal` should be either `cow`, `bird` or `snake``")
		return
	}

	fmt.Println("Created it!")
}

func GetAnimal(animalName string, animalInfo string) {
	animal, exist := animalStorage[animalName]
	if !exist {
		fmt.Println("There is no animal has the name " + animalName)
		return
	}

	switch animalInfo {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Third command (animalInfo) for `query` should be either `eat`, `move` or `speak``")
	}
}

func main() {

	for {
		fmt.Println("Enter `newanimal {animalName} {animalType}` or `query {animalName} {animalInfo}`")
		fmt.Print(">")
		var c1, c2, c3 string
		fmt.Scan(&c1, &c2, &c3)

		switch c1 {
		case "newanimal":
			CreateNewAnimal(c2, c3)
		case "query":
			GetAnimal(c2, c3)
		default:
			fmt.Printf("First command should be either `newanimal` or `query`, not %s\n", c1)
		}
	}
}
