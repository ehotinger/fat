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
	log.Println("Hello world!")
	log.Println("You can invoke this binary without any parameters to perform a 'slim' read, or pass in '--file /usr/bin/file.txt' to perform a 'fat' read")
	file := flag.String("file", "", "the file path")
	flag.Parse()

	if *file != "" {
		log.Println("Reading the entire file...")
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
		log.Println("Read through the entire file!")

	} else {
		log.Println("Skipped reading files")
	}
}
