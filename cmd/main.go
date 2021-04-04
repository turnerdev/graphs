package main

import (
	"flag"
	"fmt"
	svg "graphs/pkg"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] output.svg ...\n", os.Args[0])
	flag.PrintDefaults()
}

func render(f *os.File) {
	root := svg.New()
	root.Rect(0, 0, 100, 100).Rect(150, 150, 300, 300)

	n, err := root.Write(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	f, err := os.Create(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	render(f)
}
