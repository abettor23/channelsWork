package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)     // Канал для получения сигналов.
	signal.Notify(sigs, syscall.SIGINT) // Подписываемся на сигнал SIGINT (Ctrl+C)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-sigs:
				// Когда получаем сигнал SIGINT, выводим сообщение и выходим из горутины.
				fmt.Println()
				fmt.Println("Выхожу из программы.")
				wg.Done()
				return
			default:
				// В обычном режиме выводим квадрат текущего числа и увеличиваем счётчик.
				fmt.Printf("Квадрат числа %d равен %d\n", i, i*i)
				i++
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
}
