package main

import (
	"fmt"
	"os"
	"strconv"
)

type portNum uint16

func newPortNum(p string) portNum {
	n, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("err: bad port number provied")
		os.Exit(1)
	}
	return portNum(n)
}

func (p portNum) String() string {
	return strconv.Itoa(int(p))
}

func (p *portNum) next() string {
	*p++
	return p.String()
}
