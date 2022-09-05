// A MySQL binary log analyzer in Golang.
// Pass in the binary log file and read Insert, Update, Delete, Create, Drop, Truncate, Alter from a MySQL binary log file and print the events.
// I modified the code in this repo: https://github.com/liipx/go-mysql-binlog

package main

import (
	"flag"
	"fmt"

	binlog "github.com/ChaosHour/go-binlog/cmd/go-binlog"
)

// Define flags
var (
	file = flag.String("f", "", "MySQL binary log file")
	help = flag.Bool("h", false, "Print help")
)

// parse flags
func init() {
	flag.Parse()
}

// main function
func main() {
	flag.Parse()
	// make sure the file flag is set
	if *file == "" {
		fmt.Println("Please specify a MySQL binary log file")
		flag.PrintDefaults()
		return
	}
	if *help {
		flag.PrintDefaults()
		return
	}
	decoder, err := binlog.NewBinFileDecoder(*file)
	if err != nil {
		panic(err)
	}

	num := 0
	maxEventCount := 0
	err = decoder.WalkEvent(func(event *binlog.BinEvent) (isContinue bool, err error) {
		fmt.Printf("Got %s: \n\t", binlog.EventType2Str[event.Header.EventType])
		fmt.Println(event.Header)
		fmt.Println(event.Body)

		num++
		if num > maxEventCount {
			maxEventCount = num
		}
		return true, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d events\n", maxEventCount)
}
