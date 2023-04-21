package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	deltas = false
)

func init() {
	flag.BoolVar(&deltas, "deltas", deltas, "print deltas")
}

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	stdin := bufio.NewReader(os.Stdin)
	now := time.Now()
	start := now
	for {
		delta := ""
		if line, err := stdin.ReadString('\n'); err == nil {
			if deltas {
				newNow := time.Now()
				delta = fmt.Sprintf("d:%f D:%f ", newNow.Sub(now).Seconds(), newNow.Sub(start).Seconds())
				now = newNow
			}
			log.Print(fmt.Sprintf("%s%s", delta, line))
		} else {
			if line != "" {
				if deltas {
					newNow := time.Now()
					delta = fmt.Sprintf("d:%f D:%f ", newNow.Sub(now).Seconds(), newNow.Sub(start).Seconds())
					now = newNow
				}
				log.Print(fmt.Sprintf("%s%s", delta, line))
			}
			if err != io.EOF {
				log.Fatal(err)
			}
			return
		}
	}
}
