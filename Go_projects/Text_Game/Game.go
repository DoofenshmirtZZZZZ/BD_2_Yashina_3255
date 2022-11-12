package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Parameters struct { // Длина норы = 10		Здоровье = 100 		Уважение = 20		Вес = 30
	LenghtOfBurrow int
	health         int
	respect        int
	weight         int
}

var parameters = Parameters{10, 100, 20, 30}
var numAction, lod, hp, resp, wgh int

func main() {
	for {
		fmt.Println("Чем заняться? \n Копать нору -- 1 \n Проспать весь день -- 2 \n Поесть травки -- 3 \n Подраться -- 4")
		fmt.Print("Выберете действие, введя цифру: ")

		fmt.Scanln(&numAction)
		if numAction <= 0 || numAction > 4 {
			fmt.Println("Неверно, попробуйте снова ")
			continue
		}

		switch numAction {
		case 1:
			lod, hp, resp, wgh = parameters.Dig(numAction)
			parameters = Parameters{lod, hp, resp, wgh}
			fmt.Println("Вы копали нору, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()

			lod, hp, resp, wgh = parameters.Night()
			fmt.Println("День закончился и вы пошли спать, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()
			parameters = Parameters{lod, hp, resp, wgh}
			main()
		case 2:
			lod, hp, resp, wgh = parameters.Night()
			parameters = Parameters{lod, hp, resp, wgh}
			fmt.Println("Вы спали весь день, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()

			lod, hp, resp, wgh = parameters.Night()
			fmt.Println("День закончился и вы пошли спать, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()
			parameters = Parameters{lod, hp, resp, wgh}
			main()
		case 3:
			lod, hp, resp, wgh = parameters.Eat(numAction)
			parameters = Parameters{lod, hp, resp, wgh}
			fmt.Println("Вы поели, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()

			lod, hp, resp, wgh = parameters.Night()
			fmt.Println("День закончился и вы пошли спать, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()
			parameters = Parameters{lod, hp, resp, wgh}
			main()
		case 4:
			lod, hp, resp, wgh = parameters.Fight(numAction)
			parameters = Parameters{lod, hp, resp, wgh}
			fmt.Println("Вы подрались, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()
			parameters.Win()

			lod, hp, resp, wgh = parameters.Night()
			fmt.Println("День закончился и вы пошли спать, ваши параметры:", "\n Длина норы:", lod, "\n Здоровье:", hp, "\n Уважение:", resp, "\n Вес:", wgh)
			parameters.Death()
			parameters = Parameters{lod, hp, resp, wgh}
			main()
		}

	}
}

func (*Parameters) Night() (int, int, int, int) { // Функция для Сна, вызываем после каждого действия, возыращает все параметры
	lod, hp, resp, wgh := parameters.LenghtOfBurrow, parameters.health, parameters.respect, parameters.weight

	lod -= 2
	hp += 20
	resp -= 2
	wgh -= 5
	return lod, hp, resp, wgh

}
func (*Parameters) Dig(numAction int) (int, int, int, int) { // Функция Копать, имеет два разветвления на лениво и интенсивно, возвращает длину норы и hp
	lod, hp, resp, wgh := parameters.LenghtOfBurrow, parameters.health, parameters.respect, parameters.weight

	fmt.Println("Копать нору: интенсивно -- 1 или лениво -- 2")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scanln(&numAction)

	switch numAction {
	case 1:
		fmt.Println("Длина норы увеличилась на 5, здоровье уменьшилось на 30")
		lod += 5
		hp -= 30
	case 2:
		fmt.Println("Длина норы увеличилась на 2, здоровье уменьшилось на 10")
		lod += 2
		hp -= 10
	}
	return lod, hp, resp, wgh

}

func (*Parameters) Eat(numAction int) (int, int, int, int) { // Функция Поесть, имеет два разветвления и одно подразветвление, возвращает здоровье и вес
	lod, hp, resp, wgh := parameters.LenghtOfBurrow, parameters.health, parameters.respect, parameters.weight

	fmt.Println("Поесть травки: жухлой -- 1 или зелёной -- 2")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scanln(&numAction)

	switch numAction {
	case 1:
		fmt.Println("Здоровье увеличилось на 10, вес увеличился на 15")
		hp += 10
		wgh += 15
	case 2:
		if resp < 30 {
			fmt.Println("Здоровье уменьшилось на 30")
			hp -= 30
		} else {
			fmt.Println("Здоровье уменьшилось на 30, вес увеличился на 30")
			wgh += 30
			hp -= 30
		}
	}
	return lod, hp, resp, wgh
}

func (*Parameters) Fight(numAction int) (int, int, int, int) { // Функция Драки, при рандоме равном отнош веса к сумме весов
	lod, hp, resp, wgh := parameters.LenghtOfBurrow, parameters.health, parameters.respect, parameters.weight
	var enemyWeight, wieghtSum, dopReps, minsHp int
	var chanceWin, random float64

	fmt.Println("Подраться: со слабым -- 1 или средним -- 2 или сильным -- 3 ")
	fmt.Print("Выберете действие, введя цифру: ")
	fmt.Scanln(&numAction)

	switch numAction {
	case 1:
		enemyWeight = 30
	case 2:
		enemyWeight = 50
	case 3:
		enemyWeight = 70
	}
	wieghtSum = enemyWeight + wgh
	chanceWin = float64(wgh) / float64(wieghtSum)
	random = rand.Float64()
	rand.Seed(time.Now().UnixNano())

	if random <= chanceWin {
		fmt.Println("Вы победили, ваше уважение повысилось на 10")
		dopReps = enemyWeight - wgh
		if dopReps <= 0 {
			resp += 10
		}
		fmt.Println("Уважение:", resp, "\nСила рандома:", random, "\nШанс успеха:", chanceWin)

	} else {
		fmt.Println("Вы проиграли, ваше здоровье уменьшилось на 10")
		minsHp = enemyWeight - wgh
		if minsHp <= 0 {
			hp -= 10
		}
		fmt.Println("Здоровье:", hp, "\nСила рандома:", random, "\nШанс успеха:", chanceWin)
	}
	return lod, hp, resp, wgh
}

func (*Parameters) Death() { // Функция проверка на состояние параметров, при достижении 0 прерывает работу
	lod, hp, resp, wgh := parameters.LenghtOfBurrow, parameters.health, parameters.respect, parameters.weight

	if (lod <= 0) || (hp <= 0) || (resp <= 0) || (wgh <= 0) {
		fmt.Println("Один из ваших параметров упал до 0, игра окончена")
		os.Exit(1)
	}
}

func (*Parameters) Win() int { // Функция проверка на состояние уважения для определния победы
	resp := parameters.respect

	if resp >= 100 {
		fmt.Println("Вы добились уважения равное или превышающее 100. Победа!")
		os.Exit(1)
	}
	return resp
}
