package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var domain string
var output []string

func domName() {
	fmt.Print("Enter your domain: ")
	fmt.Scanf("%s", &domain)
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
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
} 

func main() {


	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
	} else {

		domName()
		gau()
		fmt.Println("Finished all the task")
		out := subSort(output)
		fmt.Println(out)
		fmt.Println("Done!")

	}

}
