// For creating the file:
// dd if=/dev/urandom of=file.txt bs=64M count=64 iflag=fullblock

// Example Docker commands:
// docker build -t fat .
// docker run fat -read -file /usr/bin/file.txt

package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	read := flag.Bool("read", false, "to enable read")
	file := flag.String("file", "file.txt", "the file path")
	flag.Parse()

	if *read {
		log.Println("Big read incoming...")
		file, err := os.Open(*file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			_ = scanner.Text() // Discard the allocation
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		log.Println("Big read finished")

	} else {
		log.Println("Slim")
	}
}
