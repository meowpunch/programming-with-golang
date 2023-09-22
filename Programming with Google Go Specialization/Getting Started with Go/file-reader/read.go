package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Fname string
	Lname string
}

func ReadNamesFromFile(filename string) ([]User, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ' '

	var names []User

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		name := User{
			Fname: strings.TrimSpace(record[0]),
			Lname: strings.TrimSpace(record[1]),
		}
		names = append(names, name)
	}

	return names, nil
}

func ScanFileName() string {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scanln(&filename)
	return filename
}

func PrintNames(names []User) {
	for _, name := range names {
		fmt.Println(name.Fname, name.Lname)
	}
}

func main() {
	filename := ScanFileName()

	names, err := ReadNamesFromFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	PrintNames(names)
}
