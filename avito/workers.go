package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func do(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		println("rejected")
		return
	default:
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d done\n", id)
	}
}

func main() {

	const taskCount = 10

	var (
		t  = time.Now()

		wg = &sync.WaitGroup{}
		ctx, cancel = context.WithTimeout(
			context.Background(),
			time.Millisecond * 150,
		)
	)

	defer cancel()

	wg.Add(taskCount)
	for id := 0; id < taskCount; id++ {
		go do(ctx, wg, id)
	}

	wg.Wait()

	fmt.Println(time.Now().Sub(t))
}

//func do(ctx context.Context, id int) {
//	time.Sleep(100 * time.Millisecond)
//	fmt.Printf("%d done\n", id)
//}

//1) В приведенном примере код выполняется последовательно. Чтобы запустить его в асинхронном режиме, нужно воспользоваться ключевым словом go при запуске функции do.
//2) Чтобы в основной горутине (функция main) дождаться выполнения воркеров, нужно воспользоваться примитивами синхронизации. В моем примере кода используется WaitGroup из пакета sync.
//3) Раз уж в функцию передается Context, нужно дописать его обработку в воркере. Пример: если результат работы воркера уже не ожидается.
//4) Хорошим тоном является добавлять максимальное время на выполнение воркера. Реализовать это можно с помощью context.WithTimeout.
//Плюс моего решения: В случае превышения времени выполнения программы или при закрытии окна браузера пользователем воркеры будут завершены.