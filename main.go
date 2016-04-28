package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var prefix = ""

func main() {
	flag.StringVar(&prefix, "prefix", prefix, "insert this text after the timestamp but before the input line")
	flag.Parse()
	if prefix != "" {
		prefix = prefix + " "
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	stdin := bufio.NewReader(os.Stdin)
	for {
		if line, err := stdin.ReadString('\n'); err == nil {
			log.Print(fmt.Sprintf("%s%s", prefix, line))
		} else {
			if line != "" {
				log.Print(fmt.Sprintf("%s%s", prefix, line))
			}
			if err != io.EOF {
				log.Fatal(err)
			}
			return
		}
	}
}
