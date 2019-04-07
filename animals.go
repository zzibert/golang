package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	Name string
}

type Bird struct {
	Name string
}

type Snake struct {
	Name string
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func findAnimal(name string, animals map[string]Animal, typ string) {
	for key, animal := range animals {
		if key == name {
			switch typ {
			case "move":
				animal.Move()
			case "eat":
				animal.Eat()
			case "speak":
				animal.Speak()
			}
		}
	}
}

func createAnimal(name string, animals map[string]Animal, typ string) {
	switch typ {
	case "cow":
		cow := new(Cow)
		cow.Name = name
		animals[name] = cow
	case "bird":
		bird := new(Bird)
		bird.Name = name
		animals[name] = bird
	case "snake":
		snake := new(Snake)
		snake.Name = name
		animals[name] = snake
	}
	fmt.Println("Created it!")
}

func main() {

	animals := make(map[string]Animal)
	var command, name, typ string

	for {
		fmt.Printf(">")
		fmt.Scan(&command, &name, &typ)
		switch command {
		case "newanimal":
			createAnimal(name, animals, typ)
		case "query":
			findAnimal(name, animals, typ)
		}

	}

}