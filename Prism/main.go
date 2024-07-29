package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/pavank1301/Prism/repl"
)

func main(){
	user, err := user.Current()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Hello %s!\nThis is the Prism programming language!\n",
user.Username)
	repl.Start(os.Stdin, os.Stdout)
}