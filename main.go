package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for i := 0; i < len(sa); i++ {
		nu, _ := strconv.Atoi(sa[i])
		si = append(si, nu)
	}
	return si
}
func main() {

	in := [...]string{"2010 12 30 0 0", "2018 12 30 0 0"}
	// in_slice := [...][...]string{}
	inslice := make([][]int, 2)
	temp := make([]string, 5)
	for i := 0; i < len(in); i++ {
		temp = strings.Split(in[i], " ")
		inslice[i] = sliceAtoi(temp)
	}
	// fmt.Printf("%#v\n", inslice)
	timeSlice := make([]time.Time, 2)
	for i := 0; i < len(inslice); i++ {
		timeSlice[i] = time.Date(inslice[i][0], time.Month(inslice[i][1]), inslice[i][2], inslice[i][3], inslice[i][4], 0, 0, time.UTC)
	}
	// fmt.Printf("%#v\n", timeSlice)
	duration := timeSlice[1].Sub(timeSlice[0])
	hours0 := int(duration.Hours())
	year := hours0 / (24 * 365)
	month := (hours0 % 24) / 12
	days := ((hours0 % (24 * 365)) % 12) / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	fmt.Println(year, month, days, hours, mins)
	fmt.Println((hours0 / 24) % 365)
}
