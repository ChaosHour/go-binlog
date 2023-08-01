// A MySQL binary log analyzer in Golang.
// Pass in the binary log file and read Insert, Update, Delete, Create, Drop, Truncate, Alter from a MySQL binary log file and print the events.
// I modified the code in this repo: https://github.com/liipx/go-mysql-binlog

/*
Copyright 2023 Kurt Larsen (ChaosHour)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/fatih/color"

	binlog "github.com/ChaosHour/go-binlog/cmd/go-binlog"
)

// Define flags
var (
	file = flag.String("f", "", "MySQL binary log file")
	//database    = flag.String("d", "", "MySQL database name")
	//table       = flag.String("t", "", "MySQL table name")
	startPos    = flag.Int("s", 0, "Start position")
	endPos      = flag.Int("e", 0, "End position")
	tableCounts = flag.Bool("c", false, "Table counts - experamental")
	help        = flag.Bool("h", false, "Print help")
)

// define colors
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

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
	tableCountsMap := make(map[string]int)
	err = decoder.WalkEvent(func(event *binlog.BinEvent) (isContinue bool, err error) {
		if *startPos == 0 && *endPos == 0 {
			fmt.Printf("Got %s: \n\t", binlog.EventType2Str[event.Header.EventType])
			fmt.Println(event.Header)
			fmt.Printf("LogPos: %d\n", event.Header.LogPos) // Print LogPos value
			fmt.Println(event.Body)

			num++
			if num > maxEventCount {
				maxEventCount = num
			}
			if *tableCounts {
				if event.Header.EventType == binlog.TableMapEvent {
					tableID := strconv.FormatUint(event.Body.(*binlog.BinTableMapEvent).TableID, 10)
					tableCountsMap[tableID]++
				}
			}
		} else if event.Header.LogPos >= int64(*startPos) && event.Header.LogPos <= int64(*endPos) {
			fmt.Printf("Got %s: \n\t", binlog.EventType2Str[event.Header.EventType])
			fmt.Println(event.Header)
			fmt.Printf("LogPos: %d\n", event.Header.LogPos) // Print LogPos value
			fmt.Println(event.Body)

			num++
			if num > maxEventCount {
				maxEventCount = num
			}
			switch body := event.Body.(type) {
			case *binlog.BinTableMapEvent:
				// body is of type *binlog.BinTableMapEvent
				tableID := strconv.FormatUint(body.TableID, 10)
				tableCountsMap[tableID]++
			}

		} else if event.Header.LogPos > int64(*endPos) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		panic(err)
	}

	// Get the current position in the binary log file
	pos, err := decoder.GetFilePos()
	if err != nil {
		panic(err)
	}

	// Convert the int64 value to a string
	posStr := strconv.FormatInt(pos, 10)

	fmt.Printf("Read %d events\n", maxEventCount)
	fmt.Printf("End position of binlog: %s\n", posStr)

	/*
		if *tableCounts {
			for tableID, count := range tableCountsMap {
				fmt.Printf("%d %s\n", count, tableID)
			}
	*/
	if *tableCounts {
		fmt.Println("Counting tables...")
		err = decoder.WalkEvent(func(event *binlog.BinEvent) (isContinue bool, err error) {
			if event.Header.EventType == binlog.TableMapEvent {
				tableID := strconv.FormatUint(event.Body.(*binlog.BinTableMapEvent).TableID, 10)
				tableCountsMap[tableID]++
			}
			return true, nil
		})
		if err != nil {
			panic(err)
		}
		for tableID, count := range tableCountsMap {
			fmt.Printf("%d %s\n", count, tableID)
		}
	}
}
