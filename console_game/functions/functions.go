package functions

import (
	"fmt"
	"math/rand"
)

type parameters struct { // Длина норы = 10		Здоровье = 100 		Уважение = 20		Вес = 30
	LenghtOfBurrow int
	Health         int
	Respect        int
	Weight         int
}

func DefaultStats() parameters {
	return parameters{
		LenghtOfBurrow: 10,
		Health:         100,
		Respect:        20,
		Weight:         30,
	}
}

func (pr *parameters) CheckWin() bool { // Функция для проверки победы
	if pr.Respect >= 100 {
		return true
	} else {
		return false
	}
}

func (pr *parameters) CheckDefeat() bool { // Функция для проверки Смерти
	if pr.Health <= 0 || pr.LenghtOfBurrow <= 0 || pr.Respect <= 0 || pr.Weight <= 0 {
		return true
	} else {
		return false
	}
}

func (pr *parameters) Stats() { // Функция для вывода Параметров
	fmt.Printf("Характеристики: Нора %v Здоровье  %v Уважение  %v Вес  %v\n",
		pr.LenghtOfBurrow, pr.Health, pr.Respect, pr.Weight)
}

func (pr *parameters) Day() {
	var input int
	fmt.Println("Чем заняться? \n Копать нору -- 1 \n Проспать весь день -- 2 \n Поесть травки -- 3 \n Подраться -- 4")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		pr.Dig()
	case 2:
		pr.Eat()
	case 3:
		pr.Fight()
	case 4:
		pr.Sleep()
	}
}

func (pr *parameters) Night() { // Функция для Сна, вызываем после каждого действия, меняет все параметры
	fmt.Print("Вы легли спать\n")
	pr.LenghtOfBurrow -= 2
	pr.Health += 20
	pr.Respect -= 2
	pr.Weight -= 5
	pr.Stats()
}

func (pr *parameters) Dig() { // Функция Копать, имеет два разветвления на лениво и интенсивно, меняет длину норы и hp
	var input int
	fmt.Println("Копать нору: интенсивно -- 1 или лениво -- 2")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		fmt.Println("Длина норы увеличилась на 5, здоровье уменьшилось на 30")
		pr.LenghtOfBurrow += 5
		pr.Health -= 30
	case 2:
		fmt.Println("Длина норы увеличилась на 2, здоровье уменьшилось на 10")
		pr.LenghtOfBurrow += 2
		pr.Health -= 10
	}
	pr.Stats()
}

func (pr *parameters) Eat() { // Функция Поесть, имеет два разветвления и одно подразветвление, меняет здоровье и вес
	var input int

	fmt.Println("Поесть травки: жухлой -- 1 или зелёной -- 2")
	fmt.Print("Выберете действие, введя цифру: ")

	fmt.Scan(&input)
	switch input {
	case 1:
		fmt.Println("Здоровье увеличилось на 10, вес увеличился на 15")
		pr.Weight += 15
		pr.Health += 10
	case 2:
		if pr.Respect >= 30 {
			fmt.Println("Здоровье уменьшилось на 30, вес увеличился на 30")
			pr.Health += 30
			pr.Weight += 30
		} else {
			fmt.Println("Здоровье уменьшилось на 30")
			pr.Health -= 30
		}
	}
	pr.Stats()
}

func (pr *parameters) Fight() { // Функция Драки, при рандоме равном отнош веса к сумме весов
	var input int
	var winrate, enemyWeight, exodus float32
	fmt.Println("Подраться: со слабым -- 1 или средним -- 2 или сильным -- 3 ")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scan(&input)
	switch input {
	case 1:
		enemyWeight = 30
	case 2:
		enemyWeight = 50
	case 3:
		enemyWeight = 70
	}
	winrate = float32(pr.Weight) / (float32(pr.Weight) + enemyWeight - 25)
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
			pr.Day()
			return
		}
	}
	fmt.Println("\t Бой!")

	exodus = rand.Float32()
	difference := int(enemyWeight) - pr.Weight
	if winrate > exodus {
		fmt.Println("Вы победили!")
		if pr.Weight < int(enemyWeight) {
			pr.Respect += difference + 10
		}
		if pr.Weight == int(enemyWeight) {
			pr.Respect += 15
		}
		if pr.Weight > int(enemyWeight) {
			pr.Respect += 10
			fmt.Println("Ваше уважение повысилось на 10")

		}
	} else {
		fmt.Println("Вы проиграли")
		if pr.Weight < int(enemyWeight) {
			pr.Health -= difference
		} else {
			fmt.Print("Урон по вам не был нанесен")
		}
	}
	pr.Stats()
}

func (pr *parameters) Sleep() {
	fmt.Print("Вы решили лечь спать")
	pr.Night()
}
