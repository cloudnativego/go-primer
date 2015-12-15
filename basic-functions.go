package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type dieRollFunc func(int) int

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func fakeDieRoll(size int) int {
	return 42
}

func rollTwo(size1 int, size2 int) (int, int) {
	return dieRoll(size1), dieRoll(size2)
}

func getDieRolls() []dieRollFunc {
	return []dieRollFunc{
		dieRoll,
		fakeDieRoll,
	}
}

func returnsNamed(input1 string, input2 int) (theResult string, err error) {
	theResult = "modified " + input1 + ", " + strconv.Itoa(input2)
	return theResult, err
}

func main() {
	fmt.Printf("Rolled a die of size %d, result: %d\n", 6, dieRoll(6))

	res1, res2 := rollTwo(6, 10)
	fmt.Printf("Rolled a pair of dice (%d,%d), results: %d, %d\n",
		6, 10, res1, res2)

	named, err := returnsNamed("globule", 42)
	fmt.Printf("Named params returned: '%s', %v\n", named, err)

	var rolls = getDieRolls()
	d10 := rolls[0](10)
	otherd10 := rolls[1](10)
	fmt.Printf("Die rolls invoked in slice of func ptrs: %d %d", d10, otherd10)
}
