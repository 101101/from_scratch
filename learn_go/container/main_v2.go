package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//docker run <container> cmd args
// go run main.go run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what??")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}os.Args[2]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

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

// go run main_v2.go run echo container camp  
// then try go run main_v2.go /bin/bash  
// you get a shell inside the go container  
// ps still shows host PIDs
// to get segregated PIDs, you have to run in a new rootfs   
// moving on to v3  

