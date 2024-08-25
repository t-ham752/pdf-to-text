package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %+v", err)
	}
	defer func() { _ = file.Close() }()

	cmd := exec.Command("pdftotext", "-", "-")
	cmd.Stdin = file

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if err := cmd.Run(); err != nil {
		log.Fatalf("pdftotext failed: %+v", err)
	}

	fmt.Printf("output: %s\n", buffer.String())
}
