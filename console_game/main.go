package main

import (
	"console_game/functions"
	"fmt"
)

func main() {
	parameters := functions.DefaultStats()
	fmt.Print("\tИгра началась\n")
	parameters.Stats()
	for {
		parameters.Day()
		if parameters.CheckWin() {
			fmt.Print("Вы победили!")
			break
		}
		if parameters.CheckDefeat() {
			fmt.Print("Вы проиграли")
			break
		}

		parameters.Night()
		if parameters.CheckWin() {
			fmt.Print("Вы победили!")
			break
		}
		if parameters.CheckDefeat() {
			fmt.Print("Вы проиграли")
			break
		}
	}

}
