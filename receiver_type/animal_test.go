package main

import (
	"testing"
)

func TestAnimalFood(t *testing.T) {
	cow, _ := GetAnimal("cow")
	bird, _ := GetAnimal("bird")
	snake, _ := GetAnimal("snake")

	t.Run("Cows should eat grass", func(t *testing.T) {
		got := cow.Eat()
		want := "Grass"
		assertAreEqual(got, want, t)
	})
	t.Run("Birds should eat worms", func(t *testing.T) {
		got := bird.Eat()
		want := "Worms"
		assertAreEqual(got, want, t)
	})
	t.Run("Snakes should eat mice", func(t *testing.T) {
		got := snake.Eat()
		want := "Mice"
		assertAreEqual(got, want, t)
	})
}

func TestAnimalLocomotion(t *testing.T) {
	var cow, _ = GetAnimal("cow")
	bird, _ := GetAnimal("bird")
	snake, _ := GetAnimal("snake")

	t.Run("Cows should walk", func(t *testing.T) {
		got := cow.Move()
		want := "Walk"
		assertAreEqual(got, want, t)
	})

	t.Run("Birds should fly", func(t *testing.T) {
		got := bird.Move()
		want := "Fly"
		assertAreEqual(got, want, t)
	})

	t.Run("Snakes should slither", func(t *testing.T) {
		got := snake.Move()
		want := "Slither"
		assertAreEqual(got, want, t)
	})
}

func TestAnimalSound(t *testing.T) {
	var cow, _ = GetAnimal("cow")
	bird, _ := GetAnimal("bird")
	snake, _ := GetAnimal("snake")

	t.Run("Cows should moo", func(t *testing.T) {
		got := cow.Speak()
		want := "Moo"
		assertAreEqual(got, want, t)
	})

	t.Run("Birds should peep", func(t *testing.T) {
		got := bird.Speak()
		want := "Peep"
		assertAreEqual(got, want, t)
	})
	t.Run("Snakes should do hsss", func(t *testing.T) {
		got := snake.Speak()
		want := "Hsss"
		assertAreEqual(got, want, t)
	})
}

func assertAreEqual(got string, want string, tb testing.TB) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %s want %s", got, want)
	}
}
