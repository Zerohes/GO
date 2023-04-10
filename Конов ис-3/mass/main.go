package main

import (
    "fmt"
    "math/rand"
)

func main() {

    arr := make([][]int, 5)
    for i := range arr {
        arr[i] = make([]int, 5)
    }


    set := make(map[int]bool)
    for i := range arr {
        for j := range arr[i] {
            var n int
            for {
                n = rand.Intn(25) + 1 
                if !set[n] {
                    break
                }
            }
            set[n] = true
            arr[i][j] = n
        }
    }

   
    maxSum := 0
    maxRowIndex := 0
    for i := range arr {
        rowSum := 0
        for j := range arr[i] {
            rowSum += arr[i][j]
        }
        if rowSum > maxSum {
            maxSum = rowSum
            maxRowIndex = i
        }
    }


    fmt.Println("Массив:")
    for i := range arr {
        fmt.Println(arr[i])
    }
    fmt.Printf("Строка с самой большей суммой чисел: %v\n", arr[maxRowIndex])
}