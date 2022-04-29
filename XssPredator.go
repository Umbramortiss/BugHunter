
package main

import(
 "fmt"
 "os/exec"
 "runtime"
)

var domain string

func dirCheck(){
    fileinfo, err := os.Stat("~/Bughunt/Bugxss")
    if os.IsNotExist(err) {
        os.mkdir("~/Bughunt/Bugxss")
        
        if err != os.Stat("") {
            fmt.Printf("%s", err)
        }
        fmt.Print("Your vuln den was created")
    }
    if _, err := os.Stat("~/Bughunt/Bugxss/domain"); os.IsNotExist(err){
        os.mkdir("~/Bughunt/Bugxss/domain")
            if err != nil {
            fmt.Printf("%s", err)
            }
        
    }
    fileinfo, err := os.stat("~/Bughunt/Bugxss/domain/xss")
    if os.IsNotExist(err) {
        os.mkdir("~/Bughunt/Bugxss/domain/xss")
            if err != os.Stat("") {
            fmt.Printf("%s", err)
            }
    }
}

func execute(){
    //
    fileinfo, err := os.stat("/usr/bin/gau")
    if os.IsNotExist(err) {
        out, err := exec.Command("gau","--subs").Output()
        
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Println("Running gau for fetching URLs")
        output := string(out[:])
        fmt.Println(output)
    }
    
    
}

func main() {
    
    if runtime.GOOS == "windows" {
        fmt.Println("Can't execute this on a windows machine")
    } else {
    fmt.Print("Enter your domain: ")
    fmt.Scanf("%s" &domain)
    dirCheck()
    execute()
        
    }
    
    
    

}

