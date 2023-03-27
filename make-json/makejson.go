package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ScanNameAndAddress() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	// Read name
	fmt.Print("Enter a name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Read address
	fmt.Print("Enter an address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	return name, address
}
func CreateJSON(name string, address string) ([]byte, error) {
	// Create a map with the name and address
	info := make(map[string]string)
	info["name"] = name
	info["address"] = address

	// Convert the map to JSON
	return json.Marshal(info)
}

func main() {
	name, address := ScanNameAndAddress()

	jsonData, err := CreateJSON(name, address)
	if err != nil {
		fmt.Println("Error creating JSON object:", err)
		return
	}

	fmt.Println(string(jsonData))
}
