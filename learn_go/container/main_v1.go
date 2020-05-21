package main

import (
	"fmt"
	"os"
	"os/exec"
)

//docker run <container> cmd args
// go run main.go run cmd args
func main() {
	switch os.Args[1] {
		case "run";
			run()
	default:
		panic("what??")
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// go run main_v1.go run echo container camp  
// then try go run main_v1.go /bin/bash  
// you get a shell inside the go container  
// at this point, the container can control the host  
// moving on to v2  

