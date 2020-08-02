package main

import (
	"flag"
	"io"
	"fmt"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var From = flag.String("from", "D:\\Go\\golang\\HW-9-11\\test.txt", "Original file")
var To = flag.String("to", "D:\\Go\\golang\\HW-9-11\\test-2.txt", "Destination file")
var Offset = flag.Int64("offset", 0, "Seek relative to the origin of the file")
var Limit = flag.Int64("limit", 0, "Stops with EOF after limit bytes")

func Copy(from string, to string, limit int, offset int) {

	flag.Parse()

	//Copy(from string, to string, limit int, offset int) error
	//D:\Go\golang\HW-9-11\test.txt

	f, err := os.Open(*From)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := os.Stat(*From)
	fmt.Println(b.Size())

	if *Limit == 0 {
		*Limit = b.Size()
	}

	reader := io.NewSectionReader(f, *Offset, *Limit)
	writer, _ := os.Create(*To)

	// start new bar
	bar := pb.Full.Start64(*Limit)
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	io.Copy(writer, barReader)
	// finish bar
	bar.Finish()

}

func main() {
	Copy("D:\\Go\\golang\\HW-9-11\\test.txt", "D:\\Go\\golang\\HW-9-11\\test-2.txt", 0, 0)

}
