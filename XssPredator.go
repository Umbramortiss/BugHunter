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
func gau(){
    if _,
    err := os.Stat("/usr/bin/gau"); os.IsNotExist(err) {
        out,
        err := exec.Command("gau", "%s", "--subs", domain).Output()

        if err != nil {
            fmt.Printf("%f", err)
        }
        fmt.Println("Running gau for fetching URLs")
        output := string(out[:])
        fmt.Println(output)

    }
}
    
func assetF() {
    if _,
    err := os.Stat("/usr/bin/assetfinder"); os.IsNotExist(err) {
        out,
        err := exec.Command("assetfinder", "-subs-only", "%s",  domain).Output()
        
        if err != nil {
            fmt.Printf("%f", err)
        }
        fmt.Println("Running assetfinder for domain enumeration")
        output_2 := string(out[:])
        fmt.Println(output_2)
    }
}
func subF() {
    if _,
    err := os.Stat("/usr/bin/subfinder"); os.IsNotExist(err) {
        out,
        err := exec.Command("subfinder", "-d","%s", domain ).Output()
        
        if err != nil {
            fmt.Printf("%f", err)
        }
        fmt.Println("Running subfinder for domain enumeration")
        output_3 := string(out[:])
        fmt.Println(output_3)
    }
}
func amass() {
    if _,
    err := os.Stat("/usr/bin/amass"); os.IsNotExist(err) {
        out,
        err := exec.Command("amass", "enum", "--passive", "-d", domain).Output()
        
        if err != nil {
            fmt.Printf("%f", err)
        }
        fmt.Println("Running amass for domain enumeration")
        output_4 := string(out[:])
        fmt.Println(output_4)
    }
}
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