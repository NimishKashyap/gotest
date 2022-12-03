//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health int

const (
	maxHealth = 10
	maxEnergy = 100
)

type Energy int

type Player struct {
	health    Health
	maxHealth Health

	energy    Energy
	maxEnergy Energy
}

func (player *Player) changeHealth(num Health) {
	if player.health+num < 0 {
		fmt.Println("Player is dead")
		return
	}
	if player.health+num > player.maxHealth {
		fmt.Println("Maximum health")
		return
	}
	player.health += num
	fmt.Println("Updated health:", player.health)
}
func (player *Player) changeEnergy(num Energy) {
	if player.energy+num < 0 {
		fmt.Println("Player is starving")
		return
	}
	if player.energy+num > player.maxEnergy {
		fmt.Println("Maximum Energy reached")
		return
	}
	player.energy += num
	fmt.Println("Updated energy:", player.energy)
}

func main() {

}
