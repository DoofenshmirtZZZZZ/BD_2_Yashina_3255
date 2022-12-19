package functions

import (
	"fmt"
	"math/rand"
)

type creature struct {
	Hole    int
	Health  int
	Respect int
	Weight  int
}

func DefaultStats() creature {
	return creature{
		Hole:    8,
		Health:  100,
		Respect: 10,
		Weight:  20,
	}
}

func (c *creature) CheckWin() bool {
	if c.Respect >= 100 {
		return true
	} else {
		return false
	}
}

func (c *creature) CheckDefeat() bool {
	if c.Health <= 0 || c.Hole <= 0 || c.Respect <= 0 || c.Weight <= 0 {
		return true
	} else {
		return false
	}
}

func (c *creature) Stats() {
	fmt.Printf("Характеристики: Нора %v Здоровье  %v Уважение  %v Вес  %v\n",
		c.Hole, c.Health, c.Respect, c.Weight)
}

func (c *creature) Day() {
	var input int
	fmt.Print("\tНаступает день\n")
	fmt.Print("Чем заняться?:\n" +
		"1. Копать нору\n" +
		"2. Поесть травки\n" +
		"3. Подраться\n" +
		"4. Поспать\n" +
		"Ваш выбор: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		c.dig()
	case 2:
		c.eat()
	case 3:
		c.fight()
	case 4:
		c.sleep()
	}
}

func (c *creature) Night() {
	fmt.Print("Вы легли спать\n")
	c.Hole -= 2
	c.Health += 20
	c.Respect -= 2
	c.Weight -= 5
	c.Stats()
}

func (c *creature) dig() {
	var input int
	fmt.Print("Как будете копать?:\n" +
		"1. Интенсивно\n" +
		"2. Лениво\n" +
		"Ваш выбор: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		c.Hole += 5
		c.Health -= 30
	case 2:
		c.Hole += 2
		c.Health -= 10
	}
	c.Stats()
}

func (c *creature) fight() {
	var input int
	var winrate, enemyWeight, exodus float32
	fmt.Print("С кем хотите подраться?:\n" +
		"1. Слабый (Вес 30)\n" +
		"2. Средний (Вес 50)\n" +
		"3. Сильный (Вес 70)\n" +
		"Ваш выбор: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		enemyWeight = 30
	case 2:
		enemyWeight = 50
	case 3:
		enemyWeight = 70
	}
	winrate = float32(c.Weight) / (float32(c.Weight) + enemyWeight - 25)
	fmt.Printf("Шанс на победу %.2f процента\n", winrate*100)
	if winrate*100 < 50 {
		fmt.Print("Уверены, что хотите вступить в бой?\n" +
			"1. Да\n" +
			"2. Нет\n" +
			"Ваш выбор: ")
		fmt.Scan(&input)
		switch input {
		case 1:
			break
		case 2:
			c.Day()
			return
		}
	}
	fmt.Println("\t Бой начинается!")
	exodus = rand.Float32()
	difference := int(enemyWeight) - c.Weight
	if winrate > exodus {
		fmt.Println("Вы победили!")
		if c.Weight < int(enemyWeight) {
			c.Respect += difference + 10
		}
		if c.Weight == int(enemyWeight) {
			c.Respect += 15
		}
		if c.Weight > int(enemyWeight) {
			c.Respect += 10
		}
	} else {
		fmt.Println("Вы проиграли")
		if c.Weight < int(enemyWeight) {
			c.Health -= difference
		} else {
			fmt.Print("Урон по вам не был нанесен")
		}
	}
	c.Stats()
}

func (c *creature) eat() {
	var input int
	fmt.Print("Какой травки поесть?:\n" +
		"1. Жухлой\n" +
		"2. зеленой\n" +
		"Ваш выбор: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		c.Weight += 15
		c.Health += 10
	case 2:
		if c.Respect >= 30 {
			c.Health += 30
			c.Weight += 30
		} else {
			c.Health -= 30
		}
	}
	c.Stats()
}

func (c *creature) sleep() {
	fmt.Print("Вы решили лечь спать")
	c.Night()
}
