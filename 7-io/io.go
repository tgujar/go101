package main

import (
	"bufio" // buffered io
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func readEntireFile(filename string) {
	txt, err := os.ReadFile(filename) // txt is of type []byte
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(txt))

}

func readLineByLine(filename string) {
	f, err := os.Open(filename)
	if err != nil { // nil is the zero value for pointers, interfaces, maps, slices, channels and function types
		log.Fatal(err)
	}

	defer f.Close() // defer should be after err check

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readInChunks(filename string, chunkSize int) {
	f, err := os.Open(filename) // Open is read only
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	buf := make([]byte, chunkSize)
	reader := bufio.NewReader(f)
	result := bytes.NewBuffer(nil)

	for {
		n, err := reader.Read(buf) // read up to len(buf) bytes, maybe lesser
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Println(string(buf[:n])) // Ouptut bytes read
		result.Write(buf[:n])        // write to byte buffer
	}
	fmt.Println("Complete message:", string(result.Bytes())) // full output
}

func writeBuffered(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // might cause issues -> https://www.joeshaw.org/dont-defer-close-on-writable-files/

	writer := bufio.NewWriter(f)
	b := []byte("\nBuffered!")
	_, err = writer.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush() // flush the buffer to the file
}

func flushExample(f *os.File) {
	writer := bufio.NewWriter(f)
	b := []byte("Flushed!\n")
	_, err := writer.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush() // write before sleeping (try commenting this out)
	time.Sleep(3 * time.Second)
}

func main() {
	dir, _ := os.Getwd()
	readEntireFile(path.Join(dir, "test.txt"))
	readLineByLine(path.Join(dir, "test.txt"))
	readInChunks(path.Join(dir, "test.txt"), 5)
	writeBuffered(path.Join(dir, "test.txt"))
	flushExample(os.Stdout)
}
