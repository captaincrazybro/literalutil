package main

import (
	"fmt"
	"github.com/captaincrazybro/literalutil"
)

func main() {
	s := lu.String("test")
	fmt.Printf("%q", s.Replace("e", "_", 1))
}