package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var domain string

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

	if _,
		err := os.Stat("/usr/bin/gau"); os.IsNotExist(err) {
		out,
			err := exec.Command("gau", "%s", "--subs", domain).Output()

		if err != nil {
			fmt.Printf("%f", err)
		}
		fmt.Println("Running gau for fetching URLs")
		urls := string(out[:])
		fmt.Println(urls)

	}
	
}

func httpx(wg *sync.WaitGroup) {

	defer wg.Done()

	if _,
		err := os.Stat("/usr/bin/httpx"); os.IsNotExist(err) {
		out,
			err := exec.Command("httpx", "%s", "-f urls", "-silent", domain).Output()

		if err != nil {
			fmt.Printf("%f", err)

		}
		fmt.Println("Running httpx for live URLs")
		liveOutput := string(out[:])
		fmt.Println(liveOutput)

	}

}

/*

   func subSort(s []string) []string {
       inResult := make(map[string]bool)
       var result [output,output_2,output_3,output_4]
       for _, str := rang s {
           if _, ok := inResult[str]; !ok  {
               inResult[str] = true
               result = append(result, str)
           }
       }
       return fmt.Printf(result)
   }


*/

func main() {

	var wg sync.WaitGroup

	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
	} else {

		domName()
		/*  dirCheck() */
		wg.Add(2)
		go gau(&wg)
		go httpx(&wg)

		fmt.Println("sorting or filtering results...")
		wg.Wait()
		fmt.Println("Done!")

	}

}
