package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/Bellese/test-case-engine/models"
	"github.com/ghodss/yaml"
)

// ParseConfigs recieves a filename and parses the corresponding yaml file
func ParseConfigs(filename string) models.Config {
	fmt.Println("Reading File: " + filename)

	file, err := os.Open(filename)
	data, err := ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		panic("ERROR OPENING FILE:\n" + err.Error())
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("ERROR PARSING YAML:\n" + err.Error())
	}

	return config
}

// GenerateData will create an array of data
func GenerateData(config models.Config) []map[string]string {
	data := make([]map[string]string, config.Total)

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
		data[i-1] = values
	}

	return data
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
