package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	times, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
	delay, _ := strconv.Atoi(strings.Split(os.Args[2], "=")[1])
	factor, _ := strconv.Atoi(strings.Split(os.Args[3], "=")[1])

	println("===============================")
	println(fmt.Sprintf("Times: %d", times))
	println(fmt.Sprintf("Delay: %d", delay))
	println(fmt.Sprintf("Factor: %d", factor))
	println("===============================")

	start := time.Date(2022, 4, 4, 10, 24, 44, 0, time.Local)
	total := 0
	for i := 1; i <= times; i++ {
		result := i * (delay * factor)
		total += result
		ondemand := start.Add(time.Second * time.Duration(total))
		println(fmt.Sprintf("%d | %d | %d | %d | %s", i, delay, factor, result, ondemand.Format(time.RFC3339)))
	}

	println(fmt.Sprintf("Total time in seconds: %d", total))
	println(fmt.Sprintf("Total time in minutes: %d", total/60))
	println(fmt.Sprintf("Total time in hours: %d", (total/60)/60))
	println(fmt.Sprintf("First time: %s", start.Format(time.RFC3339)))

	start = start.Add(time.Second * time.Duration(total))

	println(fmt.Sprintf("Last time: %s", start.Format(time.RFC3339)))

}
