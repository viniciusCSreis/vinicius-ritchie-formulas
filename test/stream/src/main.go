// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	input1 := os.Getenv("INPUT_TEXT")
	input2 := os.Getenv("INPUT_LIST")
	input3, _ := strconv.ParseBool(os.Getenv("INPUT_BOOLEAN"))
	input4 := os.Getenv("INPUT_PASSWORD")
	input5, err := strconv.Atoi(os.Getenv("INPUT_REPEAT"))
	if err != nil {
		panic(err)
	}
	input6, err := strconv.Atoi(os.Getenv("INPUT_REPEAT_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	sleep := time.Second * time.Duration(rand.Int31n(int32(input6)))

	formula.Formula{
		Text:     input1,
		List:     input2,
		Boolean:  input3,
		Password: input4,
		Repeat:   input5,
		Sleep:    sleep,
	}.Run(os.Stdout)
}
