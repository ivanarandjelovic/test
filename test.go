package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Man struct {
	Name     string
	LastName string
	Alive    bool
}

type Configuration struct {
	Users  []string
	Groups []string
	People []Man
}

/**
 * Load configuration object.
 */
func loadConfiguration() (configuration Configuration, err error) {
	file, _ := os.Open("test_conf.json")
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return
}

func main() {
	fmt.Println("Test v3.0")

	configuration, _ := loadConfiguration()

	fmt.Println(configuration.Users) // output: [UserA, UserB]
	fmt.Println(configuration.People)

}
