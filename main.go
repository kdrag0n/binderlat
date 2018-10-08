package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

// Android Binder event enum names
const (
	EvReply = "BR_REPLY"
	EvTransaction = "BR_TRANSACTION"
	EvComplete = "BR_TRANSACTION_COMPLETE"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("/dev/kmsg")
	check(err)
	defer file.Close()

	trStartTimes := make(map[string]time.Time)
	var totalTime float64 // usec

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		if strings.HasSuffix(split[0], "binder:") {
			// this is a binder log
			id := split[1] // 8707: 8931
			event := split[2] // BR_REPLY, BR_TRANSACTION, BR_TRANSACTION_COMPLETE

			switch (event) {
			case EvReply:
				fallthrough
			case EvTransaction:
				trStartTimes[id] = time.Now()
			case EvComplete:
				if start, ok := trStartTimes[id]; ok {
					elapsed := time.Since(start)
					usec := float64(elapsed.Nanoseconds()) / 1000
					totalTime += usec

					fmt.Printf("Transaction %11s took %6.1fus; total %4.1fms\n", id, usec, totalTime / 1000)
					delete(trStartTimes, id)
				} else {
					//fmt.Printf("WARNING: Transaction %s completed before start received\n", id)
				}
			}
		}
	}

	check(scanner.Err())
}
