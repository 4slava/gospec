package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func UserStr(text string) (error, string, string) {
	var input string
	fmt.Printf("%s\n> ", text)
	_, err := fmt.Scan(&input)
	if err != nil {
		return err, "", ""
	}
	nameAnimal := strings.ToLower(input)
	_, err = fmt.Scanln(&input)
	if err != nil {
		return err, "", ""
	}
	nameAction := strings.ToLower(input)
	return nil, nameAnimal, nameAction
}

func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	animalMap := map[string]*Animal{
		"cow":   &cow,
		"bird":  &bird,
		"snake": &snake,
	}
	actionlMap := map[string]func(animName string){
		"eat":   func(animName string) { animalMap[animName].Eat() },
		"move":  func(animName string) { animalMap[animName].Move() },
		"speak": func(animName string) { animalMap[animName].Speak() },
	}

	for {
		err, animalName, actionName := UserStr("Enter name of one animal: cow, bird or snake and space-separated one action: eat, move or speak:")
		if err != nil {
			fmt.Println(err, ", try again:")
			continue
		}
		if _, foundAnimal := animalMap[animalName]; foundAnimal != true {
			fmt.Println("There is no such animal, try again:")
			continue
		}
		if _, foundAction := actionlMap[actionName]; foundAction != true {
			fmt.Println("There is no such action, try again:")
			continue
		}
		f := actionlMap[actionName]
		f(animalName)
	}

}
