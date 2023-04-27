package main

import (
	"fmt"
	shadow "github.com/PaiGack/shadow-go"
)

func main() {
	fmt.Println("branch", shadow.Branch())
	fmt.Println("tag", shadow.Tag())
	fmt.Println("git_clean", shadow.GitClean())
	fmt.Println("git_status_file", shadow.GitStatusFile())
}
