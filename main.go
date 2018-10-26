package main

import (
	"io/ioutil"
	"os"
)

func main() {
	initLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("bedrocksrv © JoshuaDoes: 2018.")
	Info.Println(BuildID)
}
