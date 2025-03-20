package main

import "sync"

// Задача: найти проблемы, исправить и пояснить
// пояснение под программой

type Cache struct{
	mutex sync.Mutex
	data map[string]string
}

func NewCache() *Cache{
	return &Cache{
		data: make(map[string]string),
	}
}

func (c * Cache) Set(key, value string){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

func (c *Cache) Get(key string) string{
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Size() > 0{
		return c.data[key]
	}

	return ""
}

func (c *Cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.data)
}







/*
Ошибка 1: самая очевидная и находится она в методе size. При вызове функции size из функции Get
произойдет дедлок, потому что горутина, попадающая в метод get уже захватывает мьютекс
Решение оишбки 1: проверять на пустоту в методе get с помощью второго параметра, который возвращает мапа или
убрать мьютекст в методе size

*/
