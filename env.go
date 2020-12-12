package main

import (
	"fmt"
	"os"
)

func main() {
	str:= os.Getenv("hh")
	fmt.Println("哈哈",str,"end")
}
