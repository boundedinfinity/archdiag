package main

import (
	"fmt"
	"os"

	"github.com/bounded-infinity/archdiag"
)

func main() {
	d := archdiag.ArchDiag{}

	if err := d.ReadFromFile("../test/g1.json"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := d.WriteToStdout(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
