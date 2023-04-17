package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"io/ioutil"
)

var domain string
var output []string
var liveOutput []string


func domName() {
	fmt.Print("Enter your domain: ")
	fmt.Scanf("%s", &domain)
}

func installTool(tool string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("which %s", tool))
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s not found, installing...\n", tool)
		switch tool {
		case "subfinder":
			if runtime.GOOS == "windows" {
				fmt.Println("Subfinder installation is not supported on Windows")
				return
			}
			if err := exec.Command("sh", "-c", "curl -sL https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin v2.2.0").Run(); err != nil {
				fmt.Printf("Error installing subfinder: %s\n", err)
				return
			}
		case "dirsearch":
			if err := exec.Command("python3", "-m", "pip", "install", "--user", "dirsearch").Run(); err != nil {
				fmt.Printf("Error installing dirsearch: %s\n", err)
				return
			}
		case "nmap":
			if err := exec.Command("sudo", "apt-get", "install", "-y", "nmap").Run(); err != nil {
				fmt.Printf("Error installing nmap: %s\n", err)
				return
			}
		}
		fmt.Printf("%s installed successfully\n", tool)
	}
}



func dirCheck() {
    if _, err := os.Stat("$HOME/Bughunt/Bugxss"); os.IsNotExist(err) {
        os.Mkdir("$HOME/Bughunt/Bugxss", 0755)

        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Print("Your vuln den was created")

    }
//https://www.td.com/
    if _, err := os.Stat("$HOME/Bughunt/Bugxss/domain"); os.IsNotExist(err) {
        os.Mkdir("$HOME/Bughunt/Bugxss/domain", 0755)
        if err != nil {
            fmt.Printf("%s", err)
        }

    }
    if _,
    err := os.Stat("$HOME/Bughunt/Bugxss/domain/xss"); os.IsNotExist(err) {
        os.Mkdir("$HOME/Bughunt/Bugxss/domain/xss", 0755)
        if err != nil {
            fmt.Printf("%s", err)
        }

    }

}


func gau(wg *sync.WaitGroup) {

	defer wg.Done()

	if _, err := os.Stat("/usr/local/bin/gau"); os.IsNotExist(err) {

		out, err := exec.Command("gau", "--subs", domain).Output()
		fmt.Println(out)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Running gau for fetching URLs")
		output = append(output, string(out[:]))
		fmt.Println(output)

	}

}

func httpx(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/local/bin/httpx"); os.IsNotExist(err) {
			file, err := ioutil.TempFile("","httpx")
			if err =! nil {
				fmt.Printf("Error creating temp file: %s", err)
				return
			}
			defer os.Remove(file.Name())

			if _, err := file.WriteString(strings.Join(output, "\n")); err != nil {
				fmt.Printf("Error writing to temp file: %s", err)
				return
			}
		out,
			err := exec.Command("httpx", "-l",file.Name(), "--no-color", domain).Output()
		fmt.Println(err)
		if err != nil {
			fmt.Printf("%s", err)

		}
		fmt.Println("Running httpx for live URLs")
		liveOutput := string(out[:])
		fmt.Println(liveOutput)

	}

}
func assetf(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/assetfinder"); os.IsNotExist(err) {
		out,
			err := exec.Command("assetfinder", "-subs-only", domain).Output()

		if err != nil {
			fmt.Printf("Error running assetfinder: %s", err)
		}
		fmt.Println("Running assetfinder for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	}
}
func amassF(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/amass"); os.IsNotExist(err) {
		out,
			err := exec.Command("amass", "enum", "--passive", "-d", domain).Output()

		if err != nil {
			fmt.Printf("%f", err)
		}
		fmt.Println("Running amass for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	}
}
func subF(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := exec.LookPath("subfinder"); err != nil {
		out,
			err := exec.Command("subfinder", "-d", domain).Output()

		if err != nil {
			fmt.Printf("Error running subfinder: %s", err)
		}
		fmt.Println("Running subfinder for domain enumeration")
		output = append(output, string(out[:]))
		fmt.Println(output)
	}
}

func subList3r(wg *sync.WaitGroup){
	defer wg.Done()

	if _, err := os.Stat("/usr/bin/sublist3r"); os.IsNotExist(err){
		out, err := exec.Command("sublist3r", "-b", "-d", domain).Output()

		if err != nil {
			fmt.Printf("Error running sublist3r: %s", err)
		}
		fmt.Println("Running sublist3r for domain enumeration")
		output = append(output, string(out[:]))
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

func writeToCSV(fileName string, data []string)error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, file := range data {
		if err := w.Write([]string{file});err != nil{
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
		dirCheck()
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
		}else{
			fmt.Println("Saved output to CSV file:",csvFileName )
		}

		fmt.Println("Done!")

	}

}
