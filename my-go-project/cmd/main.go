package main

import (
    "fmt"
    "my-go-project/pkg/utils"
)

func main() {
    sum := utils.Add(5, 3)
    difference := utils.Subtract(5, 3)

    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", difference)
}