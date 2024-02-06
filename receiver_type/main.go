package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		animal, method, err := getClientInput()

		if err != nil {
			fmt.Printf("%v\n\n", err)
			continue
		}

		resp, err := processRequest(animal, method)

		if err != nil {
			fmt.Printf("%v\n\n", err)
			continue
		}

		fmt.Printf("%s\n\n", resp)
	}
}

func getClientInput() (animal string, action string, err error) {
	animalName := ""
	animalMethodName := ""
	fmt.Print("Enter with you request (e.g 'Cow Speak'): ")
	_, output := fmt.Scanln(&animalName, &animalMethodName)

	if output != nil {
		return "", "", fmt.Errorf("invalid input")
	}

	return animalName, animalMethodName, nil
}

func processRequest(animalName string, method string) (response string, err error) {
	animal, err := processAnimal(animalName)

	if err != nil {
		return "", err
	}

	return processMethod(method, animal)
}

func processAnimal(name string) (animal Animal, err error) {
	animal, found := GetAnimal(name)

	if !found {
		return animal, fmt.Errorf("invalid animal")
	}

	return animal, nil
}

func processMethod(name string, animal Animal) (response string, err error) {
	switch strings.ToUpper(name) {
	case "EAT":
		return animal.Eat(), nil
	case "MOVE":
		return animal.Move(), nil
	case "SPEAK":
		return animal.Speak(), nil
	default:
		return "", fmt.Errorf("invalid action")
	}
}
