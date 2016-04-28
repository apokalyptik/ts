package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	stdin := bufio.NewReader(os.Stdin)
	for {
		if line, err := stdin.ReadString('\n'); err == nil {
			log.Print(line)
		} else {
			if line != "" {
				log.Print(line)
			}
			if err != io.EOF {
				log.Fatal(err)
			}
			return
		}
	}
}
