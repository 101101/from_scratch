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

	must(syscall.Chroot("/home/rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// progressing as if you have a rootfs in folder "/home/rootfs" on the host machine  
// need to also mount /proc to make ps work  
// then try go run main_v3.go /bin/bash  
// you get a shell inside the go container  
// you can now see the "basic" isolated container, run ps to get your own process list, and see your own filesystem in /home/rootfs  
// will need to address: Users, IPC, and Networking to ship this solution...  


// Additional notes - microbadger.com - allow you to inspect layers in continers  

