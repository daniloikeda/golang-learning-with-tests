package main

import "strings"

type Animal struct {
	food     string
	movement string
	sound    string
}

type AnimalDict map[string]Animal

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.movement
}

func (animal *Animal) Speak() string {
	return animal.sound
}

func GetAnimals() map[string]Animal {
	animals := make(AnimalDict)

	animals["cow"] = Animal{food: "Grass", movement: "Walk", sound: "Moo"}
	animals["bird"] = Animal{food: "Worms", movement: "Fly", sound: "Peep"}
	animals["snake"] = Animal{food: "Mice", movement: "Slither", sound: "Hsss"}

	return animals
}

func GetAnimal(name string) (Animal, bool) {
	animals := GetAnimals()

	animal, found := animals[strings.ToLower(name)]
	return animal, found
}
