package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var rounds int
	var name string

	fmt.Println("Добро пожаловать в <GameName>\nНапечатайте имя своему герою!")
	fmt.Scan(&name)

	hero := &Character{
		Name:    name,
		Health:  100,
		Might:   5,
		Defense: 5,
	}

	enemy := &Character{
		Name:    "Zlodei",
		Health:  100,
		Might:   5,
		Defense: 5,
	}

	fmt.Println("Бой!")

	for {
		rounds++

		fmt.Printf("Раунд %d\n", rounds)

		attack(hero, enemy)
		attack(enemy, hero)

		if enemy.Health <= 0 {
			fmt.Println(hero.Name, "Победил!")
			break
		} else if hero.Health <= 0 {
			fmt.Println(enemy.Name, "Победил!")
			break
		}
	}
}

type Character struct {
	Name                   string
	Health, Might, Defense int
}

func attack(attacker, defender *Character) int {
	damage := attacker.Might + roll(7)
	diceRoll := roll(21)

	fmt.Println(attacker.Name, "Бросил", diceRoll)

	if diceRoll == 1 || diceRoll < defender.Defense {
		defender.Health -= 0

		fmt.Println(attacker.Name, "Промахнулся!")
		fmt.Println(defender.Name, defender.Health, "HP\n")
	} else if diceRoll == 20 {
		defender.Health -= damage * 2
		fmt.Println(attacker.Name, "Наносит критический удар!")
		fmt.Println(defender.Name, defender.Health, "HP\n")
	} else {
		defender.Health -= damage
		fmt.Println(attacker.Name, "Наносит обычный удар!")
		fmt.Println(defender.Name, defender.Health, "HP\n")
	}

	return damage
}

func roll(max int) int {
	min := 1
	rndm := rand.Intn(max-min) + min
	return rndm
}
