package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
		return
	}

	var prevOutput []string
	domain := getDomainInput()
	var wg sync.WaitGroup
	wg.Add(6)

	results := make(chan string, 100) // Buffered channel to store results
	// Start tools concurrently
	go runTool(&wg, "gau", []string{"--subs", domain}, results)
	go runTool(&wg, "httpx", []string{"--no-color"}, results, prevOutput) // httpx needs special handling
	go runTool(&wg, "assetfinder", []string{"-subs-only", domain}, results)
	go runTool(&wg, "amass", []string{"enum", "--passive", "-d", domain}, results)
	go runTool(&wg, "subfinder", []string{"-d", domain}, results)
	go runTool(&wg, "/Umbrabug/Sublist3r python3 sublist3r.py", []string{"-b", "-d", domain}, results)

	wg.Wait()
	close(results) // Close channel after all writes are done

	// Collect and sort results
	var finalOutput []string
	for result := range results {
		finalOutput = append(finalOutput, result)
	}
	finalOutput = unique(finalOutput)

	// Write results to CSV
	csvFileName := fmt.Sprintf("%s.csv", domain)
	if err := writeToCSV(csvFileName, finalOutput); err != nil {
		fmt.Printf("Error writing CSV file: %s\n", err)
	} else {
		fmt.Printf("Saved output to CSV file: %s\n", csvFileName)
	}
}

func getDomainInput() string {
	var domain string
	fmt.Print("Enter your domain: ")
	fmt.Scan(&domain)
	return domain
}

func runTool(wg *sync.WaitGroup, command string, args []string, results chan<- string, prevOutput []string) {
	defer wg.Done()

	// Check if the command exists
	cmdPath, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s command not found\n", command)
		return
	}

	var cmd *exec.Cmd
	// Special handling for tools that require input from a file or stdin
	if command == "httpx" && len(prevOutput) > 0 {
		// Create a temporary file
		tmpFile, err := ioutil.TempFile("", "httpx-input")
		if err != nil {
			fmt.Printf("Error creating temp file for httpx input: %s\n", err)
			return
		}
		defer os.Remove(tmpFile.Name()) // Clean up the file after

		// Write the previous output to the temp file
		for _, line := range prevOutput {
			if _, err := tmpFile.WriteString(line + "\n"); err != nil {
				fmt.Printf("Error writing to temp file for httpx: %s\n", err)
				return
			}
		}
		tmpFile.Close() // Close the file so httpx can read it

		// Adjust the command to read from the temp file
		args = append(args, "-l", tmpFile.Name())
		cmd = exec.Command(cmdPath, args...)
	} else {
		cmd = exec.Command(cmdPath, args...)
	}

	// Execute the command
	output, err := cmd.CombinedOutput() // Use CombinedOutput to get both stdout and stderr
	if err != nil {
		fmt.Printf("Error running %s: %s\n", command, err)
		return
	}

	if command == "httpx" {
		// Process httpx output specifically if needed
	} else {
		// General processing for other tools
		results <- string(output)
	}
}

// In the main or calling function, ensure to collect the previous outputs into a slice before calling runTool for httpx
// Assume this is filled with outputs from previous tools

func unique(strings []string) []string {
	seen := make(map[string]struct{})
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
