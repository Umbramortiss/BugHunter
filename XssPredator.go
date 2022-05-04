package main

import(
    "fmt"
    "os/exec"
    "os"
    "runtime"
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

func subExec() {
    //
    if _,
    err := os.Stat("/usr/bin/gau"); os.IsNotExist(err) {
        out,
        err := exec.Command("gau", "%s", "--subs", domain).Output()

        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Println("Running gau for fetching URLs")
        output := string(out[:])
        fmt.Println(output)

    }
    
    if _,
    err := os.Stat("/usr/bin/assetfinder"); os.IsNotExist(err) {
        out,
        err := exec.Command("assetfinder", "-subs-only", "%s",  domain).Output()
        
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Println("Running assetfinder for domain enumeration")
        output := string(out[:])
        fmt.Println(output)
    }
    
    if _,
    err := os.Stat("/usr/bin/subfinder"); os.IsNotExist(err) {
        out,
        err := exec.Command("subfinder", "-d","%s", domain ).Output()
        
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Println("")
    }


}

func main() {

    if runtime.GOOS == "windows" {
        fmt.Println("Can't execute this on a windows machine")
    } else {
        domName()
      /*  dirCheck() */
        subExec()

    }




}