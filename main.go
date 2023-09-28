package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func readDirectoriesFromFile(filename string) ([]string, error) {
	var directories []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		directories = append(directories, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return directories, nil
}
func fuzzDirectories(directories []string) {
	for _, dir := range directories {
		url := "https://att.com/" + dir // Replace with your target URL

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Status Code: %d\n", resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
		} else {
			fmt.Printf("Response Length: %d bytes\n", len(body))
		}
	}
}
func main() {
	filename := "directories.txt" // Replace with your input file name
	directories, err := readDirectoriesFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading directories from file: %v\n", err)
		return
	}

	fuzzDirectories(directories)
}
