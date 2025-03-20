package main

import "sync"

// Задание: найти проблемы в коде, исправить и объяснить
// Поянснение находится под программой.

type Stack struct{
	mutex sync.Mutex
	data []string
}

func NewStack() Stack{
	return Stack{}
}

func (b Stack) Push(value string){
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}

func (b Stack) Pop(){
	if len(b.data) == 0{
		panic("pop stack is empty")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = b.data[:len(b.data)- 1]
}

func (b Stack) Top() string{
	if len(b.data) == 0{ 
		panic("top: stack is empty")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.data[len(b.data) - 1]
}

var stack Stack

func producer(){
	for i := 0; i < 1000; i++{
		stack.Push("message")
	}
}

func consumer(){
	for i := 0; i < 10; i++{
		_ = stack.Top()
		stack.Pop()
	}
}

func main(){
	producer()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++{
		go func(){
			defer wg.Done()
			consumer()
		}()
	}

	wg.Wait()
}

/*
Ошибка 1: структура Stack передается по значению а не по указателю,
что приводит к созданию копий этой структуры, которые никак не связаны между собой

Ошибка 2: методы pop и top неправильно защищены мьютексом, так как он блокируется после проверки
на пустоту слайса, а значит у двух/нескольких потоков могут оказаться неактуальные данные
относительно друг друга. Простыми словами race condtition.
Пример ошибки 2: допустим в стеке остался один элемент, поток А вызывает методо top, который следит
за верхним элементом в стеке, он проверяет не пуст ли слайс, как написано выше, в нем один элемент,
соответственно условие не выполняется и код продолжается. В этот момент поток В вызывает метод pop который
удаляет элемент из стека и получается ситуация в которой поток A и потоко B имеют разные данные
об одном слайсе, что может привести к гонке данных.
Решение ошибки 2: поместить в мьютекст проверку на пустоту.

Ошибка 3: метод pop срезает последний элемент слайса, но он по прежнему остается в базовом массиве, хотя 
и не доступен в функциях, следовательно возникает утечка памяти.
Решение ошибки 3: обнуление срезанного элемента.

Ошибка 4: глобальная переменная stack которая вызывается различными функциями
может привести к ошибкам, и если ошибка возникнет при тестировании, будет непонятно
в какой функции произошла ошибка.
Решение ошибки 4: передавать переменную stack в виде параметра.

*/
