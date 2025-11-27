package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Player struct {
    Number int
    Value int
}

type Game struct {
	Players []Player
	UserChoice int
}

/*
Комбинации победы

1 - камень
2 - ножницы
3 - бумага
*/
var winComb = map[int]int{
    3: 1,
    1: 2,
	2: 3,
}

func getRandomValForBot() int {
	return rand.Intn(3) + 1
}

// Можно добавлять комменты к функциям, структурам, интерфейсам и тд и тп
// 
// Они будут видны при просмотре декларации функции
func getPlayValues(game Game) []int {
	allValues := make([]int, 0, len(game.Players)+1) // можно использовать cap, но вроде как эффективнее сделать make([]int, len) и по индексам добавлять новые элементы вместо append
	for _, p := range game.Players {
        allValues = append(allValues, p.Value)
	}
	allValues = append(allValues, game.UserChoice)
	return allValues
}

// Использование указателя на int (и других простых типах) - антипаттерн, потому что:
// - Происходит лишняя аллокация (перемещение значение в кучу и возврат указателя
// на это значение вместо прямого возврата из стека)
// - Лишняя операция по переходу по адресу значения
// 
// Как будто лучше вместо указателя возвращать само значение и вместо nil
// использовать например -1 или 0 (вроде отсчет у тебя начинается с 1)
func gameCore(game Game) (int, error) {
	playValues := getPlayValues(game)
    // Если я правильно понял, условие бессмысленное так как никогда не выполнится
	if len(playValues) < 2 {
        return -1, errors.New("недостаточное количество игроков")
    }

	win_choice := playValues[0]
	one_comb := true
	for i := 1; i < len(playValues); i++ {
		v := playValues[i]
		if v == win_choice {
			continue
		}
		if winComb[v] == win_choice {
			if !one_comb {
				return -1, nil
			}
			win_choice = v
		}
		one_comb = false
	}
	if one_comb {
		return -1, nil
	}
	return win_choice, nil
}

func getUserResult(game Game, winValue int) string {
	if game.UserChoice == winValue {
		return "Вы победили"
	}
	return "Вы проиграли"
}

func main() {
	var numBots int
	var userVal int

	for {
		fmt.Print("Введите количество ботов (1 - 6):\n")
		_, err := fmt.Scanln(&numBots)
		if err != nil {
			// fmt.Println("Ошибка ввода:", err)
			fmt.Printf("Ошибка ввода: %v\n", err) // Можно использовать Printf (print format) для удобной работы со строками и динамическими значениями
			continue
		}
		if numBots < 1 || numBots > 6 {
			fmt.Println("Количество ботов должно быть от 1 до 6")
		} else { 
			break // Не ошибка, но вроде как советую не делать так, а всегда раскрывать выражения для удобности чтения
		}
    }

	for {
		fmt.Print("Введите значение (1 - камень, 2 - ножницы, 3 - бумага):\n")
		fmt.Scanf("%d", &userVal)
		if userVal < 1 || userVal > 3 {
			fmt.Println("Значение должно быть от 1 до 3")
			continue
		} else { // Здесь вроде тоже можно не писать else
			bots := make([]Player, numBots)
			for i := 0; i < numBots; i++ {
				bots[i] = Player{
					Number: i+1,
					Value: getRandomValForBot(),
				}
			}
		
			game := Game{
				Players: bots,
				UserChoice: userVal,
			}
		
			fmt.Println(game)
		
			winValue, err := gameCore(game)
			// В Go часто вместо длинных if-else используют switch. Можешь попробовать заменить на нее
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			} else if winValue == -1 { // Здесь просто проверяем равенство с -1
				fmt.Println("Ничья")
				continue
			} // else {
			// Здесь можно не писать else, так как в любом случае сюда придем, если остальные условия не выполнились
			fmt.Println("Выбор победителя:", winValue) // так как используем напрямую значение, не приходится разыменовывать указатель
			userResult := getUserResult(game, winValue)
			fmt.Println(userResult)
			// }
		}
	}
}
