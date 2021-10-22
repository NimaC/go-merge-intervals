package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/NimaC/go-merge-intervals/handler"
	"github.com/NimaC/go-merge-intervals/interval"
)

// Parse Command Line String Input to List of Intervals
func parseInput(inputInterval string) []interval.Interval {
	r := regexp.MustCompile(`\[(?P<Start>[\d]+)\s*,\s*(?P<End>[\d]+)\]+`)
	matches := r.FindAllStringSubmatch(inputInterval, -1)
	var result []interval.Interval
	for _, match := range matches {
		start, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		result = append(result, interval.Interval{start, end})
	}
	return result
}

func main() {
	argsWithoutProg := strings.Join(os.Args[1:], " ")
	inputInterval := parseInput(argsWithoutProg)
	result := handler.GetOverlappingIntervals(inputInterval)
	fmt.Println(result)
}
