//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	return Player{
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

func TestChangeHealth(t *testing.T) {
	player := newPlayer()
	player.changeHealth(999)

	if player.health > player.maxHealth {
		t.Fatalf("Health went over limit %v, want %v", player.health, player.maxHealth)
	}

	player.changeHealth(player.health - 1)

	if player.health < 0 {
		t.Fatalf("Health: %v. Minimum: 0", player.health)
	}

	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}

}
func TestChangeEnergy(t *testing.T) {
	player := newPlayer()
	player.changeEnergy(999)

	if player.energy > player.maxEnergy {
		t.Fatalf("Health went over limit %v, want %v", player.health, player.maxHealth)
	}

	player.changeEnergy(player.energy - 1)

	if player.energy < 0 {
		t.Fatalf("Health: %v. Minimum: 0", player.health)
	}

	if player.energy > player.maxEnergy {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}

}
