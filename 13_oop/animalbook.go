package main

import (
	"fmt"
	"strings"
)

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c *Cow) Eat() {
	fmt.Println(c.food)
}

func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c *Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() {
	fmt.Println(b.food)
}

func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b *Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() {
	fmt.Println(s.food)
}

func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s *Snake) Speak() {
	fmt.Println(s.noise)
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func FillAnn(animType string, name string) Animal {
	switch animType {
	case "cow":
		ah := Cow{}
		ah.name = name
		ah.food = "grass"
		ah.locomotion = "walk"
		ah.noise = "moo"
		return &ah
	case "bird":
		ah := Bird{}
		ah.name = name
		ah.food = "worms"
		ah.locomotion = "fly"
		ah.noise = "peep"
		return &ah
	case "snake":
		ah := Snake{}
		ah.name = name
		ah.food = "mice"
		ah.locomotion = "slither"
		ah.noise = "hsss"
		return &ah
	default:
		fmt.Println("There is no such type")
		return nil
	}

}

func UserStr(text string) (error, string, string, string) {
	var input string
	fmt.Printf("%s\n> ", text)

	_, err := fmt.Scan(&input)
	if err != nil {
		return err, "", "", ""
	}
	command := strings.ToLower(input)

	_, err = fmt.Scan(&input)
	if err != nil {
		return err, "", "", ""
	}
	animal := strings.ToLower(input)

	_, err = fmt.Scanln(&input)
	if err != nil {
		return err, "", "", ""
	}
	action := strings.ToLower(input)

	return nil, command, animal, action
}
func main() {
	//var animalBook map[string]Animal
	animalBook := make(map[string]Animal)

	for {
		err, commandName, animalName, actionOrAnimal := UserStr("Create a new animal: newanimal \"animal name\" \"animal type: cow, bird, snake\"\nOr request information: query \"animal name\" \"action type: eat, move, speak\"")
		if err != nil {
			fmt.Println(err, ", try again:")
			continue
		}

		switch commandName {
		case "newanimal":
			if _, ok := animalBook[animalName]; ok {
				fmt.Println("Animal name already exist!")
				continue
			}
			if actionOrAnimal == "cow" || actionOrAnimal == "bird" || actionOrAnimal == "snake" {
				newAnn := FillAnn(actionOrAnimal, animalName)
				animalBook[animalName] = newAnn
				fmt.Printf("Created it! %s %s \n\n", actionOrAnimal, animalName)
				continue
			}
			fmt.Println("No such animal type!")
		case "query":
			_, ok := animalBook[animalName]
			if (actionOrAnimal == "eat" || actionOrAnimal == "move" || actionOrAnimal == "speak") && ok {
				actionlMap := map[string]func(animName string){
					"eat":   func(animName string) { animalBook[animalName].Eat() },
					"move":  func(animName string) { animalBook[animalName].Move() },
					"speak": func(animName string) { animalBook[animalName].Speak() },
				}
				findAnimalAction := actionlMap[actionOrAnimal]
				findAnimalAction(animalName)
				continue
			}
			fmt.Println("No such animal name or action!")
		default:
			fmt.Println("No such command!")
		}
	}

}
