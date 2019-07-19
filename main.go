package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

// ConfigField contains all information needed to define one piece of data that we are generating
type ConfigField struct {
	Name string
	Type string
	Min  int
	Max  int
}

// Config contains all information that is needed to generate test data
type Config struct {
	Title  string
	Output string
	Total  int
	Fields []ConfigField
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

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

	// fmt.Printf("Title: %#v\n", config.Title)
	// fmt.Printf("Output: %#v\n", config.Output)
	// fmt.Printf("Fields: %#v\n", config.Fields)

	generatedData := make([]map[string]string, config.Total)

	for i := 1; i < config.Total; i++ {
		values := make(map[string]string)

		for _, field := range config.Fields {
			var value string
			if field.Type == "alpha" {
				value = RandomString(field.Min, field.Max)
			} else if field.Type == "int" {
				value = strconv.Itoa(RandomNumber(field.Min, field.Max))
			} else {
				panic("Field type not supported" + field.Type)
			}
			values[field.Name] = value
		}
		generatedData[i-1] = values
	}

	fmt.Println(generatedData)

}

// RandomNumber returns an int between and including min and max
func RandomNumber(min, max int) int {
	return min + rand.Intn((max+1)-min)
}

// RandomString returns a string with random characters that is between min and max length
func RandomString(min, max int) string {
	len := RandomNumber(min, max)
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
