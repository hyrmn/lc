package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hyrmn/lc/pkg/lc"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Mode()&os.ModeCharDevice == 0 {
		reader := bufio.NewReader(os.Stdin)
		countLines(reader)
	} else {
		flag.Parse()
		filePath := flag.Arg(0)

		if filePath == "" {
			fmt.Println("Usage:\n\tlc \"path\\to\\file.txt\"")
			return
		}

		file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		countLines(file)
	}
}

func countLines(r io.Reader) {
	count, err := lc.CountLines(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
