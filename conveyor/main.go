package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Инициализация каналов для передачи чисел и их обработанных значений.
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	// Запуск горутины для чтения ввода от пользователя.
	wg.Add(1)
	go readInput(input, &wg)

	// Запуск горутины для вычисления квадрата числа.
	wg.Add(1)
	go calculateSquare(input, output, &wg)

	// Запуск горутины для вычисления произведения квадрата на 2.
	wg.Add(1)
	go calculateProduct(output, &wg)

	wg.Wait()
}

// readInput считывает числа, введенные пользователем, до получения команды "стоп".
// Числа передаются в канал input для дальнейшей обработки.
func readInput(input chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(input)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите число:")
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			break
		}
		inputStr = strings.TrimSpace(inputStr)
		if strings.ToLower(inputStr) == "стоп" {
			return
		}
		num, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Введите число или 'стоп'")
			continue
		}
		input <- num
	}
}

// calculateSquare получает числа из канала input, вычисляет их квадраты,
// и отправляет результат в канал output.
func calculateSquare(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		square := num * num
		output <- square
		fmt.Println("Квадрат:", square)
	}
	close(output)
}

// calculateProduct получает квадраты чисел из канала output,
// вычисляет произведение каждого квадрата на 2 и выводит результат.
func calculateProduct(output <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for square := range output {
		product := square * 2
		fmt.Println("Произведение:", product)
	}
}
