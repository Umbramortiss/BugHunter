package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
)

var domain string
var output []string

func domName() {
	fmt.Print("Enter your domain: ")
	fmt.Scan(&domain)
	return domain
}

func gau() {

	if _, err := os.Stat("/usr/local/bin/gau"); os.IsNotExist(err) {

		out, err := exec.Command("gau", "%s", "--subs", domain).Output()
		fmt.Println(out)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Running gau for fetching URLs")
		output := string(out[:])
		fmt.Println(output)

	}

}
func subSort(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range strings {
		if _, ok := seen[str]; !ok {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

func writeToCSV(fileName string, data []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write([]string{record}); err != nil {
			return err
		}
	}
	return nil
}
