package main

import (
	"fmt"
)

func main() {
    // Start with the number n
    n := 12385

    // Print the initial value of n
    fmt.Print(n)

    // Keep generating the next number in the sequence until we reach 1
    for n != 1 {
        if n % 2 == 0 {
            // If n is even, divide it by 2
            n = n / 2
        } else {
            // If n is odd, multiply it by 3 and add 1
            n = 3*n + 1
        }

        // Print the next number in the sequence
        fmt.Print(" ", n)
    }
}