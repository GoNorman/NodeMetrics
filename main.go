package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num = rand.Intn(10) + 1
	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "in the second")

	distance = 401000000

	fmt.Println(distance/lightSpeed, "in the second")

	fmt.Println(num)

	now := time.Now()

	fmt.Println(now)
}
