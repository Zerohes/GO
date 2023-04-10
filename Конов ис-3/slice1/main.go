package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generateRandomSlice(length, limit int) []int {
    rand.Seed(time.Now().UnixNano())
    slice := make([]int, length)
    for i := range slice {
        slice[i] = rand.Intn(limit)
    }
    return slice
}

func bubbleSort(slice []int) {
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}

func main() {
    slice := generateRandomSlice(10, 100)
    fmt.Printf("До сортировки: %v\n", slice)
    bubbleSort(slice)
    fmt.Printf("Отсортированный: %v\n", slice)
}