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

/*
func dirCheck() {
    if _, err := os.Stat("~/Bughunt/Bugxss"); os.IsNotExist(err) {
        os.Mkdir("~/Bughunt/Bugxss", 0755)

        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Print("Your vuln den was created")

    }
https://www.td.com/
    if _, err := os.Stat("~/Bughunt/Bugxss/domain"); os.IsNotExist(err) {
        os.Mkdir("~/Bughunt/Bugxss/domain", 0755)
        if err != nil {
            fmt.Printf("%s", err)
        }

    }
    if _,
    err := os.Stat("~/Bughunt/Bugxss/domain/xss"); os.IsNotExist(err) {
        os.Mkdir("~/Bughunt/Bugxss/domain/xss", 0755)
        if err != nil {
            fmt.Printf("%s", err)
        }

    }

}
*/

func gau(wg *sync.WaitGroup) {

	defer wg.Done()

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

func httpx(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/local/bin/httpx"); os.IsNotExist(err) {
		out,
			err := exec.Command("httpx", "%s", "-silent", domain).Output()
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

	var wg sync.WaitGroup

	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
	} else {

		domName()
		//dirCheck()
		wg.Add(2)
		go gau(&wg)
		//go httpx(&wg)
		go assetf(&wg)
		//go subF(&wg)
		//go amassF(&wg)
		wg.Wait()
		fmt.Println("Finished all the task")
		out := subSort(output)
		fmt.Println(out)
		fmt.Println("Done!")

	}

}
