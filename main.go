package main

import (
	"io"
	"log"
	"net/url"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("url")

	kingpin.Command("escape", "Escape data from stdin to stdout")
	kingpin.Command("unescape", "Unescape data from stdin to stdout")

	cmd := kingpin.Parse()

	var s string
	b := make([]byte, 16384)

	for {
		n, err := os.Stdin.Read(b)
		if n > 0 {
			switch cmd {
			case "escape":
				s = url.QueryEscape(string(b[:n]))
			case "unescape":
				s, err = url.QueryUnescape(string(b[:n]))
				if err != nil {
					log.Fatal(err)
				}
			}

			if _, err := os.Stdout.Write([]byte(s)); err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
