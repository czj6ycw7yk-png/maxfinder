package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000
	CHUNKS = 8

	maxRandomValue = 1_000_000
)

// generateRandomElements создаёт слайс случайных положительных чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(maxRandomValue) + 1
	}

	return elements
}

// maximum находит максимальное значение в слайсе.
func maximum(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	maxValue := elements[0]
	for _, value := range elements[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks делит слайс на CHUNKS частей и ищет максимум в каждой части
// в отдельной горутине.
func maxChunks(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	if len(elements) < CHUNKS {
		return maximum(elements)
	}

	chunkSize := len(elements) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(elements)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaximums[index] = maximum(chunk)
		}(i, elements[start:end])
	}

	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	elements := generateRandomElements(SIZE)

	start := time.Now()
	singleMax := maximum(elements)
	singleDuration := time.Since(start).Microseconds()

	start = time.Now()
	parallelMax := maxChunks(elements)
	parallelDuration := time.Since(start).Microseconds()

	fmt.Printf("Размер массива: %d\n", SIZE)
	fmt.Printf("Максимум в одном потоке: %d, время: %d мкс\n", singleMax, singleDuration)
	fmt.Printf("Максимум в %d горутинах: %d, время: %d мкс\n", CHUNKS, parallelMax, parallelDuration)

	if singleMax != parallelMax {
		fmt.Println("Ошибка: результаты не совпадают")
	}
}