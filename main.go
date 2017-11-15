package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	steps := 10
	num := rand.Intn(100)
	fmt.Println(num)
	A := []int{}
	B := []int{}
	for i := 1; i <= steps; i++ {
		A = append(A, num)
		B = append(B, num)
	}

	fmt.Println("A: ", A)
	fmt.Println("B: ", B)
	fmt.Println("Len:", len(A))

	start := time.Now()
	r := mult(A, B)

	elapsed := time.Since(start)
	log.Printf("Mult took %s", elapsed)
	fmt.Println("Result: ", r)

}

func mult(a []int, b []int) []int {
	var result []int
	for k, v := range a {
		fmt.Println("k:", k)
		for i, n := range b {
			fmt.Println("i: ", i)
			result[k] = v * n
		}
	}

	return result
}
