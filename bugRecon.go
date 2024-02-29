package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var domain string
var output []string
var liveOutput []string

func domName() {
	fmt.Print("Enter your domain: ")
	fmt.Scanf("%s", &domain)
}

func gau(wg *sync.WaitGroup) {

	defer wg.Done()

	if _, err := os.Stat("/usr/bin/gau"); err == nil {

		out, err := exec.Command("gau", "--subs", domain).Output()
		fmt.Println(out)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Running gau for fetching URLs")
		output = append(output, string(out[:]))
		fmt.Println(output)

	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
	}

}

func httpx(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/httpx"); err == nil {
		file, err := ioutil.TempFile("", "httpx")
		if err != nil {
			fmt.Printf("Error creating temp file: %s", err)
			return
		}
		defer os.Remove(file.Name())

		if _, err := file.WriteString(strings.Join(output, "\n")); err != nil {
			fmt.Printf("Error writing to temp file: %s", err)
			return
		}
		out,
			err := exec.Command("httpx", "-l", file.Name(), "--no-color", domain).Output()
		fmt.Println(err)
		if err != nil {
			fmt.Printf("%s", err)

		}
		fmt.Println("Running httpx for live URLs")
		liveOutput := string(out[:])
		fmt.Println(liveOutput)

	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
	}

}
func assetf(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/assetfinder"); err == nil {
		out,
			err := exec.Command("assetfinder", "-subs-only", domain).Output()

		if err != nil {
			fmt.Printf("Error running assetfinder: %s", err)
		}
		fmt.Println("Running assetfinder for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
	}
}
func amassF(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/amass"); err == nil {
		out,
			err := exec.Command("amass", "enum", "--passive", "-d", domain).Output()

		if err != nil {
			fmt.Printf("%f", err)
		}
		fmt.Println("Running amass for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
	}
}
func subF(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := exec.LookPath("subfinder"); err == nil {
		out,
			err := exec.Command("subfinder", "-d", domain).Output()

		if err != nil {
			fmt.Printf("Error running subfinder: %s", err)
		}
		fmt.Println("Running subfinder for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
	}
}

func subList3r(wg *sync.WaitGroup) {
	defer wg.Done()

	if _, err := os.Stat("/usr/bin/sublist3r"); err == nil {
		out, err := exec.Command("sublist3r", "-b", "-d", domain).Output()

		if err != nil {
			fmt.Printf("Error running sublist3r: %s", err)
		}
		fmt.Println("Running sublist3r for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	} else if os.IsNotExist(err) {
		fmt.Println("gau command not found")
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

func writeToCSV(fileName string, data []string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, file := range data {
		if err := w.Write([]string{file}); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	var wg sync.WaitGroup

	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
	} else {

		domName()
		wg.Add(6)
		go gau(&wg)
		go httpx(&wg)
		go assetf(&wg)
		go subF(&wg)
		go amassF(&wg)
		go subList3r(&wg)
		wg.Wait()
		fmt.Println("Finished all the task")
		out := subSort(liveOutput)
		fmt.Println(out)
		csvFileName := domain + ".csv"
		if err := writeToCSV(csvFileName, out); err != nil {
			fmt.Println("Error writing CSV file:", err)
		} else {
			fmt.Println("Saved output to CSV file:", csvFileName)
		}

		fmt.Println("Done!")

	}

}
