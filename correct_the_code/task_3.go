package main

import "sync"

// Задача: найти проблемы в коде, исправить и пояснить
// объяснение находится под программой

type Buffer struct{
	mtx sync.Mutex
	data []int
}

func NewBuffer() *Buffer{
	return &Buffer{}
}

func (b* Buffer) Add(value int){
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.data = append(b.data, value)
}

func (b *Buffer) Data() []int{
	b.mtx.Lock()
	defer b.mtx.Unlock()

	return b.data
}






/*
Ошибка: возврат слайса по указателю а не по значению.
Решение: вернуть слайс по значению, то есть копию. Так как если
будет запущено множество горутин, которые вызывают функцию add и функцию data
может произойти ситуация, что после возврата слайса, функция add продолжит его изменять.
*/
