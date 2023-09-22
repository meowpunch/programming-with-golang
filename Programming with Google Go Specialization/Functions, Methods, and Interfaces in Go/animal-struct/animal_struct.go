package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	animals := map[string]Animal{
		"cow":   Animal{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  Animal{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"},
	}

	fmt.Println("Enter name of the animal (either `cow`, `bird` or `snake`) and name of the information requested (either `eat`, `move` or `speak`")

	for {
		fmt.Print(">")

		var name, info string
		_, err := fmt.Scan(&name, &info)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		animal, exist := animals[name]

		if !exist {
			fmt.Println("Name of the animal should be either cow, bird or snake:", err)
			continue
		}

		if info == "eat" {
			animal.Eat()
		} else if info == "move" {
			animal.Move()
		} else if info == "speak" {
			animal.Speak()
		} else {
			fmt.Println("Name of the information should be either eat, move or speak:", err)
		}
	}

}
