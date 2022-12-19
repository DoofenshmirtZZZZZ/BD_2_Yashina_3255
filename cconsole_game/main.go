package main

import (
	"console_game/functions"
	"fmt"
)

func main() {
	creature := functions.DefaultStats()
	fmt.Print("\tИгра началась\n")
	creature.Stats()
	for {
		creature.Day()
		if creature.CheckWin() {
			fmt.Print("Вы победили!")
			break
		}
		if creature.CheckDefeat() {
			fmt.Print("Вы проиграли")
			break
		}

		creature.Night()
		if creature.CheckWin() {
			fmt.Print("Вы победили!")
			break
		}
		if creature.CheckDefeat() {
			fmt.Print("Вы проиграли")
			break
		}
	}

}
