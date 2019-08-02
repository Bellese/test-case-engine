package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Bellese/test-case-engine/utils"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	filename := args[0]

	config := utils.ParseConfigs(filename)

	generatedData := utils.GenerateData(config)

	fmt.Println(generatedData)

	utils.FormatData(config, generatedData)
}
