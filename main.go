package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title      string
	Output     string
	Parameters []string
}

func main() {
	args := os.Args[1:]
	filename := args[0]

	fmt.Println("Reading File: " + filename)

	file, err := os.Open(filename)
	data, err := ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		panic("ERROR OPENING FILE:\n" + err.Error())
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("ERROR PARSING YAML:\n" + err.Error())
	}

	fmt.Printf("Title: %#v\n", config.Title)
	fmt.Printf("Output: %#v\n", config.Output)
	fmt.Printf("Parameters: %#v\n", config.Parameters)

}
